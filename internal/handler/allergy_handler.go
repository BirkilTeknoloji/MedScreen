package handler

import (
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AllergyHandler handles HTTP requests for allergy operations
type AllergyHandler struct {
	service service.AllergyService
}

// NewAllergyHandler creates a new AllergyHandler instance
func NewAllergyHandler(service service.AllergyService) *AllergyHandler {
	return &AllergyHandler{service: service}
}

// CreateAllergy handles POST /api/v1/allergies
func (h *AllergyHandler) CreateAllergy(c *gin.Context) {
	var allergy models.Allergy
	if err := c.ShouldBindJSON(&allergy); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.CreateAllergy(&allergy); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create allergy", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "Allergy created successfully", allergy)
}

// GetAllergy handles GET /api/v1/allergies/:id
func (h *AllergyHandler) GetAllergy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid allergy ID", err)
		return
	}

	allergy, err := h.service.GetAllergy(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Allergy not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Allergy retrieved successfully", allergy)
}

// GetAllergies handles GET /api/v1/allergies
func (h *AllergyHandler) GetAllergies(c *gin.Context) {
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
	var severity *models.AllergySeverity

	if patientIDStr := c.Query("patient_id"); patientIDStr != "" {
		if id, err := strconv.ParseUint(patientIDStr, 10, 32); err == nil {
			uid := uint(id)
			patientID = &uid
		}
	}

	if severityStr := c.Query("severity"); severityStr != "" {
		s := models.AllergySeverity(severityStr)
		severity = &s
	}

	// Check if any filters are applied
	if patientID != nil || severity != nil {
		allergies, total, err := h.service.GetAllergiesByFilters(patientID, severity, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve allergies", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Allergies retrieved successfully", allergies, meta)
		return
	}

	// Default: get all allergies
	allergies, total, err := h.service.GetAllergies(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve allergies", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "Allergies retrieved successfully", allergies, meta)
}

// UpdateAllergy handles PUT /api/v1/allergies/:id
func (h *AllergyHandler) UpdateAllergy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid allergy ID", err)
		return
	}

	var allergy models.Allergy
	if err := c.ShouldBindJSON(&allergy); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateAllergy(uint(id), &allergy); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update allergy", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Allergy updated successfully", allergy)
}

// DeleteAllergy handles DELETE /api/v1/allergies/:id
func (h *AllergyHandler) DeleteAllergy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid allergy ID", err)
		return
	}

	if err := h.service.DeleteAllergy(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete allergy", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Allergy deleted successfully", nil)
}
