package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TibbiOrderHandler handles HTTP requests for medical orders operations (read-only)
type TibbiOrderHandler struct {
	service service.TibbiOrderService
}

// NewTibbiOrderHandler creates a new TibbiOrderHandler instance
func NewTibbiOrderHandler(service service.TibbiOrderService) *TibbiOrderHandler {
	return &TibbiOrderHandler{service: service}
}

// GetByKodu handles GET /api/v1/tibbi-order/:kodu
func (h *TibbiOrderHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_TIBBI_ORDER_KODU, "Medical order code is required", nil)
		return
	}

	order, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_TIBBI_ORDER_NOT_FOUND, "Medical order not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_TIBBI_ORDER_RETRIEVED, "Medical order retrieved successfully", order)
}

// GetByBasvuru handles GET /api/v1/tibbi-order/basvuru/:basvuru_kodu
func (h *TibbiOrderHandler) GetByBasvuru(c *gin.Context) {
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

	orders, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve medical orders", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_TIBBI_ORDERLAR_RETRIEVED, "Medical orders retrieved successfully", orders, meta)
}

// GetDetay handles GET /api/v1/tibbi-order/:kodu/detay
func (h *TibbiOrderHandler) GetDetay(c *gin.Context) {
	orderKodu := c.Param("kodu")
	if orderKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_TIBBI_ORDER_KODU, "Medical order code is required", nil)
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

	detaylar, total, err := h.service.GetDetayByOrderKodu(orderKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve order details", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_TIBBI_ORDER_DETAY_RETRIEVED, "Order details retrieved successfully", detaylar, meta)
}
