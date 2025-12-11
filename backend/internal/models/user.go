package models

import "gorm.io/gorm"

// User represents a system user (doctor, nurse, receptionist, admin)
type User struct {
	gorm.Model
	FirstName      string   `gorm:"size:100;not null" json:"first_name"`
	LastName       string   `gorm:"size:100;not null" json:"last_name"`
	Role           UserRole `gorm:"size:50;not null;index" json:"role"`
	Specialization *string  `gorm:"size:100" json:"specialization,omitempty"`
	LicenseNumber  *string  `gorm:"size:50" json:"license_number,omitempty"`
	Phone          string   `gorm:"size:20" json:"phone"`
	IsActive       bool     `gorm:"default:true" json:"is_active"`
	NFCCardID      *uint    `gorm:"index" json:"nfc_card_id,omitempty"` // Foreign key to nfc_cards.id
}
