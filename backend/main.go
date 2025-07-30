package main

import (
	"go-backend/config"
	"go-backend/models"
	"go-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// .env + DB bağlantısını başlat
	config.InitDB()

	// Veritabanı tablolarını oluştur (otomatik migrate)
	config.DB.AutoMigrate(&models.User{}, &models.PatientInfo{}, &models.Device{})

	// Fiber app başlat
	app := fiber.New()

	// Gelen istekleri, durumu, gecikmeyi vb. loglamak için logger middleware'ini kullan.
	// Daha detaylı bir format için
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${latency} -${method} ${path}\n",
	}))

	// Rotaları tanımla
	routes.SetupRoutes(app)

	// Uygulamayı çalıştır
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Fiber app'i başlatırken hata oluştu: %v", err)
	}
}
