package controllers

import (
	"bytes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/services/tendency"
	"gorm.io/gorm"
)

type TendencyController struct {
	db              *gorm.DB
	tendencyService *tendency.TendencyService
}

func NewTendencyController(db *gorm.DB, tendencyService *tendency.TendencyService) *TendencyController {
	return &TendencyController{
		db,
		tendencyService,
	}
}

func (c *TendencyController) Route(app *fiber.App) {
	app.Post("/api/v0/tendency", c.getTendencyHandler)
}

func (c *TendencyController) getTendencyHandler(ctx *fiber.Ctx) error {
	getTendencyDTO := new(dtos.GetTendencyDTO)
	if err := ctx.BodyParser(getTendencyDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrResponseDTO{
			Code: fiber.StatusBadRequest,
			Data: struct {
				Message string `json:"message"`
			}{
				Message: fmt.Sprintf("Invalid request body: %s", err),
			},
		})
	}

	bar, err := c.tendencyService.GetTendency(getTendencyDTO)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code: fiber.StatusInternalServerError,
			Data: struct {
				Message string `json:"message"`
			}{
				Message: fmt.Sprintf("Invalid request body: %s", err),
			},
		})
	}

	buffer := bytes.Buffer{}
	bar.Render(&buffer)

	_, err = ctx.Write(buffer.Bytes())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code: fiber.StatusInternalServerError,
			Data: struct {
				Message string `json:"message"`
			}{
				Message: fmt.Sprintf("Invalid request body: %s", err),
			},
		})
	}

	return nil
}
