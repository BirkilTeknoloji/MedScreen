package models

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	DeviceID   string `gorm:"uniqueIndex;not null"` // device unique ID
	UserID     uint   // 1. Foreign Key field: Stored in the database.
	User       User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 2. Relationship field: Used by GORM to load data.
	LastSeenAt *time.Time
}
