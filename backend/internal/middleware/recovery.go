package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"medscreen/internal/constants"
	"medscreen/internal/utils"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware creates a middleware handler that recovers from panics
// and returns a 500 Internal Server Error response
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log panic details with stack trace for debugging
				log.Printf("PANIC RECOVERED: %v\n", err)
				log.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
				log.Printf("Client IP: %s\n", c.ClientIP())
				log.Printf("Stack trace:\n%s\n", string(debug.Stack()))

				// Return standardized error response
				utils.SendErrorResponse(
					c,
					http.StatusInternalServerError,
					constants.ERROR_INTERNAL_SERVER,
					"Internal server error",
					fmt.Errorf("%v", err),
				)

				// Abort the request
				c.Abort()
			}
		}()

		c.Next()
	}
}
