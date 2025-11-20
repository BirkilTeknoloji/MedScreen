package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Patient represents a patient in the medical system
type Patient struct {
	gorm.Model
	UserID                *uint     `json:"user_id,omitempty"`
	FirstName             string    `gorm:"size:100;not null" json:"first_name"`
	LastName              string    `gorm:"size:100;not null;index" json:"last_name"`
	TCNumber              string    `gorm:"size:11;not null;uniqueIndex" json:"tc_number"`
	BirthDate             time.Time `gorm:"not null" json:"birth_date"`
	Gender                Gender    `gorm:"size:10;not null" json:"gender"`
	Phone                 string    `gorm:"size:20" json:"phone"`
	Email                 *string   `gorm:"size:100" json:"email,omitempty"`
	Address               *string   `gorm:"type:text" json:"address,omitempty"`
	EmergencyContactName  *string   `gorm:"size:100" json:"emergency_contact_name,omitempty"`
	EmergencyContactPhone *string   `gorm:"size:20" json:"emergency_contact_phone,omitempty"`
	BloodType             *string   `gorm:"size:5" json:"blood_type,omitempty"`
	Height                *float64  `json:"height,omitempty"`
	Weight                *float64  `json:"weight,omitempty"`
	PrimaryDoctorID       *uint     `gorm:"index" json:"primary_doctor_id,omitempty"`
	PrimaryDoctor         *User     `gorm:"foreignKey:PrimaryDoctorID" json:"primary_doctor,omitempty"`
}

// BeforeCreate validates the patient data before creating
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	return p.validate()
}

// BeforeUpdate validates the patient data before updating
func (p *Patient) BeforeUpdate(tx *gorm.DB) error {
	return p.validate()
}

// validate performs validation on patient data
func (p *Patient) validate() error {
	// Validate TC number is exactly 11 digits
	if len(p.TCNumber) != 11 {
		return errors.New("tc_number must be exactly 11 digits")
	}

	// Validate TC number contains only digits
	for _, char := range p.TCNumber {
		if char < '0' || char > '9' {
			return errors.New("tc_number must contain only digits")
		}
	}

	// Validate gender enum
	if p.Gender != GenderMale && p.Gender != GenderFemale {
		return errors.New("gender must be either 'male' or 'female'")
	}

	return nil
}
