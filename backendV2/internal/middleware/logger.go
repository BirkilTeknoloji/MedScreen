package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware creates a middleware handler that logs HTTP requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Get status code
		status := c.Writer.Status()

		// Build log message
		if query != "" {
			path = path + "?" + query
		}

		log.Printf("[%s] %s %s %d %v %s",
			method,
			path,
			c.ClientIP(),
			status,
			duration,
			c.Errors.String(),
		)
	}
}
