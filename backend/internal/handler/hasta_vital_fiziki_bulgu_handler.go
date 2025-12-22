package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// HastaVitalFizikiBulguHandler handles HTTP requests for vital signs operations (read-only)
type HastaVitalFizikiBulguHandler struct {
	service service.HastaVitalFizikiBulguService
}

// NewHastaVitalFizikiBulguHandler creates a new HastaVitalFizikiBulguHandler instance
func NewHastaVitalFizikiBulguHandler(service service.HastaVitalFizikiBulguService) *HastaVitalFizikiBulguHandler {
	return &HastaVitalFizikiBulguHandler{service: service}
}

// GetByKodu handles GET /api/v1/vital-bulgu/:kodu
func (h *HastaVitalFizikiBulguHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_VITAL_BULGU_KODU, "Vital sign code is required", nil)
		return
	}

	bulgu, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_VITAL_BULGU_NOT_FOUND, "Vital sign not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_VITAL_BULGU_RETRIEVED, "Vital sign retrieved successfully", bulgu)
}

// GetByBasvuru handles GET /api/v1/vital-bulgu/basvuru/:basvuru_kodu
func (h *HastaVitalFizikiBulguHandler) GetByBasvuru(c *gin.Context) {
	basvuruKodu := c.Param("basvuru_kodu")
	if basvuruKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_HASTA_BASVURU_KODU, "Visit code is required", nil)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	bulgular, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve vital signs", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_VITAL_BULGULAR_RETRIEVED, "Vital signs retrieved successfully", bulgular, meta)
}

// GetByDateRange handles GET /api/v1/vital-bulgu/date-range
func (h *HastaVitalFizikiBulguHandler) GetByDateRange(c *gin.Context) {
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	if startStr == "" || endStr == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_DATE_RANGE, "Start date and end date are required", nil)
		return
	}

	startDate, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_DATE_RANGE, "Invalid start date format (use YYYY-MM-DD)", err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_DATE_RANGE, "Invalid end date format (use YYYY-MM-DD)", err)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	bulgular, total, err := h.service.GetByDateRange(startDate, endDate, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve vital signs", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_VITAL_BULGULAR_RETRIEVED, "Vital signs retrieved successfully", bulgular, meta)
}
