package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
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
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	var (
		recommendation bool
		err            error
	)

	switch getRecommendationDTO.IndicatorName {
	case models.HeartRateIndicator:
		recommendation, err = c.recommendationService.PatientHeartRateInNorm(getRecommendationDTO.PatientName, getRecommendationDTO.Indicators)
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal error",
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
