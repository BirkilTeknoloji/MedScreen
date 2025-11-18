package handler

import (
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AppointmentHandler handles HTTP requests for appointment operations
type AppointmentHandler struct {
	service service.AppointmentService
}

// NewAppointmentHandler creates a new AppointmentHandler instance
func NewAppointmentHandler(service service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

// CreateAppointment handles POST /api/v1/appointments
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.CreateAppointment(&appointment); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_APPOINTMENT_CREATE_FAILED, "Failed to create appointment", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, utils.SUCCESS_APPOINTMENT_CREATED, "Appointment created successfully", appointment)
}

// GetAppointment handles GET /api/v1/appointments/:id
func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_APPOINTMENT_ID, "Invalid appointment ID", err)
		return
	}

	appointment, err := h.service.GetAppointment(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, utils.ERROR_APPOINTMENT_NOT_FOUND, "Appointment not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_APPOINTMENT_RETRIEVED, "Appointment retrieved successfully", appointment)
}

// GetAppointments handles GET /api/v1/appointments
func (h *AppointmentHandler) GetAppointments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse filter parameters
	var doctorID, patientID *uint
	var status *models.AppointmentStatus
	var startDate, endDate *time.Time

	if doctorIDStr := c.Query("doctor_id"); doctorIDStr != "" {
		if id, err := strconv.ParseUint(doctorIDStr, 10, 32); err == nil {
			uid := uint(id)
			doctorID = &uid
		}
	}

	if patientIDStr := c.Query("patient_id"); patientIDStr != "" {
		if id, err := strconv.ParseUint(patientIDStr, 10, 32); err == nil {
			uid := uint(id)
			patientID = &uid
		}
	}

	if statusStr := c.Query("status"); statusStr != "" {
		s := models.AppointmentStatus(statusStr)
		status = &s
	}

	if startDateStr := c.Query("start_date"); startDateStr != "" {
		if t, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &t
		}
	}

	if endDateStr := c.Query("end_date"); endDateStr != "" {
		if t, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &t
		}
	}

	// Check if any filters are applied
	if doctorID != nil || patientID != nil || status != nil || startDate != nil || endDate != nil {
		appointments, total, err := h.service.GetAppointmentsByFilters(doctorID, patientID, status, startDate, endDate, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, utils.SUCCESS_APPOINTMENTS_RETRIEVED, "Appointments retrieved successfully", appointments, meta)
		return
	}

	// Default: get all appointments
	appointments, total, err := h.service.GetAppointments(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, utils.SUCCESS_APPOINTMENTS_RETRIEVED, "Appointments retrieved successfully", appointments, meta)
}

// UpdateAppointment handles PUT /api/v1/appointments/:id
func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_APPOINTMENT_ID, "Invalid appointment ID", err)
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateAppointment(uint(id), &appointment); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_APPOINTMENT_UPDATE_FAILED, "Failed to update appointment", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_APPOINTMENT_UPDATED, "Appointment updated successfully", appointment)
}

// DeleteAppointment handles DELETE /api/v1/appointments/:id
func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_APPOINTMENT_ID, "Invalid appointment ID", err)
		return
	}

	if err := h.service.DeleteAppointment(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_APPOINTMENT_DELETE_FAILED, "Failed to delete appointment", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_APPOINTMENT_DELETED, "Appointment deleted successfully", nil)
}
