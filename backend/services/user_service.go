package services

import (
	"errors"
	"go-backend/config"
	"go-backend/models"
)

func CreateUser(name, role, cardID string) (models.User, error) {
	// Role doğrulaması (güvenlik için)
	if role != "doctor" && role != "patient" {
		return models.User{}, errors.New("Geçersiz rol. 'doctor' veya 'patient' olmalı")
	}

	// Kart ID daha önce kullanılmış mı?
	var existing models.User
	if err := config.DB.Where("card_id = ?", cardID).First(&existing).Error; err == nil {
		return models.User{}, errors.New("Bu kart zaten başka bir kullanıcıya atanmış")
	}

	user := models.User{
		Name:   name,
		Role:   role,
		CardID: cardID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
