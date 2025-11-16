package models

import (
	"time"

	"gorm.io/gorm"
)

// MedicalTest represents a medical test ordered for a patient
type MedicalTest struct {
	gorm.Model
	PatientID         uint         `gorm:"not null;index" json:"patient_id"`
	Patient           *Patient     `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	AppointmentID     *uint        `gorm:"index" json:"appointment_id,omitempty"`
	Appointment       *Appointment `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	OrderedByDoctorID uint         `gorm:"not null;index" json:"ordered_by_doctor_id"`
	OrderedByDoctor   *User        `gorm:"foreignKey:OrderedByDoctorID" json:"ordered_by_doctor,omitempty"`
	TestType          TestType     `gorm:"size:50;not null;index" json:"test_type"`
	TestName          string       `gorm:"size:200;not null" json:"test_name"`
	OrderedDate       time.Time    `gorm:"not null" json:"ordered_date"`
	ScheduledDate     *time.Time   `json:"scheduled_date,omitempty"`
	CompletedDate     *time.Time   `json:"completed_date,omitempty"`
	Results           *string      `gorm:"type:text" json:"results,omitempty"`
	ResultFilePath    *string      `gorm:"size:500" json:"result_file_path,omitempty"`
	Status            TestStatus   `gorm:"size:50;not null;index" json:"status"`
	LabName           *string      `gorm:"size:200" json:"lab_name,omitempty"`
	Notes             *string      `gorm:"type:text" json:"notes,omitempty"`
}
