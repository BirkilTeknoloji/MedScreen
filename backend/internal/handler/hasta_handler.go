package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HastaHandler handles HTTP requests for patient operations (read-only)
type HastaHandler struct {
	service service.HastaService
}

// NewHastaHandler creates a new HastaHandler instance
func NewHastaHandler(service service.HastaService) *HastaHandler {
	return &HastaHandler{service: service}
}

// GetByKodu handles GET /api/v1/hasta/:kodu
func (h *HastaHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_HASTA_KODU, "Patient code is required", nil)
		return
	}

	hasta, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_HASTA_NOT_FOUND, "Patient not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_HASTA_RETRIEVED, "Patient retrieved successfully", hasta)
}

// GetByTCKimlik handles GET /api/v1/hasta/tc/:tc_kimlik
func (h *HastaHandler) GetByTCKimlik(c *gin.Context) {
	tcKimlik := c.Param("tc_kimlik")
	if tcKimlik == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "TC Kimlik number is required", nil)
		return
	}

	hasta, err := h.service.GetByTCKimlik(tcKimlik)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_HASTA_NOT_FOUND, "Patient not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_HASTA_RETRIEVED, "Patient retrieved successfully", hasta)
}

// GetAll handles GET /api/v1/hasta
func (h *HastaHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	hastalar, total, err := h.service.GetAll(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve patients", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTALAR_RETRIEVED, "Patients retrieved successfully", hastalar, meta)
}

// Search handles GET /api/v1/hasta/search
func (h *HastaHandler) Search(c *gin.Context) {
	ad := c.Query("ad")
	soyadi := c.Query("soyadi")
	if ad == "" && soyadi == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Search ad or soyadi is required", nil)
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

	hastalar, total, err := h.service.SearchByAdSoyadi(ad, soyadi, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to search patients", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTALAR_RETRIEVED, "Patients search results retrieved successfully", hastalar, meta)
}
