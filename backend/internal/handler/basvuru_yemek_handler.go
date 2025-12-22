package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BasvuruYemekHandler handles HTTP requests for diet/meal operations (read-only)
type BasvuruYemekHandler struct {
	service service.BasvuruYemekService
}

// NewBasvuruYemekHandler creates a new BasvuruYemekHandler instance
func NewBasvuruYemekHandler(service service.BasvuruYemekService) *BasvuruYemekHandler {
	return &BasvuruYemekHandler{service: service}
}

// GetByKodu handles GET /api/v1/basvuru-yemek/:kodu
func (h *BasvuruYemekHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_BASVURU_YEMEK_KODU, "Meal code is required", nil)
		return
	}

	yemek, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_BASVURU_YEMEK_NOT_FOUND, "Meal not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_BASVURU_YEMEK_RETRIEVED, "Meal retrieved successfully", yemek)
}

// GetByBasvuru handles GET /api/v1/basvuru-yemek/basvuru/:basvuru_kodu
func (h *BasvuruYemekHandler) GetByBasvuru(c *gin.Context) {
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

	yemekler, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve meals", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_BASVURU_YEMEKLER_RETRIEVED, "Meals retrieved successfully", yemekler, meta)
}

// GetByTuru handles GET /api/v1/basvuru-yemek/turu/:yemek_turu
func (h *BasvuruYemekHandler) GetByTuru(c *gin.Context) {
	yemekTuru := c.Param("yemek_turu")
	if yemekTuru == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Meal type is required", nil)
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

	yemekler, total, err := h.service.GetByTuru(yemekTuru, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve meals", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_BASVURU_YEMEKLER_RETRIEVED, "Meals retrieved successfully", yemekler, meta)
}
