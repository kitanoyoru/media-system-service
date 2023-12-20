package recommendation

import (
	"errors"

	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"github.com/kitanoyoru/media-system-service/pkg/helpers"
	"gorm.io/gorm"
)

const (
	HeartBeatNormDifference     = 10.0
	BloodPressureNormDifference = 15.0
	WeightNormDifference        = 20.0
)

type RecommendationService struct {
	db *gorm.DB
}

func NewRecommendationService(db *gorm.DB) *RecommendationService {
	return &RecommendationService{
		db,
	}
}

func (rs *RecommendationService) GetRecommendationByName(indicatorName, patientName string, indicators []float64) (bool, error) {
	switch indicatorName {
	case models.HeartRateIndicator:
		return rs.PatientHeartRateInNorm(patientName, indicators)
	case models.BloodPressureIndicator:
		return rs.PatientBloodPressureInNorm(patientName, indicators)
	default:
		return false, errors.New("failed to get recommendation")
	}
}

func (rs *RecommendationService) PatientHeartRateInNorm(patientName string, indicators []float64) (bool, error) {
	patient, err := rs.getPatientByName(patientName)
	if err != nil {
		return false, err
	}

	storedIndicators, err := patient.IndicatorInteraction.GetDynamicIndicators(models.HeartRateIndicator)
	if err != nil {
		return false, err
	}

	return rs.checkIsNorm(indicators, storedIndicators, HeartBeatNormDifference), nil
}

func (rs *RecommendationService) PatientBloodPressureInNorm(patientName string, indicators []float64) (bool, error) {
	patient, err := rs.getPatientByName(patientName)
	if err != nil {
		return false, err
	}

	storedIndicators, err := patient.IndicatorInteraction.GetDynamicIndicators(models.BloodPressureIndicator)
	if err != nil {
		return false, err
	}

	return rs.checkIsNorm(indicators, storedIndicators, BloodPressureNormDifference), nil
}

func (rs *RecommendationService) PatientWeightInNorm(patientName string, indicator float64) (bool, error) {
	patient, err := rs.getPatientByName(patientName)
	if err != nil {
		return false, err
	}

	storedIndicator, err := patient.IndicatorInteraction.GetStaticIndicators(models.WeightIndicator)
	if err != nil {
		return false, err
	}

	return rs.checkIsNorm([]float64{indicator}, []float64{storedIndicator}, WeightNormDifference), nil
}

func (rs *RecommendationService) getPatientByName(name string) (*models.Patient, error) {
	patient := models.Patient{}
	if err := rs.db.Preload("IndicatorInteraction").Preload("MedicalWorker").First(&patient, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (rs *RecommendationService) checkIsNorm(requested, stored []float64, difference float64) bool {
	rAvg, sAvg := helpers.Avg(requested), helpers.Avg(stored)

	return rAvg >= sAvg-difference && rAvg <= sAvg+difference
}
