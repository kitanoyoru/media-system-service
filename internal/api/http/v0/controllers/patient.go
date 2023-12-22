package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"github.com/kitanoyoru/media-system-service/internal/services/patient"
	"gorm.io/gorm"
)

type PatientController struct {
	db             *gorm.DB
	patientService *patient.PatientService
}

func NewPatientController(db *gorm.DB, patientService *patient.PatientService) *PatientController {
	return &PatientController{
		db,
		patientService,
	}
}

func (c *PatientController) Route(app *fiber.App) {
	app.Post("/api/v0/patients", c.getPatientsHandler)
}

func (c *PatientController) getPatientsHandler(ctx *fiber.Ctx) error {
	getPatientsDTO := new(dtos.GetPatientsRequestDTO)
	if err := ctx.BodyParser(getPatientsDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ErrResponseDTO{
			Code: fiber.StatusBadRequest,
			Data: struct {
				Message string `json:"message"`
			}{
				Message: fmt.Sprintf("Invalid request body: %s", err),
			},
		})
	}

	patients, err := c.patientService.GetWorkerPatientsByName(getPatientsDTO.Username)
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

	response := dtos.GetPatientsResponseDTO{
		Code: fiber.StatusOK,
		Data: struct {
			Patients []models.Patient `json:"patients"`
		}{
			Patients: patients,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
