package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Role   string `gorm:"not null"` // "patient" veya "doctor"
	CardID string `gorm:"uniqueIndex;not null"`

	// Eğer hasta ise hasta bilgisi olabilir
	PatientInfo PatientInfo `gorm:"foreignKey:UserID"`
}

// type PatientInfo struct {
// 	gorm.Model
// 	UserID     uint `gorm:"uniqueIndex"` // Her hastaya bir bilgi
// 	Diagnosis  string
// 	Allergies  string
// 	Treatments string
// }
