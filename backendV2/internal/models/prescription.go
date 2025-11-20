package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Prescription represents a medication prescription for a patient
type Prescription struct {
	gorm.Model
	PatientID      uint               `gorm:"not null;index" json:"patient_id"`
	Patient        *Patient           `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	AppointmentID  *uint              `gorm:"index" json:"appointment_id,omitempty"`
	Appointment    *Appointment       `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	DoctorID       uint               `gorm:"not null;index" json:"doctor_id"`
	Doctor         *User              `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	PrescribedDate time.Time          `gorm:"not null;index" json:"prescribed_date"`
	MedicationName string             `gorm:"size:200;not null" json:"medication_name"`
	Dosage         string             `gorm:"size:100;not null" json:"dosage"`
	Frequency      string             `gorm:"size:100;not null" json:"frequency"`
	Duration       string             `gorm:"size:100;not null" json:"duration"`
	Quantity       int                `gorm:"not null" json:"quantity"`
	RefillsAllowed int                `gorm:"default:0" json:"refills_allowed"`
	Instructions   *string            `gorm:"type:text" json:"instructions,omitempty"`
	Status         PrescriptionStatus `gorm:"size:50;not null" json:"status"`
}

// BeforeCreate validates the prescription data before creating
func (p *Prescription) BeforeCreate(tx *gorm.DB) error {
	return p.validate()
}

// BeforeUpdate validates the prescription data before updating
func (p *Prescription) BeforeUpdate(tx *gorm.DB) error {
	return p.validate()
}

// validate performs validation on prescription data
func (p *Prescription) validate() error {
	// Validate quantity is positive
	if p.Quantity <= 0 {
		return errors.New("quantity must be a positive integer")
	}

	return nil
}
