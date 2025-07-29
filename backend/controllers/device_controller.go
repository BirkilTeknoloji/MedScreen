package controllers

import (
	"errors"
	"go-backend/services"

	"github.com/gofiber/fiber/v2"
)

// DeviceViewRequest, cihazdan gelen istek formatını tanımlar.
type DeviceViewRequest struct {
	ScannerCardID string `json:"scanner_card_id" validate:"required"`
	PatientUserID uint   `json:"patient_user_id" validate:"required"`
}

// QRMedicationRequest, QR kod ile ilaç sorgulama isteği formatını tanımlar.
type QRMedicationRequest struct {
	ScannerCardID string `json:"scanner_card_id" validate:"required"`
	QRData        string `json:"qr_data" validate:"required"`
}

// GetPatientDataForDevice, yatak başı cihazından gelen isteği işler.
// Kartı okutan kişinin rolüne göre hastanın bilgilerini gösterir veya erişimi reddeder.
func GetPatientDataForDevice(c *fiber.Ctx) error {
	var input DeviceViewRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek: scanner_card_id ve patient_user_id gereklidir."})
	}

	// Servis katmanını çağır
	patientData, err := services.GetPatientDataForView(input.ScannerCardID, input.PatientUserID)
	if err != nil {
		// Servisten dönen özel hataları yakala ve uygun HTTP durum kodlarını döndür
		switch {
		case errors.Is(err, services.ErrScannerNotFound), errors.Is(err, services.ErrPatientNotFound), errors.Is(err, services.ErrRequestedUserNotPatient):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, services.ErrPermissionDenied), errors.Is(err, services.ErrPatientSelfViewOnly):
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		default:
			// Diğer beklenmedik veritabanı veya sistem hataları için
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "İç sunucu hatası: " + err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(patientData)
}

// GetMedicationByQRCode, QR kod ile gelen isteği işler ve hastanın ilaç bilgilerini döndürür.
func GetMedicationByQRCode(c *fiber.Ctx) error {
	var input QRMedicationRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek: scanner_card_id ve qr_data gereklidir."})
	}

	// Servis katmanını çağır
	medicationData, err := services.GetMedicationByQRCode(input.ScannerCardID, input.QRData)
	if err != nil {
		// Servisten dönen özel hataları yakala
		switch {
		case errors.Is(err, services.ErrScannerNotFound), errors.Is(err, services.ErrPatientFromQRNotFound):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, services.ErrScannerNotAuthorized):
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "İç sunucu hatası: " + err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(medicationData)
}
