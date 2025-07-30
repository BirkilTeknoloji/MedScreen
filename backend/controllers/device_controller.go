package controllers

import (
	"errors"
	"go-backend/config"
	"go-backend/services"
	"strconv"

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

// RegisterDeviceRequest, cihaz kaydı için gelen isteğin yapısını tanımlar.
type RegisterDeviceRequest struct {
	DeviceID string `json:"device_id" validate:"required"`
	UserID   uint   `json:"user_id" validate:"required"`
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

type DeviceController struct {
	Service *services.DeviceService
}

func NewDeviceController(s *services.DeviceService) *DeviceController {
	return &DeviceController{Service: s}
}

func (c *DeviceController) Register(ctx *fiber.Ctx) error {
	var req struct {
		DeviceID string `json:"device_id"`
		UserID   uint   `json:"user_id"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	d, err := c.Service.RegisterDevice(req.DeviceID, req.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(d)
}

func (c *DeviceController) ListByUser(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("userID"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid userID"})
	}
	devices, err := c.Service.ListUserDevices(uint(uid))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(devices)
}

func (c *DeviceController) CountByUser(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("userID"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid userID"})
	}
	cnt, err := c.Service.CountUserDevices(uint(uid))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"count": cnt})
}

// Register, yeni bir cihazı kaydeder veya mevcutsa günceller.
// @Router /api/v1/devices/register [post]
func Register(c *fiber.Ctx) error {
	// 1. Gelen JSON isteğini ayrıştır
	var req RegisterDeviceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// 2. Gerekli alanların dolu olduğunu kontrol et
	if req.DeviceID == "" || req.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "device_id and user_id are required",
		})
	}

	// 3. İlgili servisi çağır
	deviceService := services.NewDeviceService(config.DB)
	device, err := deviceService.RegisterDevice(req.DeviceID, req.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 4. Başarılı yanıtı döndür
	return c.Status(fiber.StatusCreated).JSON(device)
}

// ListByUser, belirli bir kullanıcıya ait tüm cihazları listeler.
// @Router /api/v1/devices/user/:userID [get]
func ListByUser(c *fiber.Ctx) error {
	// 1. URL'den userID'yi al ve sayıya çevir
	userIDStr := c.Params("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	// 2. Servisi çağır
	deviceService := services.NewDeviceService(config.DB)
	devices, err := deviceService.ListUserDevices(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(devices)
}

// CountByUser, belirli bir kullanıcıya ait cihaz sayısını verir.
// @Router /api/v1/devices/user/:userID/count [get]
func CountByUser(c *fiber.Ctx) error {
	userIDStr := c.Params("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	deviceService := services.NewDeviceService(config.DB)
	count, err := deviceService.CountUserDevices(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"count": count})
}
