package controllers

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/services/tendencies"
	"gorm.io/gorm"
)

type TendencyController struct {
	db              *gorm.DB
	tendencyService *tendencies.TendencyService
}

func NewTendencyController(db *gorm.DB, tendencyService *tendencies.TendencyService) *TendencyController {
	return &TendencyController{
		db,
		tendencyService,
	}
}

func (c *TendencyController) Route(app *fiber.App) {
	app.Get("/api/v0/tendency", c.getTendencyHandler)
}

func (c *TendencyController) getTendencyHandler(ctx *fiber.Ctx) error {
	getTendencyDTO := new(dtos.GetTendencyDTO)
	if err := ctx.QueryParser(getTendencyDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	bar, err := c.tendencyService.GetTendency(getTendencyDTO)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal error",
		})
	}

	buffer := bytes.Buffer{}
	bar.Render(&buffer)

	_, err = ctx.Write(buffer.Bytes())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal error",
		})
	}

	return nil
}
