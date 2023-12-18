package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	// static indicators
	HeightIndicator = iota
	WeightIndicator

	// dynamic indicators
	HeartRateIndicator
	BloodPressureIndicator
)

var (
	ErrGetStaticIndicators  = errors.New("failed to get static indicators")
	ErrGetDynamicIndicators = errors.New("failed to get dynamic indicators")
)

type Height struct {
	gorm.Model
	Value         float64   `gorm:"column:value"`
	Timestamp     time.Time `gorm:"column:timestamp"`
	InteractionID uint      `gorm:"column:interaction_id"`
}

type Weight struct {
	gorm.Model
	Value         float64   `gorm:"column:value"`
	Timestamp     time.Time `gorm:"column:timestamp"`
	InteractionID uint      `gorm:"column:interaction_id"`
}

type HeartRate struct {
	gorm.Model
	Value         int64     `gorm:"column:value"`
	Timestamp     time.Time `gorm:"column:timestamp"`
	InteractionID uint      `gorm:"column:interaction_id"`
}

type BloodPressure struct {
	gorm.Model
	Value         int64     `gorm:"column:value"`
	Timestamp     time.Time `gorm:"column:timestamp"`
	InteractionID uint      `gorm:"column:interaction_id"`
}

type IndicatorInteraction struct {
	gorm.Model

	Height         Height          `gorm:"foreignKey:InteractionID"`
	Weight         Weight          `gorm:"foreignKey:InteractionID"`
	HeartRates     []HeartRate     `gorm:"foreignKey:InteractionID"`
	BloodPressures []BloodPressure `gorm:"foreignKey:InteractionID"`

	PatientID uint
}

func (i IndicatorInteraction) GetDynamicIndicators(indicatorName int) ([]float64, error) {
	switch indicatorName {
	case HeartRateIndicator:
		return i.getHeartRates(), nil
	case BloodPressureIndicator:
		return i.getBloodPressures(), nil
	default:
		return nil, ErrGetDynamicIndicators
	}
}

func (i IndicatorInteraction) SendDynamicIndicators(indicatorName int, indicatorValues []float64) {
	switch indicatorName {
	case HeartRateIndicator:
		i.addHeartRates(indicatorValues)
	case BloodPressureIndicator:
		i.addBloodPressures(indicatorValues)
	}
}

func (i IndicatorInteraction) GetStaticIndicators(indicatorName int) (float64, error) {
	switch indicatorName {
	case HeightIndicator:
		return i.getHeight(), nil
	case WeightIndicator:
		return i.getWeight(), nil
	default:
		return 0.0, ErrGetStaticIndicators
	}
}

func (i IndicatorInteraction) SetStaticIndicators(indicatorName int, value float64) {
	switch indicatorName {
	case HeightIndicator:
		i.setHeight(value)
	case WeightIndicator:
		i.setWeight(value)
	}
}

func (i IndicatorInteraction) getHeartRates() []float64 {
	heartRates := make([]float64, len(i.HeartRates))
	for idx, hr := range i.HeartRates {
		heartRates[idx] = float64(hr.Value)
	}
	return heartRates
}

func (i IndicatorInteraction) addHeartRates(values []float64) {
	currentTime := time.Now()
	for _, value := range values {
		heartRate := HeartRate{
			Value:         int64(value),
			Timestamp:     currentTime,
			InteractionID: i.ID,
		}
		i.HeartRates = append(i.HeartRates, heartRate)
	}
}

func (i IndicatorInteraction) getBloodPressures() []float64 {
	bloodPressures := make([]float64, len(i.BloodPressures))
	for idx, bp := range i.BloodPressures {
		bloodPressures[idx] = float64(bp.Value)
	}
	return bloodPressures
}

func (i IndicatorInteraction) addBloodPressures(values []float64) {
	currentTime := time.Now()
	for _, value := range values {
		bloodPressure := BloodPressure{
			Value:         int64(value),
			Timestamp:     currentTime,
			InteractionID: i.ID,
		}
		i.BloodPressures = append(i.BloodPressures, bloodPressure)
	}
}

func (i IndicatorInteraction) getHeight() float64 {
	return i.Height.Value
}

func (i IndicatorInteraction) setHeight(value float64) {
	i.Height.Value = value
}

func (i IndicatorInteraction) getWeight() float64 {
	return i.Weight.Value
}

func (i IndicatorInteraction) setWeight(value float64) {
	i.Weight.Value = value
}
