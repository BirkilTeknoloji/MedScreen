package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BasvuruTaniHandler handles HTTP requests for diagnosis operations (read-only)
type BasvuruTaniHandler struct {
	service service.BasvuruTaniService
}

// NewBasvuruTaniHandler creates a new BasvuruTaniHandler instance
func NewBasvuruTaniHandler(service service.BasvuruTaniService) *BasvuruTaniHandler {
	return &BasvuruTaniHandler{service: service}
}

// GetByKodu handles GET /api/v1/basvuru-tani/:kodu
func (h *BasvuruTaniHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_BASVURU_TANI_KODU, "Diagnosis code is required", nil)
		return
	}

	tani, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_BASVURU_TANI_NOT_FOUND, "Diagnosis not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_BASVURU_TANI_RETRIEVED, "Diagnosis retrieved successfully", tani)
}

// GetByHasta handles GET /api/v1/basvuru-tani/hasta/:hasta_kodu
func (h *BasvuruTaniHandler) GetByHasta(c *gin.Context) {
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

	tanilar, total, err := h.service.GetByHastaKodu(hastaKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve diagnoses", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_BASVURU_TANILAR_RETRIEVED, "Diagnoses retrieved successfully", tanilar, meta)
}

// GetByBasvuru handles GET /api/v1/basvuru-tani/basvuru/:basvuru_kodu
func (h *BasvuruTaniHandler) GetByBasvuru(c *gin.Context) {
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

	tanilar, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve diagnoses", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_BASVURU_TANILAR_RETRIEVED, "Diagnoses retrieved successfully", tanilar, meta)
}
