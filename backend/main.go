package main

import (
	"go-backend/config"
	"go-backend/models"
	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// .env + DB bağlantısını başlat
	config.InitDB()

	// Veritabanı tablolarını oluştur (otomatik migrate)
	config.DB.AutoMigrate(&models.User{})

	// Gin router başlat
	r := gin.Default()

	// Rotaları tanımla
	routes.SetupRoutes(r)

	// Uygulamayı çalıştır
	r.Run(":8080")
}
