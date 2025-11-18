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

// MedicalTestHandler handles HTTP requests for medical test operations
type MedicalTestHandler struct {
	service service.MedicalTestService
}

// NewMedicalTestHandler creates a new MedicalTestHandler instance
func NewMedicalTestHandler(service service.MedicalTestService) *MedicalTestHandler {
	return &MedicalTestHandler{service: service}
}

// CreateMedicalTest handles POST /api/v1/medical-tests
func (h *MedicalTestHandler) CreateMedicalTest(c *gin.Context) {
	var test models.MedicalTest
	if err := c.ShouldBindJSON(&test); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.CreateMedicalTest(&test); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_MEDICAL_TEST_CREATE_FAILED, "Failed to create medical test", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, utils.SUCCESS_MEDICAL_TEST_CREATED, "Medical test created successfully", test)
}

// GetMedicalTest handles GET /api/v1/medical-tests/:id
func (h *MedicalTestHandler) GetMedicalTest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_MEDICAL_TEST_ID, "Invalid medical test ID", err)
		return
	}

	test, err := h.service.GetMedicalTest(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, utils.ERROR_MEDICAL_TEST_NOT_FOUND, "Medical test not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_MEDICAL_TEST_RETRIEVED, "Medical test retrieved successfully", test)
}

// GetMedicalTests handles GET /api/v1/medical-tests
func (h *MedicalTestHandler) GetMedicalTests(c *gin.Context) {
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
	var testType *models.TestType
	var status *models.TestStatus
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

	if testTypeStr := c.Query("test_type"); testTypeStr != "" {
		tt := models.TestType(testTypeStr)
		testType = &tt
	}

	if statusStr := c.Query("status"); statusStr != "" {
		s := models.TestStatus(statusStr)
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
	if patientID != nil || doctorID != nil || testType != nil || status != nil || startDate != nil || endDate != nil {
		tests, total, err := h.service.GetMedicalTestsByFilters(patientID, doctorID, testType, status, startDate, endDate, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_INTERNAL_SERVER, "Failed to retrieve medical tests", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, utils.SUCCESS_MEDICAL_TESTS_RETRIEVED, "Medical tests retrieved successfully", tests, meta)
		return
	}

	// Default: get all medical tests
	tests, total, err := h.service.GetMedicalTests(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_INTERNAL_SERVER, "Failed to retrieve medical tests", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, utils.SUCCESS_MEDICAL_TESTS_RETRIEVED, "Medical tests retrieved successfully", tests, meta)
}

// UpdateMedicalTest handles PUT /api/v1/medical-tests/:id
func (h *MedicalTestHandler) UpdateMedicalTest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_MEDICAL_TEST_ID, "Invalid medical test ID", err)
		return
	}

	var test models.MedicalTest
	if err := c.ShouldBindJSON(&test); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateMedicalTest(uint(id), &test); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_MEDICAL_TEST_UPDATE_FAILED, "Failed to update medical test", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_MEDICAL_TEST_UPDATED, "Medical test updated successfully", test)
}

// DeleteMedicalTest handles DELETE /api/v1/medical-tests/:id
func (h *MedicalTestHandler) DeleteMedicalTest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ERROR_INVALID_MEDICAL_TEST_ID, "Invalid medical test ID", err)
		return
	}

	if err := h.service.DeleteMedicalTest(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ERROR_MEDICAL_TEST_DELETE_FAILED, "Failed to delete medical test", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, utils.SUCCESS_MEDICAL_TEST_DELETED, "Medical test deleted successfully", nil)
}
