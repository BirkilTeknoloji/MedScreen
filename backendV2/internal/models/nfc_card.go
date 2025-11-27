package models

import (
	"time"

	"gorm.io/gorm"
)

// NFCCard represents an NFC card used for user authentication
type NFCCard struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	CardUID         string     `gorm:"size:100;not null;uniqueIndex:idx_nfc_cards_card_uid" json:"card_uid"`
	AssignedUserID  *uint      `gorm:"index" json:"assigned_user_id,omitempty"`
	AssignedUser    *User      `gorm:"foreignKey:AssignedUserID" json:"assigned_user,omitempty"`
	IsActive        bool       `gorm:"default:true;index" json:"is_active"`
	IssuedAt        time.Time  `gorm:"not null" json:"issued_at"`
	LastUsedAt      *time.Time `json:"last_used_at,omitempty"`
	CreatedByUserID uint       `json:"created_by_user_id"`
	CreatedByUser   *User      `gorm:"foreignKey:CreatedByUserID" json:"created_by_user,omitempty"`
}

// BeforeUpdate hook to update last_used_at timestamp
func (n *NFCCard) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now()
	n.LastUsedAt = &now
	return nil
}
