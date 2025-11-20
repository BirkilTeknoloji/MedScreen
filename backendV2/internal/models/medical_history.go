package models

import (
	"time"

	"gorm.io/gorm"
)

// MedicalHistory represents a patient's medical history entry
type MedicalHistory struct {
	gorm.Model
	PatientID       uint                 `gorm:"not null;index" json:"patient_id"`
	Patient         *Patient             `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	ConditionName   string               `gorm:"size:200;not null" json:"condition_name"`
	DiagnosedDate   time.Time            `gorm:"not null" json:"diagnosed_date"`
	ResolvedDate    *time.Time           `json:"resolved_date,omitempty"`
	Status          MedicalHistoryStatus `gorm:"size:50;not null;index" json:"status"`
	Notes           *string              `gorm:"type:text" json:"notes,omitempty"`
	AddedByDoctorID uint                 `gorm:"not null" json:"added_by_doctor_id"`
	AddedByDoctor   *User                `gorm:"foreignKey:AddedByDoctorID" json:"added_by_doctor,omitempty"`
}

// TableName overrides the default table name
func (MedicalHistory) TableName() string {
	return "medical_history"
}
