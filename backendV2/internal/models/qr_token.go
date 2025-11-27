package models

import (
	"time"

	"gorm.io/gorm"
)

// QRToken represents a QR code token for patient assignment or prescription info
type QRToken struct {
	gorm.Model
	Token     string     `gorm:"uniqueIndex:idx_qr_tokens_token;not null" json:"token"`
	Type      QRType     `gorm:"type:varchar(50);not null;index" json:"type"`
	PatientID uint       `gorm:"not null;index" json:"patient_id"`
	Patient   *Patient   `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	DeviceID  *uint      `gorm:"index" json:"device_id,omitempty"`
	Device    *Device    `gorm:"foreignKey:DeviceID" json:"device,omitempty"`
	ExpiresAt time.Time  `gorm:"not null;index" json:"expires_at"`
	IsUsed    bool       `gorm:"default:false;index" json:"is_used"`
	UsedAt    *time.Time `json:"used_at,omitempty"`
}

// TableName specifies the table name for QRToken
func (QRToken) TableName() string {
	return "qr_tokens"
}

// IsExpired checks if the token has expired
func (q *QRToken) IsExpired() bool {
	return time.Now().After(q.ExpiresAt)
}

// CanBeUsed checks if the token can be used (not expired and not used for patient assignment)
func (q *QRToken) CanBeUsed() bool {
	if q.IsExpired() {
		return false
	}
	// Patient assignment tokens are one-time use
	if q.Type == QRTypePatientAssignment && q.IsUsed {
		return false
	}
	// Prescription info tokens can be used multiple times
	return true
}
