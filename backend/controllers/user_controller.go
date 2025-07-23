package controllers

import (
	"go-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input struct {
		Name   string `json:"Name" binding:"required"`
		Role   string `json:"Role" binding:"required"`
		CardID string `json:"CardID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tüm alanlar gereklidir"})
		return
	}

	user, err := services.CreateUser(input.Name, input.Role, input.CardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByCardID(c *gin.Context) {
	cardID := c.Param("card_id")

	user, err := services.GetUserByCardID(cardID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kart ile eşleşen kullanıcı bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, user)
}
