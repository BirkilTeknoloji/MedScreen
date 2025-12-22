package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TetkikSonucHandler handles HTTP requests for test results operations (read-only)
type TetkikSonucHandler struct {
	service service.TetkikSonucService
}

// NewTetkikSonucHandler creates a new TetkikSonucHandler instance
func NewTetkikSonucHandler(service service.TetkikSonucService) *TetkikSonucHandler {
	return &TetkikSonucHandler{service: service}
}

// GetByKodu handles GET /api/v1/tetkik-sonuc/:kodu
func (h *TetkikSonucHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_TETKIK_SONUC_KODU, "Test result code is required", nil)
		return
	}

	sonuc, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_TETKIK_SONUC_NOT_FOUND, "Test result not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_TETKIK_SONUC_RETRIEVED, "Test result retrieved successfully", sonuc)
}

// GetByBasvuru handles GET /api/v1/tetkik-sonuc/basvuru/:basvuru_kodu
func (h *TetkikSonucHandler) GetByBasvuru(c *gin.Context) {
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

	sonuclar, total, err := h.service.GetByBasvuruKodu(basvuruKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve test results", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_TETKIK_SONUCLAR_RETRIEVED, "Test results retrieved successfully", sonuclar, meta)
}
