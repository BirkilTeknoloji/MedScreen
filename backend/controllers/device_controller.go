package controllers

import (
	"errors"
	"go-backend/config"
	"go-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeviceViewRequest defines the request format from the device.
type DeviceViewRequest struct {
	ScannerCardID string `json:"scanner_card_id" validate:"required"`
	PatientUserID uint   `json:"patient_user_id" validate:"required"`
}

// QRMedicationRequest, defines the format of the drug query request via QR code.
type QRMedicationRequest struct {
	ScannerCardID string `json:"scanner_card_id" validate:"required"`
	QRData        string `json:"qr_data" validate:"required"`
}

// RegisterDeviceRequest, defines the structure of the request coming for device registration.
type RegisterDeviceRequest struct {
	DeviceID string `json:"device_id" validate:"required"`
	UserID   uint   `json:"user_id" validate:"required"`
}

type DeviceController struct {
	Service *services.DeviceService
}

func NewDeviceController(s *services.DeviceService) *DeviceController {
	return &DeviceController{Service: s}
}

func (c *DeviceController) Register(ctx *fiber.Ctx) error {
	var req RegisterDeviceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if req.DeviceID == "" || req.UserID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "device_id and user_id are required"})
	}
	d, err := c.Service.RegisterDevice(req.DeviceID, req.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(d)
}

func (c *DeviceController) ListByUser(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("userID"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}
	devices, err := c.Service.ListUserDevices(uint(uid))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(devices)
}

func (c *DeviceController) CountByUser(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("userID"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}
	cnt, err := c.Service.CountUserDevices(uint(uid))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"count": cnt})
}

// GetPatientDataForDevice, processes the request from the bedside device.
// It displays the patient's information or denies access, depending on the role of the person who scans the card.
func GetPatientDataForDevice(c *fiber.Ctx) error {
	var input DeviceViewRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek: scanner_card_id ve patient_user_id gereklidir."})
	}

	// Invoke the service layer
	patientData, err := services.GetPatientDataForView(input.ScannerCardID, input.PatientUserID)
	if err != nil {
		// Catch custom errors returned from the service and return appropriate HTTP status codes
		switch {
		case errors.Is(err, services.ErrScannerNotFound), errors.Is(err, services.ErrPatientNotFound), errors.Is(err, services.ErrRequestedUserNotPatient):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, services.ErrPermissionDenied), errors.Is(err, services.ErrPatientSelfViewOnly):
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		default:
			// For other unexpected database or system errors
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error: " + err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(patientData)
}

// GetMedicationByQRCode, it processes the request received via QR code and returns the patient's medication information.
func GetMedicationByQRCode(c *fiber.Ctx) error {
	var input QRMedicationRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request: scanner_card_id and qr_data are required."})
	}

	// Invoke the service layer
	medicationData, err := services.GetMedicationByQRCode(input.ScannerCardID, input.QRData)
	if err != nil {
		// Catch custom errors returned from the service
		switch {
		case errors.Is(err, services.ErrScannerNotFound), errors.Is(err, services.ErrPatientFromQRNotFound):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, services.ErrScannerNotAuthorized):
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error: " + err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(medicationData)
}

// Function-based handlers delegate to struct-based controller for consistency

// Register, registers a new device or updates an existing one.
// @Router /api/v1/devices/register [post]
func Register(c *fiber.Ctx) error {
	deviceService := services.NewDeviceService(config.DB)
	controller := NewDeviceController(deviceService)
	return controller.Register(c)
}

// ListByUser, lists all devices belonging to a specific user.
// @Router /api/v1/devices/user/:userID [get]
func ListByUser(c *fiber.Ctx) error {
	deviceService := services.NewDeviceService(config.DB)
	controller := NewDeviceController(deviceService)
	return controller.ListByUser(c)
}

// CountByUser, returns the number of devices belonging to a specific user.
// @Router /api/v1/devices/user/:userID/count [get]
func CountByUser(c *fiber.Ctx) error {
	deviceService := services.NewDeviceService(config.DB)
	controller := NewDeviceController(deviceService)
	return controller.CountByUser(c)
}
