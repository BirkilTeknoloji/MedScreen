package models

import (
	"time"

	"gorm.io/gorm"
)

// Diagnosis represents a medical diagnosis for a patient
type Diagnosis struct {
	gorm.Model
	PatientID     uint              `gorm:"not null;index" json:"patient_id"`
	Patient       *Patient          `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	AppointmentID *uint             `gorm:"index" json:"appointment_id,omitempty"`
	Appointment   *Appointment      `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	DoctorID      uint              `gorm:"not null;index" json:"doctor_id"`
	Doctor        *User             `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	DiagnosisDate time.Time         `gorm:"not null;index" json:"diagnosis_date"`
	ICDCode       string            `gorm:"size:20;index" json:"icd_code"`
	DiagnosisName string            `gorm:"size:200;not null" json:"diagnosis_name"`
	Description   *string           `gorm:"type:text" json:"description,omitempty"`
	Severity      DiagnosisSeverity `gorm:"size:50;not null" json:"severity"`
	Status        DiagnosisStatus   `gorm:"size:50;not null" json:"status"`
}
