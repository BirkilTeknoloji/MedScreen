package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PatientHandler handles HTTP requests for patient operations
type PatientHandler struct {
	service service.PatientService
}

// NewPatientHandler creates a new PatientHandler instance
func NewPatientHandler(service service.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

// CreatePatient handles POST /api/v1/patients
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.CreatePatient(&patient); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_PATIENT_CREATE_FAILED, "Failed to create patient", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, constants.SUCCESS_PATIENT_CREATED, "Patient created successfully", patient)
}

// GetPatient handles GET /api/v1/patients/:id
func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PATIENT_ID, "Invalid patient ID", err)
		return
	}

	patient, err := h.service.GetPatient(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_PATIENT_NOT_FOUND, "Patient not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_PATIENT_RETRIEVED, "Patient retrieved successfully", patient)
}

// GetPatients handles GET /api/v1/patients
func (h *PatientHandler) GetPatients(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Check for search by name
	if name := c.Query("name"); name != "" {
		patients, total, err := h.service.SearchPatientsByName(name, page, limit)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_PATIENT_SEARCH_FAILED, "Failed to search patients", err)
			return
		}

		meta := utils.CalculateMeta(page, limit, total)
		utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_PATIENTS_RETRIEVED, "Patients retrieved successfully", patients, meta)
		return
	}

	// Check for search by TC number
	if tcNumber := c.Query("tc_number"); tcNumber != "" {
		patient, err := h.service.GetPatientByTCNumber(tcNumber)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_PATIENT_NOT_FOUND, "Patient not found", err)
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_PATIENT_RETRIEVED, "Patient retrieved successfully", patient)
		return
	}

	// Default: get all patients
	patients, total, err := h.service.GetPatients(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve patients", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_PATIENTS_RETRIEVED, "Patients retrieved successfully", patients, meta)
}

// UpdatePatient handles PUT /api/v1/patients/:id
func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PATIENT_ID, "Invalid patient ID", err)
		return
	}

	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.UpdatePatient(uint(id), &patient); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_PATIENT_UPDATE_FAILED, "Failed to update patient", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_PATIENT_UPDATED, "Patient updated successfully", patient)
}

// DeletePatient handles DELETE /api/v1/patients/:id
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PATIENT_ID, "Invalid patient ID", err)
		return
	}

	if err := h.service.DeletePatient(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_PATIENT_DELETE_FAILED, "Failed to delete patient", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_PATIENT_DELETED, "Patient deleted successfully", nil)
}

// GetPatientMedicalHistory handles GET /api/v1/patients/:id/medical-history
func (h *PatientHandler) GetPatientMedicalHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PATIENT_ID, "Invalid patient ID", err)
		return
	}

	history, err := h.service.GetPatientMedicalHistory(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve patient medical history", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_PATIENT_MEDICAL_HISTORY_RETRIEVED, "Patient medical history retrieved successfully", history)
}
