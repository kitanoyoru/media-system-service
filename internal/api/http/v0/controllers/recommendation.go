package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/services/recommendation"
	"gorm.io/gorm"
)

type RecommendationController struct {
	db                    *gorm.DB
	recommendationService *recommendation.RecommendationService
}

func NewRecommendationController(db *gorm.DB, recommendationService *recommendation.RecommendationService) *RecommendationController {
	return &RecommendationController{
		db,
		recommendationService,
	}
}

func (c *RecommendationController) Route(app *fiber.App) {
	app.Post("/api/v0/recommendation", c.getRecommendationHandler)
}

func (c *RecommendationController) getRecommendationHandler(ctx *fiber.Ctx) error {
	getRecommendationDTO := new(dtos.PostRecommendationRequestDTO)
	if err := ctx.BodyParser(getRecommendationDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrResponseDTO{
			Code: fiber.StatusBadRequest,
			Data: struct {
				Message string `json:"message"`
			}{
				Message: "Invalid request body",
			},
		})
	}

	recommendation, err := c.recommendationService.GetRecommendationByName(getRecommendationDTO.IndicatorName, getRecommendationDTO.PatientName, getRecommendationDTO.Indicators)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code: fiber.StatusInternalServerError,
			Data: struct {
				Message string `json:"message"`
			}{
				Message: "Internal Error",
			},
		})
	}

	response := dtos.PostRecommendationResponseDTO{
		Code: fiber.StatusOK,
		Data: struct {
			Answer bool `json:"answer"`
		}{
			Answer: recommendation,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
