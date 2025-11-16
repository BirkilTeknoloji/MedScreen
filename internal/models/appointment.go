package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment represents a scheduled appointment between a patient and doctor
type Appointment struct {
	gorm.Model
	PatientID       uint              `gorm:"not null;index" json:"patient_id"`
	Patient         *Patient          `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	DoctorID        uint              `gorm:"not null;index" json:"doctor_id"`
	Doctor          *User             `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	AppointmentDate time.Time         `gorm:"not null;index" json:"appointment_date"`
	DurationMinutes int               `gorm:"not null" json:"duration_minutes"`
	AppointmentType AppointmentType   `gorm:"size:50;not null" json:"appointment_type"`
	Status          AppointmentStatus `gorm:"size:50;not null;index" json:"status"`
	Reason          *string           `gorm:"type:text" json:"reason,omitempty"`
	Notes           *string           `gorm:"type:text" json:"notes,omitempty"`
	CreatedByUserID uint              `gorm:"not null" json:"created_by_user_id"`
	CreatedByUser   *User             `gorm:"foreignKey:CreatedByUserID" json:"created_by_user,omitempty"`
}
