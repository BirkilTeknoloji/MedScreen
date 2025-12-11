package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// VitalSign represents a patient's vital signs measurement
type VitalSign struct {
	ID                     uint         `gorm:"primaryKey" json:"id"`
	CreatedAt              time.Time    `json:"created_at"`
	PatientID              uint         `gorm:"not null;index" json:"patient_id"`
	Patient                *Patient     `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	AppointmentID          *uint        `gorm:"index" json:"appointment_id,omitempty"`
	Appointment            *Appointment `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	RecordedByUserID       uint         `gorm:"not null" json:"recorded_by_user_id"`
	RecordedByUser         *User        `gorm:"foreignKey:RecordedByUserID" json:"recorded_by_user,omitempty"`
	RecordedAt             time.Time    `gorm:"not null;index" json:"recorded_at"`
	BloodPressureSystolic  *int         `json:"blood_pressure_systolic,omitempty"`
	BloodPressureDiastolic *int         `json:"blood_pressure_diastolic,omitempty"`
	HeartRate              *int         `json:"heart_rate,omitempty"`
	Temperature            *float64     `json:"temperature,omitempty"`
	RespiratoryRate        *int         `json:"respiratory_rate,omitempty"`
	OxygenSaturation       *float64     `json:"oxygen_saturation,omitempty"`
	Height                 *float64     `json:"height,omitempty"`
	Weight                 *float64     `json:"weight,omitempty"`
	BMI                    *float64     `json:"bmi,omitempty"`
	Notes                  *string      `gorm:"type:text" json:"notes,omitempty"`
}

// BeforeCreate hook to calculate BMI and validate data
func (v *VitalSign) BeforeCreate(tx *gorm.DB) error {
	v.calculateBMI()
	return v.validate()
}

// BeforeUpdate hook to calculate BMI and validate data
func (v *VitalSign) BeforeUpdate(tx *gorm.DB) error {
	v.calculateBMI()
	return v.validate()
}

// calculateBMI calculates BMI if height and weight are provided
func (v *VitalSign) calculateBMI() {
	if v.Height != nil && v.Weight != nil && *v.Height > 0 {
		// BMI = weight (kg) / (height (m))^2
		heightInMeters := *v.Height / 100.0
		bmi := *v.Weight / (heightInMeters * heightInMeters)
		v.BMI = &bmi
	}
}

// validate performs validation on vital sign data
func (v *VitalSign) validate() error {
	// Validate blood pressure: systolic > diastolic
	if v.BloodPressureSystolic != nil && v.BloodPressureDiastolic != nil {
		if *v.BloodPressureSystolic <= *v.BloodPressureDiastolic {
			return errors.New("blood_pressure_systolic must be greater than blood_pressure_diastolic")
		}
	}

	return nil
}
