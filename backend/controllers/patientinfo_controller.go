package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"go-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// GetPatientInfoForUser, retrieves patient information for a specific user.
func GetPatientInfoForUser(c *fiber.Ctx) error {
	userID, err := utils.ParseUserID(c, "id")
	if err != nil {
		return nil
	}
	patientInfo, err := services.GetPatientInfoByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(patientInfo)
}

// UpdatePatientInfoForUser, updates patient information for a specific user.
func UpdatePatientInfoForUser(c *fiber.Ctx) error {
	userID, err := utils.ParseUserID(c, "id")
	if err != nil {
		return nil
	}
	var input models.PatientInfo
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid data format: " + err.Error()})
	}
	updatedInfo, err := services.UpdatePatientInfoByUserID(userID, &input)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(updatedInfo)
}

// GetPatientInfoByDeviceId, retrieves patient information by device ID.
func GetPatientInfoByDeviceId(c *fiber.Ctx) error {
	deviceId := c.Params("deviceId")
	if deviceId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Device ID is required"})
	}
	patientInfo, err := services.GetPatientInfoByDeviceId(deviceId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(patientInfo)
}
