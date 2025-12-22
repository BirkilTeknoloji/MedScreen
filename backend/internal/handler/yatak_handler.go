package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// YatakHandler handles HTTP requests for bed operations (read-only)
type YatakHandler struct {
	service service.YatakService
}

// NewYatakHandler creates a new YatakHandler instance
func NewYatakHandler(service service.YatakService) *YatakHandler {
	return &YatakHandler{service: service}
}

// GetByKodu handles GET /api/v1/yatak/:kodu
func (h *YatakHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_YATAK_KODU, "Bed code is required", nil)
		return
	}

	yatak, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_YATAK_NOT_FOUND, "Bed not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_YATAK_RETRIEVED, "Bed retrieved successfully", yatak)
}

// GetByBirimOda handles GET /api/v1/yatak/birim/:birim_kodu/oda/:oda_kodu
func (h *YatakHandler) GetByBirimOda(c *gin.Context) {
	birimKodu := c.Param("birim_kodu")
	odaKodu := c.Param("oda_kodu")

	if birimKodu == "" || odaKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Unit code and room code are required", nil)
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

	yataklar, total, err := h.service.GetByBirimAndOda(birimKodu, odaKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve beds", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_YATAKLAR_RETRIEVED, "Beds retrieved successfully", yataklar, meta)
}

// GetAll handles GET /api/v1/yatak
func (h *YatakHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	yataklar, total, err := h.service.GetAll(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve beds", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_YATAKLAR_RETRIEVED, "Beds retrieved successfully", yataklar, meta)
}
