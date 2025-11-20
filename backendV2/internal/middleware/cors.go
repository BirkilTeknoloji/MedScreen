package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware creates a middleware handler for CORS with configurable allowed origins
func CORSMiddleware(allowedOrigins, allowedMethods, allowedHeaders []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if origin is allowed
		if origin != "" && isOriginAllowed(origin, allowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		} else if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
			// Allow all origins if wildcard is specified
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		// Set allowed methods
		if len(allowedMethods) > 0 {
			methods := ""
			for i, method := range allowedMethods {
				if i > 0 {
					methods += ", "
				}
				methods += method
			}
			c.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		}

		// Set allowed headers
		if len(allowedHeaders) > 0 {
			headers := ""
			for i, header := range allowedHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
		}

		// Handle preflight OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// isOriginAllowed checks if the given origin is in the allowed origins list
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if allowed == origin || allowed == "*" {
			return true
		}
	}
	return false
}
