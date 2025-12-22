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

// KlinikSeyirHandler handles HTTP requests for clinical progress notes operations (read-only)
type KlinikSeyirHandler struct {
	service service.KlinikSeyirService
}

// NewKlinikSeyirHandler creates a new KlinikSeyirHandler instance
func NewKlinikSeyirHandler(service service.KlinikSeyirService) *KlinikSeyirHandler {
	return &KlinikSeyirHandler{service: service}
}

// GetByKodu handles GET /api/v1/klinik-seyir/:kodu
func (h *KlinikSeyirHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_KLINIK_SEYIR_KODU, "Clinical note code is required", nil)
		return
	}

	seyir, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_KLINIK_SEYIR_NOT_FOUND, "Clinical note not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_KLINIK_SEYIR_RETRIEVED, "Clinical note retrieved successfully", seyir)
}

// GetByBasvuru handles GET /api/v1/klinik-seyir/basvuru/:basvuru_kodu
func (h *KlinikSeyirHandler) GetByBasvuru(c *gin.Context) {
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

	seyirler, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve clinical notes", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_KLINIK_SEYIRLER_RETRIEVED, "Clinical notes retrieved successfully", seyirler, meta)
}

// GetByFilters handles GET /api/v1/klinik-seyir/filter
func (h *KlinikSeyirHandler) GetByFilters(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	var seyirTipi *string
	if tipi := c.Query("seyir_tipi"); tipi != "" {
		seyirTipi = &tipi
	}

	var startDate, endDate *time.Time
	if start := c.Query("start_date"); start != "" {
		t, err := time.Parse("2006-01-02", start)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_DATE_RANGE, "Invalid start date format (use YYYY-MM-DD)", err)
			return
		}
		startDate = &t
	}

	if end := c.Query("end_date"); end != "" {
		t, err := time.Parse("2006-01-02", end)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_DATE_RANGE, "Invalid end date format (use YYYY-MM-DD)", err)
			return
		}
		endDate = &t
	}

	seyirler, total, err := h.service.GetByFilters(seyirTipi, startDate, endDate, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve clinical notes", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_KLINIK_SEYIRLER_RETRIEVED, "Clinical notes retrieved successfully", seyirler, meta)
}
