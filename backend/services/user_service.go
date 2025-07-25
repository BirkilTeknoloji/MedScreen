package services

import (
	"errors"
	"fmt"
	"go-backend/config"
	"go-backend/models"
	"strings"
)

func CreateUser(user models.User, patientInfo *models.PatientInfo) (models.User, error) {
	role := strings.ToLower(user.Role)
	if role != "doctor" && role != "patient" {
		return models.User{}, errors.New("Geçersiz rol: sadece 'doctor' veya 'patient' olabilir")
	}

	// Gelen verileri ata
	newUser := models.User{
		Name:   user.Name,
		Role:   role,
		CardID: user.CardID,
	}

	// Transaction başlat
	tx := config.DB.Begin()
	if tx.Error != nil {
		return models.User{}, fmt.Errorf("transaction başlatılamadı: %w", tx.Error)
	}

	// 1. Kullanıcıyı oluştur
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		// Hata mesajını daha anlaşılır hale getirelim
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return models.User{}, errors.New("bu CardID veya başka bir benzersiz alan zaten kullanılıyor")
		}
		return models.User{}, fmt.Errorf("kullanıcı oluşturulamadı: %w", err)
	}

	// Eğer kullanıcı hasta ise ve hasta bilgisi varsa PatientInfo oluştur
	if role == "patient" && patientInfo != nil {
		patientInfo.UserID = newUser.ID // Oluşturulan User'ın ID'sini ata
		if err := tx.Create(patientInfo).Error; err != nil {
			tx.Rollback()
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return models.User{}, errors.New("bu TC Kimlik Numarası veya UserID zaten kullanılıyor")
			}
			return models.User{}, fmt.Errorf("hasta bilgisi oluşturulamadı: %w", err)
		}
		newUser.PatientInfo = *patientInfo
	}

	// Her şey yolundaysa transaction'ı onayla
	return newUser, tx.Commit().Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	// Preload ile ilişkili PatientInfo verisini de çekiyoruz
	if err := config.DB.Preload("PatientInfo").First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Kullanıcıyı ID ile sil
func DeleteUserByID(id uint) error {
	var user models.User

	// Kullanıcıyı önce bul
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("Kullanıcı bulunamadı")
	}

	// Sil (soft delete)
	// Not: İlişkili PatientInfo'yu da silmek isterseniz, burada transaction içinde ek bir silme işlemi gerekir.
	// Şimdilik sadece User siliniyor.
	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// CardID ile kullanıcı getir
func GetUserByCardID(cardID string) (models.User, error) {
	var user models.User
	if err := config.DB.Preload("PatientInfo").Where("card_id = ?", cardID).First(&user).Error; err != nil {
		return models.User{}, errors.New("Kart ile eşleşen kullanıcı bulunamadı")
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("PatientInfo").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
