package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type PatientInfo struct {
	gorm.Model
	UserID         uint   `gorm:"uniqueIndex;not null"` // User modeline referans
	TCNumber       string `gorm:"uniqueIndex;not null"` // Türkiye Cumhuriyeti No
	BirthDate      string
	Gender         string
	Phone          string
	Address        string
	Appointments   time.Time
	Diagnosis      []string        `gorm:"type:text[]"`
	Prescriptions  []string        `gorm:"type:text[]"`
	Notes          []string        `gorm:"type:text[]"`
	Tests          json.RawMessage `gorm:"type:jsonb"`
	DoctorID       uint            `gorm:"not null"`   // İlgili doktorun ID'si
	MedicalHistory json.RawMessage `gorm:"type:jsonb"` // Tıbbi geçmiş bilgisi JSON formatında
	SurgeryHistory json.RawMessage `gorm:"type:jsonb"` // Cerrahi geçmiş bilgisi JSON formatında
	Height         float32
	Weight         float32
	Allergies      []string `gorm:"type:text[]"`
	BloodType      string
}
