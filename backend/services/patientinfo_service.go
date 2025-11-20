package services

import (
	"errors"
	"fmt"
	"go-backend/config"
	"go-backend/models"

	"gorm.io/gorm"
)

// GetPatientInfoByUserID returns patient info for a given user ID.
func GetPatientInfoByUserID(userID uint) (models.PatientInfo, error) {
	var info models.PatientInfo
	err := config.DB.Where("user_id = ?", userID).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.PatientInfo{}, errors.New("No patient information found for this user")
	}
	return info, err
}

// GetPatientInfoByDeviceId returns patient info for a given device ID.
func GetPatientInfoByDeviceId(deviceID string) (*models.PatientInfo, error) {
	var device models.Device
	if err := config.DB.Where("device_id = ?", deviceID).First(&device).Error; err != nil {
		return nil, err
	}

	var patientInfo models.PatientInfo
	if err := config.DB.Where("user_id = ?", device.UserID).First(&patientInfo).Error; err != nil {
		return nil, err
	}

	return &patientInfo, nil
}

// UpdatePatientInfoByUserID updates patient info for a given user ID.
func UpdatePatientInfoByUserID(userID uint, input *models.PatientInfo) (models.PatientInfo, error) {
	existingInfo, err := GetPatientInfoByUserID(userID)
	if err != nil {
		return models.PatientInfo{}, err
	}

	if err := config.DB.Model(&existingInfo).Updates(input).Error; err != nil {
		return models.PatientInfo{}, fmt.Errorf("Patient information could not be updated: %w", err)
	}

	return existingInfo, nil
}
