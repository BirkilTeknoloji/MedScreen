package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AnlikYatanHastaHandler handles HTTP requests for current inpatient operations (read-only)
type AnlikYatanHastaHandler struct {
	service service.AnlikYatanHastaService
}

// NewAnlikYatanHastaHandler creates a new AnlikYatanHastaHandler instance
func NewAnlikYatanHastaHandler(service service.AnlikYatanHastaService) *AnlikYatanHastaHandler {
	return &AnlikYatanHastaHandler{service: service}
}

// GetByKodu handles GET /api/v1/anlik-yatan-hasta/:kodu
func (h *AnlikYatanHastaHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_ANLIK_YATAN_HASTA_KODU, "Inpatient code is required", nil)
		return
	}

	yatanHasta, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_ANLIK_YATAN_HASTA_NOT_FOUND, "Inpatient not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_ANLIK_YATAN_HASTA_RETRIEVED, "Inpatient retrieved successfully", yatanHasta)
}

// GetByYatak handles GET /api/v1/anlik-yatan-hasta/yatak/:yatak_kodu
func (h *AnlikYatanHastaHandler) GetByYatak(c *gin.Context) {
	yatakKodu := c.Param("yatak_kodu")
	if yatakKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_YATAK_KODU, "Bed code is required", nil)
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

	yatanHastalar, total, err := h.service.GetByYatakKodu(yatakKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve inpatients", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_ANLIK_YATAN_HASTALAR_RETRIEVED, "Inpatients retrieved successfully", yatanHastalar, meta)
}

// GetByHasta handles GET /api/v1/anlik-yatan-hasta/hasta/:hasta_kodu
func (h *AnlikYatanHastaHandler) GetByHasta(c *gin.Context) {
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

	yatanHastalar, total, err := h.service.GetByHastaKodu(hastaKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve inpatients", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_ANLIK_YATAN_HASTALAR_RETRIEVED, "Inpatients retrieved successfully", yatanHastalar, meta)
}

// GetByBirim handles GET /api/v1/anlik-yatan-hasta/birim/:birim_kodu
func (h *AnlikYatanHastaHandler) GetByBirim(c *gin.Context) {
	birimKodu := c.Param("birim_kodu")
	if birimKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Unit code is required", nil)
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

	yatanHastalar, total, err := h.service.GetByBirimKodu(birimKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve inpatients", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_ANLIK_YATAN_HASTALAR_RETRIEVED, "Inpatients retrieved successfully", yatanHastalar, meta)
}
