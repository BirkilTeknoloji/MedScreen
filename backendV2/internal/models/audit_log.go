package models

import (
	"time"
)

// AuditLog represents a record of a change in the system
type AuditLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     *uint     `gorm:"index" json:"user_id,omitempty"` // Nullable for system actions
	User       *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Action     string    `gorm:"size:50;not null;index" json:"action"`       // CREATE, UPDATE, DELETE
	EntityName string    `gorm:"size:100;not null;index" json:"entity_name"` // Table name
	EntityID   uint      `gorm:"not null;index" json:"entity_id"`
	OldValues  string    `gorm:"type:text" json:"old_values,omitempty"` // JSON string
	NewValues  string    `gorm:"type:text" json:"new_values,omitempty"` // JSON string
	IPAddress  string    `gorm:"size:50" json:"ip_address,omitempty"`
	UserAgent  string    `gorm:"size:255" json:"user_agent,omitempty"`
	CreatedAt  time.Time `gorm:"index" json:"created_at"`
}

// BeforeCreate hook is not needed here as we will handle creation manually or via a separate mechanism
