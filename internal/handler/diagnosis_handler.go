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

// DiagnosisHandler handles HTTP requests for diagnosis operations
type DiagnosisHandler struct {
	service service.DiagnosisService
}

// NewDiagnosisHandler creates a new DiagnosisHandler instance
func NewDiagnosisHandler(service service.DiagnosisService) *DiagnosisHandler {
	return &DiagnosisHandler{service: service}
}

// CreateDiagnosis handles POST /api/v1/diagnoses
func (h *DiagnosisHandler) CreateDiagnosis(c *gin.Context) {
	var diagnosis models.Diagnosis
	if err := c.ShouldBindJSON(&diagnosis); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.CreateDiagnosis(&diagnosis); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_DIAGNOSIS_CREATE_FAILED, "Failed to create diagnosis", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, utils.SUCCESS_DIAGNOSIS_CREATED, "Diagnosis created successfully", diagnosis)
}

// GetDiagnosis handles GET /api/v1/diagnoses/:id
func (h *DiagnosisHandler) GetDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_DIAGNOSIS_ID, "Invalid diagnosis ID", err)
		return
	}

	diagnosis, err := h.service.GetDiagnosis(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, utils.ERROR_DIAGNOSIS_NOT_FOUND, "Diagnosis not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_DIAGNOSIS_RETRIEVED, "Diagnosis retrieved successfully", diagnosis)
}

// GetDiagnoses handles GET /api/v1/diagnoses
func (h *DiagnosisHandler) GetDiagnoses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse filter parameters
	var patientID, doctorID, appointmentID *uint
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
	if patientID != nil || doctorID != nil || appointmentID != nil || startDate != nil || endDate != nil {
		diagnoses, total, err := h.service.GetDiagnosesByFilters(patientID, doctorID, appointmentID, startDate, endDate, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_INTERNAL_SERVER, "Failed to retrieve diagnoses", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, utils.SUCCESS_DIAGNOSES_RETRIEVED, "Diagnoses retrieved successfully", diagnoses, meta)
		return
	}

	// Default: get all diagnoses
	diagnoses, total, err := h.service.GetDiagnoses(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_INTERNAL_SERVER, "Failed to retrieve diagnoses", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, utils.SUCCESS_DIAGNOSES_RETRIEVED, "Diagnoses retrieved successfully", diagnoses, meta)
}

// UpdateDiagnosis handles PUT /api/v1/diagnoses/:id
func (h *DiagnosisHandler) UpdateDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_DIAGNOSIS_ID, "Invalid diagnosis ID", err)
		return
	}

	var diagnosis models.Diagnosis
	if err := c.ShouldBindJSON(&diagnosis); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateDiagnosis(uint(id), &diagnosis); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_DIAGNOSIS_UPDATE_FAILED, "Failed to update diagnosis", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_DIAGNOSIS_UPDATED, "Diagnosis updated successfully", diagnosis)
}

// DeleteDiagnosis handles DELETE /api/v1/diagnoses/:id
func (h *DiagnosisHandler) DeleteDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_DIAGNOSIS_ID, "Invalid diagnosis ID", err)
		return
	}

	if err := h.service.DeleteDiagnosis(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_DIAGNOSIS_DELETE_FAILED, "Failed to delete diagnosis", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_DIAGNOSIS_DELETED, "Diagnosis deleted successfully", nil)
}
