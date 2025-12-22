package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RiskSkorlamaHandler handles HTTP requests for risk scoring operations (read-only)
type RiskSkorlamaHandler struct {
	service service.RiskSkorlamaService
}

// NewRiskSkorlamaHandler creates a new RiskSkorlamaHandler instance
func NewRiskSkorlamaHandler(service service.RiskSkorlamaService) *RiskSkorlamaHandler {
	return &RiskSkorlamaHandler{service: service}
}

// GetByKodu handles GET /api/v1/risk-skorlama/:kodu
func (h *RiskSkorlamaHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_RISK_SKORLAMA_KODU, "Risk score code is required", nil)
		return
	}

	skorlama, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_RISK_SKORLAMA_NOT_FOUND, "Risk score not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_RISK_SKORLAMA_RETRIEVED, "Risk score retrieved successfully", skorlama)
}

// GetByBasvuru handles GET /api/v1/risk-skorlama/basvuru/:basvuru_kodu
func (h *RiskSkorlamaHandler) GetByBasvuru(c *gin.Context) {
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

	skorlamalar, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve risk scores", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RISK_SKORLAMALAR_RETRIEVED, "Risk scores retrieved successfully", skorlamalar, meta)
}

// GetByTuru handles GET /api/v1/risk-skorlama/turu/:turu
func (h *RiskSkorlamaHandler) GetByTuru(c *gin.Context) {
	turu := c.Param("turu")
	if turu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Risk score type is required", nil)
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

	skorlamalar, total, err := h.service.GetByTuru(turu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve risk scores", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_RISK_SKORLAMALAR_RETRIEVED, "Risk scores retrieved successfully", skorlamalar, meta)
}
