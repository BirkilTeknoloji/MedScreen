package routes

import (
	"go-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rotaları gruplamak için (örneğin /api/v1)
	api := app.Group("/api/v1")

	users := api.Group("/users")
	users.Post("/", controllers.CreateUser)
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUserByID)
	users.Delete("/:id", controllers.DeleteUserByID)
	users.Get("/:id/qrcode", controllers.GetUserQRCode) // Yeni QR kod rotası
	users.Get("/card/:card_id", controllers.GetUserByCardID)

	// PatientInfo Rotaları (User'a bağlı)
	users.Get("/:id/patientinfo", controllers.GetPatientInfoForUser)
	users.Put("/:id/patientinfo", controllers.UpdatePatientInfoForUser)

	// Cihaz Rotaları
	device := api.Group("/device")
	device.Post("/view", controllers.GetPatientDataForDevice)
	device.Post("/medication-by-qr", controllers.GetMedicationByQRCode)

	// /api/devices altındaki route grubu
	deviceRoutes := api.Group("/devices")

	// Router tanımları
	deviceRoutes.Post("/register", controllers.Register)
	deviceRoutes.Get("/user/:userID", controllers.ListByUser)
	deviceRoutes.Get("/user/:userID/count", controllers.CountByUser)
}
