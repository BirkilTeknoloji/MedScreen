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

// RandevuHandler handles HTTP requests for appointment operations (read-only)
type RandevuHandler struct {
	service service.RandevuService
}

// NewRandevuHandler creates a new RandevuHandler instance
func NewRandevuHandler(service service.RandevuService) *RandevuHandler {
	return &RandevuHandler{service: service}
}

// GetByKodu handles GET /api/v1/randevu/:kodu
func (h *RandevuHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_RANDEVU_KODU, "Appointment code is required", nil)
		return
	}

	randevu, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_RANDEVU_NOT_FOUND, "Appointment not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_RANDEVU_RETRIEVED, "Appointment retrieved successfully", randevu)
}

// GetByHasta handles GET /api/v1/randevu/hasta/:hasta_kodu
func (h *RandevuHandler) GetByHasta(c *gin.Context) {
	hastaKodu := c.Param("hasta_kodu")
	if hastaKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_HASTA_KODU, "Patient code is required", nil)
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

	randevular, total, err := h.service.GetByHastaKodu(hastaKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RANDEVULAR_RETRIEVED, "Appointments retrieved successfully", randevular, meta)
}

// GetByBasvuru handles GET /api/v1/randevu/basvuru/:basvuru_kodu
func (h *RandevuHandler) GetByBasvuru(c *gin.Context) {
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

	randevular, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RANDEVULAR_RETRIEVED, "Appointments retrieved successfully", randevular, meta)
}

// GetByHekim handles GET /api/v1/randevu/hekim/:hekim_kodu
func (h *RandevuHandler) GetByHekim(c *gin.Context) {
	hekimKodu := c.Param("hekim_kodu")
	if hekimKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PERSONEL_KODU, "Physician code is required", nil)
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

	randevular, total, err := h.service.GetByHekimKodu(hekimKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RANDEVULAR_RETRIEVED, "Appointments retrieved successfully", randevular, meta)
}

// GetByTuru handles GET /api/v1/randevu/turu/:randevu_turu
func (h *RandevuHandler) GetByTuru(c *gin.Context) {
	randevuTuru := c.Param("randevu_turu")
	if randevuTuru == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Appointment type is required", nil)
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

	randevular, total, err := h.service.GetByTuru(randevuTuru, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RANDEVULAR_RETRIEVED, "Appointments retrieved successfully", randevular, meta)
}

// GetByDateRange handles GET /api/v1/randevu/date-range
func (h *RandevuHandler) GetByDateRange(c *gin.Context) {
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

	randevular, total, err := h.service.GetByDateRange(startDate, endDate, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve appointments", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RANDEVULAR_RETRIEVED, "Appointments retrieved successfully", randevular, meta)
}
