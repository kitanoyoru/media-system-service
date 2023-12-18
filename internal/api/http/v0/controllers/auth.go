package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"github.com/kitanoyoru/media-system-service/internal/services/auth"
	"gorm.io/gorm"
)

type AuthController struct {
	db          *gorm.DB
	authService *auth.AuthService
}

func NewAuthController(db *gorm.DB, authService *auth.AuthService) *AuthController {
	return &AuthController{
		db:          db,
		authService: authService,
	}
}

func (c *AuthController) Route(app *fiber.App) {
	app.Post("/login", c.loginHandler)
	app.Post("/register", c.registerHandler)
}

func (c *AuthController) loginHandler(ctx *fiber.Ctx) error {
	loginDTO := new(dtos.LoginRequestDTO)
	if err := ctx.BodyParser(loginDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	token, err := c.authService.GetJWTToken(loginDTO)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to get JWT token",
		})
	}

	response := dtos.LoginResponseDTO{
		Code: fiber.StatusOK,
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		}}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *AuthController) registerHandler(ctx *fiber.Ctx) error {
	registerDTO := new(dtos.RegisterRequestDTO)
	if err := ctx.BodyParser(registerDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	medicalWorker := models.MedicalWorker{
		Username: registerDTO.Username,
		Password: registerDTO.Password,
	}

	if err := c.db.Save(&medicalWorker).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ErrResponseDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal error",
		})
	}

	response := dtos.RegisterResponseDTO{
		Code: fiber.StatusOK,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
