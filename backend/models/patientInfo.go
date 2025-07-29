package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type PatientInfo struct {
	gorm.Model
	UserID         uint `gorm:"uniqueIndex;not null"`
	DeviceID       string
	Name           string
	TCNumber       string `gorm:"uniqueIndex;not null"`
	BirthDate      string
	Gender         string
	Phone          string
	Address        string
	Appointments   Appointments    `gorm:"type:jsonb"`
	Diagnosis      Diagnoses       `gorm:"type:jsonb"`
	Prescriptions  Prescriptions   `gorm:"type:jsonb"`
	Notes          Notes           `gorm:"type:jsonb"`
	Tests          Tests           `gorm:"type:jsonb"`
	DoctorID       uint            `gorm:"not null"`
	MedicalHistory json.RawMessage `gorm:"type:jsonb"`
	SurgeryHistory json.RawMessage `gorm:"type:jsonb"`
	Height         float32
	Weight         float32
	Allergies      Allergies `gorm:"type:jsonb"`
	BloodType      string
}

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

type Test struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Result string `json:"result"`
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
