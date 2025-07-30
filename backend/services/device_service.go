package services

import (
	"errors"
	"go-backend/models"
	"time"

	"gorm.io/gorm"
)

type DeviceService struct {
	DB *gorm.DB
}

func NewDeviceService(db *gorm.DB) *DeviceService {
	return &DeviceService{DB: db}
}

func (s *DeviceService) RegisterDevice(deviceID string, userID uint) (*models.Device, error) {
	var existing models.Device
	err := s.DB.Where("device_id = ?", deviceID).First(&existing).Error
	if err == nil {
		existing.LastSeenAt = ptrTime(time.Now())
		s.DB.Save(&existing)
		return &existing, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	d := models.Device{
		DeviceID:   deviceID,
		UserID:     userID, // Bu satırın `UserID` olduğundan emin olun, `User` değil.
		LastSeenAt: ptrTime(time.Now()),
	}
	if err := s.DB.Create(&d).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (s *DeviceService) ListUserDevices(userID uint) ([]models.Device, error) {
	var devices []models.Device
	err := s.DB.Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}

func (s *DeviceService) CountUserDevices(userID uint) (int64, error) {
	var cnt int64
	err := s.DB.Model(&models.Device{}).Where("user_id = ?", userID).Count(&cnt).Error
	return cnt, err
}

func ptrTime(t time.Time) *time.Time {
	return &t
}
