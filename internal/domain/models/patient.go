package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model

	Name string `gorm:"column:name;not null" json:"name"`

	MedicalWorkerID      uint                 `json:"-"`
	IndicatorInteraction IndicatorInteraction `json:"-"`
}
