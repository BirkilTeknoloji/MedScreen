package services

import (
	"errors"
	"fmt"
	"go-backend/config"
	"go-backend/models"

	"gorm.io/gorm"
)

// GetPatientInfoByUserID, verilen kullanıcı ID'sine ait hasta bilgilerini getirir.
func GetPatientInfoByUserID(userID uint) (models.PatientInfo, error) {
	var info models.PatientInfo
	err := config.DB.Where("user_id = ?", userID).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.PatientInfo{}, errors.New("bu kullanıcıya ait hasta bilgisi bulunamadı")
	}
	return info, err
}

// GetPatientInfoByDeviceId, verilen cihaz ID'sine ait hasta bilgilerini getirir.
func GetPatientInfoByDeviceId(DeviceID string) (*models.PatientInfo, error) {
	var device models.Device
	if err := config.DB.Where("device_id = ?", DeviceID).First(&device).Error; err != nil {
		return nil, err
	}

	var patientInfo models.PatientInfo
	if err := config.DB.Where("user_id = ?", device.UserID).First(&patientInfo).Error; err != nil {
		return nil, err
	}

	return &patientInfo, nil
}

// UpdatePatientInfoByUserID, verilen kullanıcı ID'sine ait hasta bilgilerini günceller.
func UpdatePatientInfoByUserID(userID uint, input *models.PatientInfo) (models.PatientInfo, error) {
	// Önce mevcut kaydı bulalım
	existingInfo, err := GetPatientInfoByUserID(userID)
	if err != nil {
		return models.PatientInfo{}, err // Kayıt yoksa hata döner
	}

	// Gelen input'u mevcut kayıt üzerine işle (Model'deki ID, UserID gibi alanları korur)
	if err := config.DB.Model(&existingInfo).Updates(input).Error; err != nil {
		return models.PatientInfo{}, fmt.Errorf("hasta bilgisi güncellenemedi: %w", err)
	}

	return existingInfo, nil
}
