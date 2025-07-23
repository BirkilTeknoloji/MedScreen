package services

import (
	"go-backend/config"
	"go-backend/models"
)

// Tüm kullanıcıları getir
func GetAllUsers() []models.User {
	var users []models.User
	config.DB.Find(&users)
	return users
}

// Yeni kullanıcı oluştur
func CreateUser(user models.User) models.User {
	config.DB.Create(&user)
	return user
}

// ID ile kullanıcı getir
func GetUserByID(id string) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

// ID ile kullanıcı sil
func DeleteUser(id string) error {
	return config.DB.Delete(&models.User{}, id).Error
}
