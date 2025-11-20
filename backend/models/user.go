package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Role   string `gorm:"not null"` // "patient" or "personnel"
	CardID string `gorm:"uniqueIndex;not null"`

	// If the user is a patient, there may be patient information
	PatientInfo PatientInfo `gorm:"foreignKey:UserID"`
}
