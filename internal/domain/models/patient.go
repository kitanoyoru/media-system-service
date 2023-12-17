package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model

	Name string `gorm:"not null"`

	MedicalWorkerID      uint
	IndicatorInteraction IndicatorInteraction
}
