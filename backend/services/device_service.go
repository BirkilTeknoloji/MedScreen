package services

import (
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

// RegisterDevice finds or creates a device and updates its userID, then loads related User and PatientInfo.
func (s *DeviceService) RegisterDevice(deviceID string, userID uint) (*models.Device, error) {
	var device models.Device

	if err := s.DB.Where(models.Device{DeviceID: deviceID}).
		Assign(models.Device{UserID: userID}).
		FirstOrCreate(&device).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Preload("User.PatientInfo").
		First(&device, device.ID).Error; err != nil {
		return nil, err
	}

	return &device, nil
}

// ListUserDevices returns all devices for a given user.
func (s *DeviceService) ListUserDevices(userID uint) ([]models.Device, error) {
	var devices []models.Device
	err := s.DB.Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}

// CountUserDevices returns the count of devices for a given user.
func (s *DeviceService) CountUserDevices(userID uint) (int64, error) {
	var cnt int64
	err := s.DB.Model(&models.Device{}).Where("user_id = ?", userID).Count(&cnt).Error
	return cnt, err
}

func ptrTime(t time.Time) *time.Time {
	return &t
}
