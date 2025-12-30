package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HastaUyariHandler handles HTTP requests for patient warnings operations (read-only)
type HastaUyariHandler struct {
	service service.HastaUyariService
}

// NewHastaUyariHandler creates a new HastaUyariHandler instance
func NewHastaUyariHandler(service service.HastaUyariService) *HastaUyariHandler {
	return &HastaUyariHandler{service: service}
}

// GetByKodu handles GET /api/v1/hasta-uyari/:kodu
func (h *HastaUyariHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_HASTA_UYARI_KODU, "Warning code is required", nil)
		return
	}

	uyari, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_HASTA_UYARI_NOT_FOUND, "Warning not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_HASTA_UYARI_RETRIEVED, "Warning retrieved successfully", uyari)
}

// GetByBasvuru handles GET /api/v1/hasta-uyari/basvuru/:basvuru_kodu
func (h *HastaUyariHandler) GetByBasvuru(c *gin.Context) {
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

	uyarilar, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve warnings", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_UYARILAR_RETRIEVED, "Warnings retrieved successfully", uyarilar, meta)
}

// GetByFilters handles GET /api/v1/hasta-uyari/filter
func (h *HastaUyariHandler) GetByFilters(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	var uyariTuru *string
	if turu := c.Query("uyari_turu"); turu != "" {
		uyariTuru = &turu
	}

	var aktiflik *int
	if aktifStr := c.Query("aktiflik"); aktifStr != "" {
		aktifVal, err := strconv.Atoi(aktifStr)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid aktiflik value", err)
			return
		}
		aktiflik = &aktifVal
	} else if aktifStr := c.Query("aktiflik_bilgisi"); aktifStr != "" {
		aktifVal, err := strconv.Atoi(aktifStr)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid aktiflik_bilgisi value", err)
			return
		}
		aktiflik = &aktifVal
	}

	if uyariTuru == nil && aktiflik == nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "At least one filter (uyari_turu or aktiflik/aktiflik_bilgisi) is required", nil)
		return
	}

	uyarilar, total, err := h.service.GetByFilters(uyariTuru, aktiflik, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve warnings", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_HASTA_UYARILAR_RETRIEVED, "Warnings retrieved successfully", uyarilar, meta)
}
