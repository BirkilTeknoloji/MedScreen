package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"go-backend/services"
	"go-backend/utils"
	"strings"

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

// GetPatientInfoDetailByQR, bir hastanın bilgilerinden belirli bir öğeyi (test veya reçete gibi) alır.
// Bu fonksiyon, hastanın belirli bir verisine bağlanan QR kodlarıyla kullanılmak üzere tasarlanmıştır.
func GetPatientInfoDetailByQR(c *fiber.Ctx) error {
	userID := c.Params("id")
	field := c.Params("field")
	itemID := c.Params("itemId")

	var patientInfo models.PatientInfo
	if err := config.DB.Where("user_id = ?", userID).First(&patientInfo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Verilen kullanıcı ID'si için hasta bilgisi bulunamadı",
		})
	}

	switch strings.ToLower(field) {
	case "appointments":
		for _, item := range patientInfo.Appointments {
			if item.ID == itemID {
				return c.JSON(item)
			}
		}
	case "diagnosis":
		for _, item := range patientInfo.Diagnosis {
			if item.ID == itemID {
				return c.JSON(item)
			}
		}
	case "prescriptions":
		for _, item := range patientInfo.Prescriptions {
			if item.ID == itemID {
				return c.JSON(item)
			}
		}
	case "notes":
		for _, item := range patientInfo.Notes {
			if item.ID == itemID {
				return c.JSON(item)
			}
		}
	case "tests":
		for _, item := range patientInfo.Tests {
			if item.ID == itemID {
				return c.JSON(item)
			}
		}
	case "allergies":
		for _, item := range patientInfo.Allergies {
			if item.ID == itemID {
				return c.JSON(item)
			}
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz alan belirtildi. Geçerli alanlar: appointments, diagnosis, prescriptions, notes, tests, allergies.",
		})
	}

	// Eğer buraya gelinirse, belirtilen alandaki listede öğe bulunamamıştır.
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Belirtilen alanda verilen ID ile öğe bulunamadı",
	})
}
