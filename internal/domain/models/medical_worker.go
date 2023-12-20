package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MedicalWorker struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	Patients []Patient
}

func (u *MedicalWorker) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
