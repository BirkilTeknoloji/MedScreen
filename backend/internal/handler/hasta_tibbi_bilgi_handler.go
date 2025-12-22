package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HastaTibbiBilgiHandler handles HTTP requests for patient medical information operations (read-only)
type HastaTibbiBilgiHandler struct {
	service service.HastaTibbiBilgiService
}

// NewHastaTibbiBilgiHandler creates a new HastaTibbiBilgiHandler instance
func NewHastaTibbiBilgiHandler(service service.HastaTibbiBilgiService) *HastaTibbiBilgiHandler {
	return &HastaTibbiBilgiHandler{service: service}
}

// GetByKodu handles GET /api/v1/hasta-tibbi-bilgi/:kodu
func (h *HastaTibbiBilgiHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_HASTA_TIBBI_BILGI_KODU, "Medical information code is required", nil)
		return
	}

	bilgi, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_HASTA_TIBBI_BILGI_NOT_FOUND, "Medical information not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_HASTA_TIBBI_BILGI_RETRIEVED, "Medical information retrieved successfully", bilgi)
}

// GetByHasta handles GET /api/v1/hasta-tibbi-bilgi/hasta/:hasta_kodu
func (h *HastaTibbiBilgiHandler) GetByHasta(c *gin.Context) {
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

	bilgiler, total, err := h.service.GetByHastaKodu(hastaKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve medical information", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_TIBBI_BILGILER_RETRIEVED, "Medical information retrieved successfully", bilgiler, meta)
}

// GetByTuru handles GET /api/v1/hasta-tibbi-bilgi/turu/:turu_kodu
func (h *HastaTibbiBilgiHandler) GetByTuru(c *gin.Context) {
	turuKodu := c.Param("turu_kodu")
	if turuKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Type code is required", nil)
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

	bilgiler, total, err := h.service.GetByTuru(turuKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve medical information", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_TIBBI_BILGILER_RETRIEVED, "Medical information retrieved successfully", bilgiler, meta)
}
