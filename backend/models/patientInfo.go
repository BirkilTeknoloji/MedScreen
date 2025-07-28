package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type PatientInfo struct {
	gorm.Model
	UserID         uint   `gorm:"uniqueIndex;not null"`
	TCNumber       string `gorm:"uniqueIndex;not null"`
	BirthDate      string
	Gender         string
	Phone          string
	Address        string
	Appointments   time.Time
	Diagnosis      []string        `gorm:"serializer:json"`
	Prescriptions  []string        `gorm:"serializer:json"`
	Notes          []string        `gorm:"serializer:json"`
	Tests          json.RawMessage `gorm:"type:jsonb"`
	DoctorID       uint            `gorm:"not null"`
	MedicalHistory json.RawMessage `gorm:"type:jsonb"`
	SurgeryHistory json.RawMessage `gorm:"type:jsonb"`
	Height         float32
	Weight         float32
	Allergies      []string `gorm:"serializer:json"`
	BloodType      string
}
