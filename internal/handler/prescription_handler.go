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

// PrescriptionHandler handles HTTP requests for prescription operations
type PrescriptionHandler struct {
	service service.PrescriptionService
}

// NewPrescriptionHandler creates a new PrescriptionHandler instance
func NewPrescriptionHandler(service service.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{service: service}
}

// CreatePrescription handles POST /api/v1/prescriptions
func (h *PrescriptionHandler) CreatePrescription(c *gin.Context) {
	var prescription models.Prescription
	if err := c.ShouldBindJSON(&prescription); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.CreatePrescription(&prescription); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create prescription", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "Prescription created successfully", prescription)
}

// GetPrescription handles GET /api/v1/prescriptions/:id
func (h *PrescriptionHandler) GetPrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid prescription ID", err)
		return
	}

	prescription, err := h.service.GetPrescription(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Prescription not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Prescription retrieved successfully", prescription)
}

// GetPrescriptions handles GET /api/v1/prescriptions
func (h *PrescriptionHandler) GetPrescriptions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse filter parameters
	var patientID, doctorID *uint
	var status *models.PrescriptionStatus
	var startDate, endDate *time.Time

	if patientIDStr := c.Query("patient_id"); patientIDStr != "" {
		if id, err := strconv.ParseUint(patientIDStr, 10, 32); err == nil {
			uid := uint(id)
			patientID = &uid
		}
	}

	if doctorIDStr := c.Query("doctor_id"); doctorIDStr != "" {
		if id, err := strconv.ParseUint(doctorIDStr, 10, 32); err == nil {
			uid := uint(id)
			doctorID = &uid
		}
	}

	if statusStr := c.Query("status"); statusStr != "" {
		s := models.PrescriptionStatus(statusStr)
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
	if patientID != nil || doctorID != nil || status != nil || startDate != nil || endDate != nil {
		prescriptions, total, err := h.service.GetPrescriptionsByFilters(patientID, doctorID, status, startDate, endDate, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve prescriptions", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Prescriptions retrieved successfully", prescriptions, meta)
		return
	}

	// Default: get all prescriptions
	prescriptions, total, err := h.service.GetPrescriptions(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve prescriptions", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Prescriptions retrieved successfully", prescriptions, meta)
}

// UpdatePrescription handles PUT /api/v1/prescriptions/:id
func (h *PrescriptionHandler) UpdatePrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid prescription ID", err)
		return
	}

	var prescription models.Prescription
	if err := c.ShouldBindJSON(&prescription); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.UpdatePrescription(uint(id), &prescription); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update prescription", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Prescription updated successfully", prescription)
}

// DeletePrescription handles DELETE /api/v1/prescriptions/:id
func (h *PrescriptionHandler) DeletePrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid prescription ID", err)
		return
	}

	if err := h.service.DeletePrescription(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete prescription", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Prescription deleted successfully", nil)
}
