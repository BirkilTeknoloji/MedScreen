package models

import (
	"time"

	"gorm.io/gorm"
)

// Allergy represents a patient's allergy information
type Allergy struct {
	gorm.Model
	PatientID       uint            `gorm:"not null;index" json:"patient_id"`
	Patient         *Patient        `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	Allergen        string          `gorm:"size:200;not null" json:"allergen"`
	AllergyType     AllergyType     `gorm:"size:50;not null" json:"allergy_type"`
	Reaction        string          `gorm:"size:500;not null" json:"reaction"`
	Severity        AllergySeverity `gorm:"size:50;not null;index" json:"severity"`
	DiagnosedDate   time.Time       `gorm:"not null" json:"diagnosed_date"`
	Notes           *string         `gorm:"type:text" json:"notes,omitempty"`
	AddedByDoctorID uint            `gorm:"not null" json:"added_by_doctor_id"`
	AddedByDoctor   *User           `gorm:"foreignKey:AddedByDoctorID" json:"added_by_doctor,omitempty"`
}
