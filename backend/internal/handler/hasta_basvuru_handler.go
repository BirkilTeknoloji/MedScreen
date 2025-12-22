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

// HastaBasvuruHandler handles HTTP requests for patient visit operations (read-only)
type HastaBasvuruHandler struct {
	service service.HastaBasvuruService
}

// NewHastaBasvuruHandler creates a new HastaBasvuruHandler instance
func NewHastaBasvuruHandler(service service.HastaBasvuruService) *HastaBasvuruHandler {
	return &HastaBasvuruHandler{service: service}
}

// GetByKodu handles GET /api/v1/hasta-basvuru/:kodu
func (h *HastaBasvuruHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_HASTA_BASVURU_KODU, "Visit code is required", nil)
		return
	}

	basvuru, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_HASTA_BASVURU_NOT_FOUND, "Visit not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_HASTA_BASVURU_RETRIEVED, "Visit retrieved successfully", basvuru)
}

// GetByHasta handles GET /api/v1/hasta-basvuru/hasta/:hasta_kodu
func (h *HastaBasvuruHandler) GetByHasta(c *gin.Context) {
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

	basvurular, total, err := h.service.GetByHastaKodu(hastaKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve visits", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_BASVURULAR_RETRIEVED, "Visits retrieved successfully", basvurular, meta)
}

// GetByHekim handles GET /api/v1/hasta-basvuru/hekim/:hekim_kodu
func (h *HastaBasvuruHandler) GetByHekim(c *gin.Context) {
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

	basvurular, total, err := h.service.GetByHekimKodu(hekimKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve visits", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_BASVURULAR_RETRIEVED, "Visits retrieved successfully", basvurular, meta)
}

// GetByFilters handles GET /api/v1/hasta-basvuru/filter
func (h *HastaBasvuruHandler) GetByFilters(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	var durum *string
	if d := c.Query("durum"); d != "" {
		durum = &d
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

	basvurular, total, err := h.service.GetByFilters(durum, startDate, endDate, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve visits", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_BASVURULAR_RETRIEVED, "Visits retrieved successfully", basvurular, meta)
}
