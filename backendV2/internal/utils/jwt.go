package utils

import (
	"time"

	"medscreen/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecretKey is the secret key used for signing JWT tokens
var jwtSecretKey []byte

// SetJWTSecretKey sets the JWT secret key from configuration
func SetJWTSecretKey() {
	// Load JWT secret key from configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	jwtSecretKey = []byte(cfg.JWT.SecretKey)
}

// GenerateJWT generates a JWT token for a given user ID and role
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 hours expiration
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// ParseJWT parses and validates a JWT token and returns the user ID and role
func ParseJWT(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, okID := claims["user_id"].(float64)
		role, okRole := claims["role"].(string)

		if okID && okRole {
			return uint(userIDFloat), role, nil
		}
	}
	return 0, "", jwt.ErrInvalidKey
}
