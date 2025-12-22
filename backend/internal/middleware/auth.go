package middleware

import (
	"medscreen/internal/models"
	"medscreen/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware for JWT authentication
// For VEM 2.0 read-only system, this validates JWT tokens without audit logging
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}
		// Parse the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userID, role, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set("userID", userID)
		c.Set("userRole", role)

		// Note: Audit context removed for read-only VEM 2.0 system
		// No write operations are performed, so audit logging is not needed

		c.Next()
	}
}

// RoleMiddleware checks if the user has one of the allowed roles
// Uses PersonelGorevKodu from VEM 2.0 schema
func RoleMiddleware(allowedRoles ...models.PersonelGorevKodu) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleString, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User role not found"})
			return
		}

		userRole := models.PersonelGorevKodu(roleString.(string))
		for _, allowed := range allowedRoles {
			if userRole == allowed {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
	}
}
