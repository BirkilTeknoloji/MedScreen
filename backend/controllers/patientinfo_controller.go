package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetPatientInfoForUser, belirli bir kullanıcının hasta bilgilerini getirir.
func GetPatientInfoForUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	patientInfo, err := services.GetPatientInfoByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(patientInfo)
}

// UpdatePatientInfoForUser, belirli bir kullanıcının hasta bilgilerini günceller.
func UpdatePatientInfoForUser(c *fiber.Ctx) error {
	idParam := c.Params("id") // Bu User'ın ID'si olmalı
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	var input models.PatientInfo
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz veri formatı: " + err.Error()})
	}

	// Servis katmanını çağırarak güncelleme yap
	updatedInfo, err := services.UpdatePatientInfoByUserID(uint(userID), &input)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedInfo)
}
