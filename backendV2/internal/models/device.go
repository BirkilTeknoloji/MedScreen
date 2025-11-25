package models

import (
	"gorm.io/gorm"
)

// Device represents a bedside tablet or other device in the system
type Device struct {
	gorm.Model
	MACAddress  string   `gorm:"size:50;uniqueIndex;not null" json:"mac_address"`
	PatientID   *uint    `gorm:"index" json:"patient_id,omitempty"`
	Patient     *Patient `gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"patient,omitempty"`
	RoomNumber  *string  `gorm:"size:20" json:"room_number,omitempty"`
	Description *string  `gorm:"type:text" json:"description,omitempty"`
	IsActive    bool     `gorm:"default:true" json:"is_active"`
}
