package patient

import (
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"gorm.io/gorm"
)

type PatientService struct {
	db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{
		db,
	}
}

func (ps *PatientService) GetWorkerPatientsByName(name string) ([]models.Patient, error) {
	var worker models.MedicalWorker

	if err := ps.db.Preload("Patients").Model(&worker).Where("username = ?", name).First(&worker).Error; err != nil {
		return nil, err
	}

	return worker.Patients, nil
}
