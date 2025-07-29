package models

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	DeviceID   string `gorm:"uniqueIndex;not null"` // cihaz benzersiz kimliği
	UserID     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LastSeenAt *time.Time
}
