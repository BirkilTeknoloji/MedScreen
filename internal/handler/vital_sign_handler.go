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

// VitalSignHandler handles HTTP requests for vital sign operations
type VitalSignHandler struct {
	service service.VitalSignService
}

// NewVitalSignHandler creates a new VitalSignHandler instance
func NewVitalSignHandler(service service.VitalSignService) *VitalSignHandler {
	return &VitalSignHandler{service: service}
}

// CreateVitalSign handles POST /api/v1/vital-signs
func (h *VitalSignHandler) CreateVitalSign(c *gin.Context) {
	var vitalSign models.VitalSign
	if err := c.ShouldBindJSON(&vitalSign); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.CreateVitalSign(&vitalSign); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create vital sign", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "Vital sign created successfully", vitalSign)
}

// GetVitalSign handles GET /api/v1/vital-signs/:id
func (h *VitalSignHandler) GetVitalSign(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid vital sign ID", err)
		return
	}

	vitalSign, err := h.service.GetVitalSign(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Vital sign not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Vital sign retrieved successfully", vitalSign)
}

// GetVitalSigns handles GET /api/v1/vital-signs
func (h *VitalSignHandler) GetVitalSigns(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse filter parameters
	var patientID, appointmentID *uint
	var startDate, endDate *time.Time

	if patientIDStr := c.Query("patient_id"); patientIDStr != "" {
		if id, err := strconv.ParseUint(patientIDStr, 10, 32); err == nil {
			uid := uint(id)
			patientID = &uid
		}
	}

	if appointmentIDStr := c.Query("appointment_id"); appointmentIDStr != "" {
		if id, err := strconv.ParseUint(appointmentIDStr, 10, 32); err == nil {
			uid := uint(id)
			appointmentID = &uid
		}
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
	if patientID != nil || appointmentID != nil || startDate != nil || endDate != nil {
		vitalSigns, total, err := h.service.GetVitalSignsByFilters(patientID, appointmentID, startDate, endDate, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve vital signs", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Vital signs retrieved successfully", vitalSigns, meta)
		return
	}

	// Default: get all vital signs
	vitalSigns, total, err := h.service.GetVitalSigns(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve vital signs", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Vital signs retrieved successfully", vitalSigns, meta)
}

// UpdateVitalSign handles PUT /api/v1/vital-signs/:id
func (h *VitalSignHandler) UpdateVitalSign(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid vital sign ID", err)
		return
	}

	var vitalSign models.VitalSign
	if err := c.ShouldBindJSON(&vitalSign); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateVitalSign(uint(id), &vitalSign); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update vital sign", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Vital sign updated successfully", vitalSign)
}

// DeleteVitalSign handles DELETE /api/v1/vital-signs/:id
func (h *VitalSignHandler) DeleteVitalSign(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid vital sign ID", err)
		return
	}

	if err := h.service.DeleteVitalSign(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete vital sign", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Vital sign deleted successfully", nil)
}
