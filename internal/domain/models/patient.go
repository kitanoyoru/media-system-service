package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model

	Name string `gorm:"column:name;not null"`

	MedicalWorkerID      uint
	IndicatorInteraction IndicatorInteraction
}
