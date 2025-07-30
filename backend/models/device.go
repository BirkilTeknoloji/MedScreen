package models

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	DeviceID   string `gorm:"uniqueIndex;not null"` // cihaz benzersiz kimliği
	UserID     uint   // 1. Foreign Key alanı: Veritabanında saklanır.
	User       User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 2. İlişki alanı: GORM tarafından veri yüklemek için kullanılır.
	LastSeenAt *time.Time
}
