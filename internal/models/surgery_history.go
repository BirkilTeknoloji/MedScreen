package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// SurgeryHistory represents a patient's surgery history entry
type SurgeryHistory struct {
	gorm.Model
	PatientID       uint      `gorm:"not null;index" json:"patient_id"`
	Patient         *Patient  `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	ProcedureName   string    `gorm:"size:200;not null" json:"procedure_name"`
	SurgeryDate     time.Time `gorm:"not null;index" json:"surgery_date"`
	SurgeonName     string    `gorm:"size:200;not null" json:"surgeon_name"`
	Complications   *string   `gorm:"type:text" json:"complications,omitempty"`
	Notes           *string   `gorm:"type:text" json:"notes,omitempty"`
	AddedByDoctorID uint      `gorm:"not null" json:"added_by_doctor_id"`
	AddedByDoctor   *User     `gorm:"foreignKey:AddedByDoctorID" json:"added_by_doctor,omitempty"`
}

// TableName overrides the default table name
func (SurgeryHistory) TableName() string {
	return "surgery_history"
}

// BeforeCreate validates the surgery history data before creating
func (s *SurgeryHistory) BeforeCreate(tx *gorm.DB) error {
	return s.validate()
}

// BeforeUpdate validates the surgery history data before updating
func (s *SurgeryHistory) BeforeUpdate(tx *gorm.DB) error {
	return s.validate()
}

// validate performs validation on surgery history data
func (s *SurgeryHistory) validate() error {
	// Validate surgery_date is not in the future
	if s.SurgeryDate.After(time.Now()) {
		return errors.New("surgery_date cannot be in the future")
	}

	return nil
}
