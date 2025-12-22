package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TabletCihazHandler handles HTTP requests for tablet device operations (read-only)
type TabletCihazHandler struct {
	service service.TabletCihazService
}

// NewTabletCihazHandler creates a new TabletCihazHandler instance
func NewTabletCihazHandler(service service.TabletCihazService) *TabletCihazHandler {
	return &TabletCihazHandler{service: service}
}

// GetByKodu handles GET /api/v1/tablet-cihaz/:kodu
func (h *TabletCihazHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_TABLET_CIHAZ_KODU, "Device code is required", nil)
		return
	}

	cihaz, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_TABLET_CIHAZ_NOT_FOUND, "Device not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_TABLET_CIHAZ_RETRIEVED, "Device retrieved successfully", cihaz)
}

// GetByYatak handles GET /api/v1/tablet-cihaz/yatak/:yatak_kodu
func (h *TabletCihazHandler) GetByYatak(c *gin.Context) {
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

	cihazlar, total, err := h.service.GetByYatakKodu(yatakKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve devices", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_TABLET_CIHAZLAR_RETRIEVED, "Devices retrieved successfully", cihazlar, meta)
}

// GetAll handles GET /api/v1/tablet-cihaz
func (h *TabletCihazHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	cihazlar, total, err := h.service.GetAll(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve devices", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_TABLET_CIHAZLAR_RETRIEVED, "Devices retrieved successfully", cihazlar, meta)
}
