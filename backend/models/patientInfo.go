package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// PatientInfo represents detailed patient information stored in the database.
type PatientInfo struct {
	gorm.Model
	UserID         uint `gorm:"uniqueIndex;not null"` // Links to the User table
	Name           string
	TCNumber       string `gorm:"uniqueIndex;not null"` // Turkish National ID
	BirthDate      string
	Gender         string
	Phone          string
	Address        string
	Appointments   Appointments    `gorm:"type:jsonb"` // JSONB field
	Diagnosis      Diagnoses       `gorm:"type:jsonb"`
	Prescriptions  Prescriptions   `gorm:"type:jsonb"`
	Notes          Notes           `gorm:"type:jsonb"`
	Tests          Tests           `gorm:"type:jsonb"`
	DoctorID       uint            `gorm:"not null"`
	MedicalHistory json.RawMessage `gorm:"type:jsonb"` // Free-form JSON
	SurgeryHistory json.RawMessage `gorm:"type:jsonb"`
	Height         float32
	Weight         float32
	Allergies      Allergies `gorm:"type:jsonb"`
	BloodType      string
}

// Appointment represents a scheduled meeting with a doctor.
type Appointment struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type Appointments []Appointment

func (a *Appointments) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Appointments: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

func (a Appointments) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Diagnosis represents a medical diagnosis entry.
type Diagnosis struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Diagnoses []Diagnosis

func (d *Diagnoses) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Diagnoses: %v", value)
	}
	return json.Unmarshal(bytes, d)
}

func (d Diagnoses) Value() (driver.Value, error) {
	return json.Marshal(d)
}

// Prescription represents a prescribed medication.
type Prescription struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Prescriptions []Prescription

func (p *Prescriptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Prescriptions: %v", value)
	}
	return json.Unmarshal(bytes, p)
}

func (p Prescriptions) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Note represents a doctor's note.
type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Notes []Note

func (n *Notes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Notes: %v", value)
	}
	return json.Unmarshal(bytes, n)
}

func (n Notes) Value() (driver.Value, error) {
	return json.Marshal(n)
}

// Allergy represents a known allergy.
type Allergy struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Allergies []Allergy

func (a *Allergies) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Allergies: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

func (a Allergies) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Test represents a medical test.
type Test struct {
	ID     string          `json:"id"`
	Title  string          `json:"title"`
	Result json.RawMessage `json:"result"` // Flexible: string or array
}

type Tests []Test

func (t *Tests) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Tests: %v", value)
	}
	return json.Unmarshal(bytes, t)
}

func (t Tests) Value() (driver.Value, error) {
	return json.Marshal(t)
}
