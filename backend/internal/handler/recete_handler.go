package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReceteHandler handles HTTP requests for prescription operations (read-only)
type ReceteHandler struct {
	service service.ReceteService
}

// NewReceteHandler creates a new ReceteHandler instance
func NewReceteHandler(service service.ReceteService) *ReceteHandler {
	return &ReceteHandler{service: service}
}

// GetByKodu handles GET /api/v1/recete/:kodu
func (h *ReceteHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_RECETE_KODU, "Prescription code is required", nil)
		return
	}

	recete, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_RECETE_NOT_FOUND, "Prescription not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_RECETE_RETRIEVED, "Prescription retrieved successfully", recete)
}

// GetByBasvuru handles GET /api/v1/recete/basvuru/:basvuru_kodu
func (h *ReceteHandler) GetByBasvuru(c *gin.Context) {
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

	receteler, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve prescriptions", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RECETELER_RETRIEVED, "Prescriptions retrieved successfully", receteler, meta)
}

// GetByHekim handles GET /api/v1/recete/hekim/:hekim_kodu
func (h *ReceteHandler) GetByHekim(c *gin.Context) {
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

	receteler, total, err := h.service.GetByHekimKodu(hekimKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve prescriptions", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RECETELER_RETRIEVED, "Prescriptions retrieved successfully", receteler, meta)
}

// GetIlaclar handles GET /api/v1/recete/:kodu/ilaclar
func (h *ReceteHandler) GetIlaclar(c *gin.Context) {
	receteKodu := c.Param("kodu")
	if receteKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_RECETE_KODU, "Prescription code is required", nil)
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

	ilaclar, total, err := h.service.GetIlaclar(receteKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve medications", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RECETE_ILACLAR_RETRIEVED, "Medications retrieved successfully", ilaclar, meta)
}
