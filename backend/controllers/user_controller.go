package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tüm kullanıcıları getir
func GetUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// Yeni kullanıcı oluştur
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CreateUser(user)
	c.JSON(http.StatusCreated, result)
}

// Belirli kullanıcıyı getir
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Belirli kullanıcıyı sil
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı silinemedi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı silindi"})
}
