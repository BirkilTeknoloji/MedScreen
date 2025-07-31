package routes

import (
	"go-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	{
		users.Post("/", controllers.CreateUser)
		users.Get("/", controllers.GetAllUsers)
		users.Get("/:id", controllers.GetUserByID)
		users.Delete("/:id", controllers.DeleteUserByID)
		users.Get("/:id/qrcode", controllers.GetUserQRCode)
		users.Get("/card/:card_id", controllers.GetUserByCardID)

		// PatientInfo routes (User related)
		users.Get("/:id/patientinfo", controllers.GetPatientInfoForUser)
		users.Put("/:id/patientinfo", controllers.UpdatePatientInfoForUser)
		users.Get("/device/:deviceId/patientinfo", controllers.GetPatientInfoByDeviceId)
	}

	// Device routes
	device := api.Group("/device")
	{
		device.Post("/view", controllers.GetPatientDataForDevice)
		device.Post("/medication-by-qr", controllers.GetMedicationByQRCode)
	}

	// Device management routes
	deviceRoutes := api.Group("/devices")
	{
		deviceRoutes.Post("/register", controllers.Register)
		deviceRoutes.Get("/user/:userID", controllers.ListByUser)
		deviceRoutes.Get("/user/:userID/count", controllers.CountByUser)
	}
}
