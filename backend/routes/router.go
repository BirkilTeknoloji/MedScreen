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
	users.Get("/card/:card_id", controllers.GetUserByCardID)
}
