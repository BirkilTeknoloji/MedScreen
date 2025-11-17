package handler

import (
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MedicalHistoryHandler handles HTTP requests for medical history operations
type MedicalHistoryHandler struct {
	service service.MedicalHistoryService
}

// NewMedicalHistoryHandler creates a new MedicalHistoryHandler instance
func NewMedicalHistoryHandler(service service.MedicalHistoryService) *MedicalHistoryHandler {
	return &MedicalHistoryHandler{service: service}
}

// CreateMedicalHistory handles POST /api/v1/medical-histories
func (h *MedicalHistoryHandler) CreateMedicalHistory(c *gin.Context) {
	var history models.MedicalHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.CreateMedicalHistory(&history); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create medical history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "Medical history created successfully", history)
}

// GetMedicalHistory handles GET /api/v1/medical-histories/:id
func (h *MedicalHistoryHandler) GetMedicalHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid medical history ID", err)
		return
	}

	history, err := h.service.GetMedicalHistory(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Medical history not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Medical history retrieved successfully", history)
}

// GetMedicalHistories handles GET /api/v1/medical-histories
func (h *MedicalHistoryHandler) GetMedicalHistories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse filter parameters
	var patientID *uint
	var status *models.MedicalHistoryStatus

	if patientIDStr := c.Query("patient_id"); patientIDStr != "" {
		if id, err := strconv.ParseUint(patientIDStr, 10, 32); err == nil {
			uid := uint(id)
			patientID = &uid
		}
	}

	if statusStr := c.Query("status"); statusStr != "" {
		s := models.MedicalHistoryStatus(statusStr)
		status = &s
	}

	// Check if any filters are applied
	if patientID != nil || status != nil {
		histories, total, err := h.service.GetMedicalHistoriesByFilters(patientID, status, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve medical histories", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Medical histories retrieved successfully", histories, meta)
		return
	}

	// Default: get all medical histories
	histories, total, err := h.service.GetMedicalHistories(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve medical histories", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Medical histories retrieved successfully", histories, meta)
}

// UpdateMedicalHistory handles PUT /api/v1/medical-histories/:id
func (h *MedicalHistoryHandler) UpdateMedicalHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid medical history ID", err)
		return
	}

	var history models.MedicalHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateMedicalHistory(uint(id), &history); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update medical history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Medical history updated successfully", history)
}

// DeleteMedicalHistory handles DELETE /api/v1/medical-histories/:id
func (h *MedicalHistoryHandler) DeleteMedicalHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid medical history ID", err)
		return
	}

	if err := h.service.DeleteMedicalHistory(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete medical history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Medical history deleted successfully", nil)
}
