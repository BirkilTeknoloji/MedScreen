package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// PatientInfo represents detailed patient information stored in the database.
// It includes personal info, appointments, diagnoses, tests, and other medical data.
type PatientInfo struct {
	gorm.Model
	UserID         uint `gorm:"uniqueIndex;not null"` // Links to the User table
	Name           string
	TCNumber       string `gorm:"uniqueIndex;not null"` // Turkish National ID
	BirthDate      string
	Gender         string
	Phone          string
	Address        string
	Appointments   Appointments    `gorm:"type:jsonb"` // List of appointments in JSONB
	Diagnosis      Diagnoses       `gorm:"type:jsonb"` // List of diagnoses in JSONB
	Prescriptions  Prescriptions   `gorm:"type:jsonb"` // List of prescriptions in JSONB
	Notes          Notes           `gorm:"type:jsonb"` // Doctor notes in JSONB
	Tests          Tests           `gorm:"type:jsonb"` // Medical tests in JSONB
	DoctorID       uint            `gorm:"not null"`   // ID of the assigned doctor
	MedicalHistory json.RawMessage `gorm:"type:jsonb"` // Free-form medical history in JSONB
	SurgeryHistory json.RawMessage `gorm:"type:jsonb"` // Free-form surgery history in JSONB
	Height         float32
	Weight         float32
	Allergies      Allergies `gorm:"type:jsonb"` // List of known allergies
	BloodType      string
}

// Appointment represents a scheduled meeting with a doctor.
type Appointment struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

// Appointments is a slice of Appointment objects.
type Appointments []Appointment

// Scan implements the sql.Scanner interface for Appointments.
// It converts JSONB from the database into the Go struct.
func (a *Appointments) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Appointments: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

// Value implements the driver.Valuer interface for Appointments.
// It converts the Go struct into JSON to be stored in the database.
func (a Appointments) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Diagnosis represents a medical diagnosis entry.
type Diagnosis struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Diagnoses is a slice of Diagnosis objects.
type Diagnoses []Diagnosis

// Scan implements the sql.Scanner interface for Diagnoses.
func (d *Diagnoses) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Diagnoses: %v", value)
	}
	return json.Unmarshal(bytes, d)
}

// Value implements the driver.Valuer interface for Diagnoses.
func (d Diagnoses) Value() (driver.Value, error) {
	return json.Marshal(d)
}

// Prescription represents a prescribed medication.
type Prescription struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Prescriptions is a slice of Prescription objects.
type Prescriptions []Prescription

// Scan implements the sql.Scanner interface for Prescriptions.
func (p *Prescriptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Prescriptions: %v", value)
	}
	return json.Unmarshal(bytes, p)
}

// Value implements the driver.Valuer interface for Prescriptions.
func (p Prescriptions) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Note represents a note written by a doctor.
type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Notes is a slice of Note objects.
type Notes []Note

// Scan implements the sql.Scanner interface for Notes.
func (n *Notes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Notes: %v", value)
	}
	return json.Unmarshal(bytes, n)
}

// Value implements the driver.Valuer interface for Notes.
func (n Notes) Value() (driver.Value, error) {
	return json.Marshal(n)
}

// Allergy represents a known allergy.
type Allergy struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Allergies is a slice of Allergy objects.
type Allergies []Allergy

// Scan implements the sql.Scanner interface for Allergies.
func (a *Allergies) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Allergies: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

// Value implements the driver.Valuer interface for Allergies.
func (a Allergies) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Test represents a medical test and its result.
type Test struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Result string `json:"result"`
}

// Tests is a slice of Test objects.
type Tests []Test

// Scan implements the sql.Scanner interface for Tests.
func (t *Tests) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Tests: %v", value)
	}
	return json.Unmarshal(bytes, t)
}

// Value implements the driver.Valuer interface for Tests.
func (t Tests) Value() (driver.Value, error) {
	return json.Marshal(t)
}
