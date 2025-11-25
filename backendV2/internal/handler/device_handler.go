package handler

import (
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	deviceService service.DeviceService
}

func NewDeviceHandler(deviceService service.DeviceService) *DeviceHandler {
	return &DeviceHandler{deviceService: deviceService}
}

// RegisterDevice registers a new device
func (h *DeviceHandler) RegisterDevice(c *gin.Context) {
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid request payload", err)
		return
	}

	if err := h.deviceService.RegisterDevice(&device); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "REGISTRATION_FAILED", "Failed to register device", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "DEVICE_REGISTERED", "Device registered successfully", device)
}

// GetDeviceByMAC returns device info by MAC address
func (h *DeviceHandler) GetDeviceByMAC(c *gin.Context) {
	mac := c.Param("mac_address")
	if mac == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "MISSING_MAC", "MAC address is required", nil)
		return
	}

	device, err := h.deviceService.GetDeviceByMAC(mac)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "DEVICE_NOT_FOUND", "Device not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "DEVICE_RETRIEVED", "Device retrieved successfully", device)
}

// AssignPatient assigns a patient to a device
func (h *DeviceHandler) AssignPatient(c *gin.Context) {
	mac := c.Param("mac_address")
	if mac == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "MISSING_MAC", "MAC address is required", nil)
		return
	}

	var req struct {
		PatientID uint `json:"patient_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid request payload", err)
		return
	}

	if err := h.deviceService.AssignPatient(mac, req.PatientID); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "ASSIGNMENT_FAILED", "Failed to assign patient", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "PATIENT_ASSIGNED", "Patient assigned successfully", nil)
}

// UnassignPatient removes patient from device
func (h *DeviceHandler) UnassignPatient(c *gin.Context) {
	mac := c.Param("mac_address")
	if mac == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "MISSING_MAC", "MAC address is required", nil)
		return
	}

	if err := h.deviceService.UnassignPatient(mac); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "UNASSIGNMENT_FAILED", "Failed to unassign patient", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "PATIENT_UNASSIGNED", "Patient unassigned successfully", nil)
}

// GetDevices returns devices with optional filters
func (h *DeviceHandler) GetDevices(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Parse optional filters
	var roomNumber *string
	if rn := c.Query("room_number"); rn != "" {
		roomNumber = &rn
	}

	var patientID *uint
	if pidStr := c.Query("patient_id"); pidStr != "" {
		pid, err := strconv.ParseUint(pidStr, 10, 32)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, "INVALID_PATIENT_ID", "Invalid patient ID", err)
			return
		}
		pidUint := uint(pid)
		patientID = &pidUint
	}

	devices, total, err := h.deviceService.GetDevicesByFilters(roomNumber, patientID, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "FETCH_FAILED", "Failed to fetch devices", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "DEVICES_RETRIEVED", "Devices retrieved successfully", devices, meta)
}

// UpdateDevice updates device fields by MAC address
func (h *DeviceHandler) UpdateDevice(c *gin.Context) {
	mac := c.Param("mac_address")
	if mac == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "MISSING_MAC", "MAC address is required", nil)
		return
	}

	var req service.DeviceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid request payload", err)
		return
	}

	device, err := h.deviceService.UpdateDevice(mac, &req)
	if err != nil {
		if err.Error() == "device not found" {
			utils.SendErrorResponse(c, http.StatusNotFound, "DEVICE_NOT_FOUND", "Device not found", err)
			return
		}
		utils.SendErrorResponse(c, http.StatusInternalServerError, "UPDATE_FAILED", "Failed to update device", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "DEVICE_UPDATED", "Device updated successfully", device)
}

// DeleteDevice deletes a device by MAC address
func (h *DeviceHandler) DeleteDevice(c *gin.Context) {
	mac := c.Param("mac_address")
	if mac == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "MISSING_MAC", "MAC address is required", nil)
		return
	}

	if err := h.deviceService.DeleteDevice(mac); err != nil {
		if err.Error() == "device not found" {
			utils.SendErrorResponse(c, http.StatusNotFound, "DEVICE_NOT_FOUND", "Device not found", err)
			return
		}
		utils.SendErrorResponse(c, http.StatusInternalServerError, "DELETE_FAILED", "Failed to delete device", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "DEVICE_DELETED", "Device deleted successfully", nil)
}

// GetDeviceByID returns device info by database ID
func (h *DeviceHandler) GetDeviceByID(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "MISSING_ID", "Device ID is required", nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "INVALID_ID", "Invalid device ID", err)
		return
	}

	device, err := h.deviceService.GetDeviceByID(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "DEVICE_NOT_FOUND", "Device not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "DEVICE_RETRIEVED", "Device retrieved successfully", device)
}
