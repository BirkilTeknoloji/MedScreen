package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{

		v1.POST("/users", controllers.CreateUser)
		v1.GET("/users/:id", controllers.GetUserByID)

	}

}
