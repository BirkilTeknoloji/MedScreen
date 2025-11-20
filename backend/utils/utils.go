package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParseUserID parses user ID from params and returns error response if invalid.
func ParseUserID(c *fiber.Ctx, param string) (uint, error) {
	idStr := c.Params(param)
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
		return 0, err
	}
	return uint(id), nil
}
