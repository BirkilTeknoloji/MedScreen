package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QRHandler handles HTTP requests for QR code operations
type QRHandler struct {
	qrService     service.QRService
	deviceService service.DeviceService
}

// NewQRHandler creates a new QRHandler instance
func NewQRHandler(qrService service.QRService, deviceService service.DeviceService) *QRHandler {
	return &QRHandler{
		qrService:     qrService,
		deviceService: deviceService,
	}
}

// GeneratePatientQRRequest represents the request for generating a patient assignment QR
type GeneratePatientQRRequest struct {
	ExpiryHours int `json:"expiry_hours" binding:"required,min=1,max=168"` // Max 7 days
}

// GeneratePrescriptionQRRequest represents the request for generating a prescription info QR
type GeneratePrescriptionQRRequest struct {
	ExpiryHours int `json:"expiry_hours" binding:"required,min=1,max=720"` // Max 30 days
}

// ScanPatientQRRequest represents the request for scanning a patient QR code
type ScanPatientQRRequest struct {
	Token string `json:"token" binding:"required"`
}

// QRResponse represents the QR code generation response
type QRResponse struct {
	Token   string `json:"token"`
	QRImage string `json:"qr_image"` // Base64-encoded PNG
	Type    string `json:"type"`
	Expires string `json:"expires_at"`
}

// GeneratePatientQR handles POST /api/v1/patients/:id/generate-qr
func (h *QRHandler) GeneratePatientQR(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid patient ID", err)
		return
	}

	var request GeneratePatientQRRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	token, qrImage, err := h.qrService.GeneratePatientAssignmentQR(uint(patientID), request.ExpiryHours)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to generate QR code", err)
		return
	}

	response := QRResponse{
		Token:   token,
		QRImage: qrImage,
		Type:    "patient_assignment",
		Expires: "",
	}

	utils.SendSuccessResponse(c, http.StatusCreated, constants.SUCCESS_OPERATION_COMPLETED, "QR code generated successfully", response)
}

// ScanPatientQR handles POST /api/v1/devices/:mac_address/scan-patient-qr
func (h *QRHandler) ScanPatientQR(c *gin.Context) {
	macAddress := c.Param("mac_address")
	if macAddress == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "MAC address is required", nil)
		return
	}

	var request ScanPatientQRRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	err := h.qrService.AssignPatientToDevice(request.Token, macAddress)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INTERNAL_SERVER, err.Error(), err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_OPERATION_COMPLETED, "Patient assigned to device successfully", nil)
}

// GeneratePrescriptionQR handles POST /api/v1/devices/:mac_address/generate-prescription-qr/:patient_id
func (h *QRHandler) GeneratePrescriptionQR(c *gin.Context) {
	macAddress := c.Param("mac_address")
	if macAddress == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "MAC address is required", nil)
		return
	}

	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid patient ID", err)
		return
	}

	var request GeneratePrescriptionQRRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	// Get device by MAC address
	device, err := h.deviceService.GetDeviceByMAC(macAddress)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_INTERNAL_SERVER, "Device not found", err)
		return
	}

	token, qrImage, err := h.qrService.GeneratePrescriptionInfoQR(uint(patientID), device.ID, request.ExpiryHours)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to generate QR code", err)
		return
	}

	response := QRResponse{
		Token:   token,
		QRImage: qrImage,
		Type:    "prescription_info",
		Expires: "",
	}

	utils.SendSuccessResponse(c, http.StatusCreated, constants.SUCCESS_OPERATION_COMPLETED, "QR code generated successfully", response)
}

// ValidateQRToken handles GET /api/v1/qr-tokens/:token/validate
func (h *QRHandler) ValidateQRToken(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Token is required", nil)
		return
	}

	qrToken, err := h.qrService.ValidateToken(token)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INTERNAL_SERVER, err.Error(), err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_OPERATION_COMPLETED, "Token validated successfully", qrToken)
}
