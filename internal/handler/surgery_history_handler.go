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

// SurgeryHistoryHandler handles HTTP requests for surgery history operations
type SurgeryHistoryHandler struct {
	service service.SurgeryHistoryService
}

// NewSurgeryHistoryHandler creates a new SurgeryHistoryHandler instance
func NewSurgeryHistoryHandler(service service.SurgeryHistoryService) *SurgeryHistoryHandler {
	return &SurgeryHistoryHandler{service: service}
}

// CreateSurgeryHistory handles POST /api/v1/surgery-histories
func (h *SurgeryHistoryHandler) CreateSurgeryHistory(c *gin.Context) {
	var surgery models.SurgeryHistory
	if err := c.ShouldBindJSON(&surgery); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.CreateSurgeryHistory(&surgery); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create surgery history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "Surgery history created successfully", surgery)
}

// GetSurgeryHistory handles GET /api/v1/surgery-histories/:id
func (h *SurgeryHistoryHandler) GetSurgeryHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid surgery history ID", err)
		return
	}

	surgery, err := h.service.GetSurgeryHistory(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Surgery history not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Surgery history retrieved successfully", surgery)
}

// GetSurgeryHistories handles GET /api/v1/surgery-histories
func (h *SurgeryHistoryHandler) GetSurgeryHistories(c *gin.Context) {
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
	var startDate, endDate *time.Time

	if patientIDStr := c.Query("patient_id"); patientIDStr != "" {
		if id, err := strconv.ParseUint(patientIDStr, 10, 32); err == nil {
			uid := uint(id)
			patientID = &uid
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
	if patientID != nil || startDate != nil || endDate != nil {
		surgeries, total, err := h.service.GetSurgeryHistoriesByFilters(patientID, startDate, endDate, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve surgery histories", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Surgery histories retrieved successfully", surgeries, meta)
		return
	}

	// Default: get all surgery histories
	surgeries, total, err := h.service.GetSurgeryHistories(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve surgery histories", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Surgery histories retrieved successfully", surgeries, meta)
}

// UpdateSurgeryHistory handles PUT /api/v1/surgery-histories/:id
func (h *SurgeryHistoryHandler) UpdateSurgeryHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid surgery history ID", err)
		return
	}

	var surgery models.SurgeryHistory
	if err := c.ShouldBindJSON(&surgery); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateSurgeryHistory(uint(id), &surgery); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update surgery history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Surgery history updated successfully", surgery)
}

// DeleteSurgeryHistory handles DELETE /api/v1/surgery-histories/:id
func (h *SurgeryHistoryHandler) DeleteSurgeryHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid surgery history ID", err)
		return
	}

	if err := h.service.DeleteSurgeryHistory(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete surgery history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Surgery history deleted successfully", nil)
}
