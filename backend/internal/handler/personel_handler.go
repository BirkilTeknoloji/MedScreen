package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PersonelHandler handles HTTP requests for personnel operations (read-only)
type PersonelHandler struct {
	service service.PersonelService
}

// NewPersonelHandler creates a new PersonelHandler instance
func NewPersonelHandler(service service.PersonelService) *PersonelHandler {
	return &PersonelHandler{service: service}
}

// GetByKodu handles GET /api/v1/personel/:kodu
func (h *PersonelHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PERSONEL_KODU, "Personnel code is required", nil)
		return
	}

	personel, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_PERSONEL_NOT_FOUND, "Personnel not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_PERSONEL_RETRIEVED, "Personnel retrieved successfully", personel)
}

// GetAll handles GET /api/v1/personel
func (h *PersonelHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	personeller, total, err := h.service.GetAll(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve personnel", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_PERSONELLER_RETRIEVED, "Personnel list retrieved successfully", personeller, meta)
}

// GetByGorev handles GET /api/v1/personel/gorev/:gorev_kodu
func (h *PersonelHandler) GetByGorev(c *gin.Context) {
	gorevKodu := c.Param("gorev_kodu")
	if gorevKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Role code is required", nil)
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

	personeller, total, err := h.service.GetByGorevKodu(gorevKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve personnel by role", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_PERSONELLER_RETRIEVED, "Personnel by role retrieved successfully", personeller, meta)
}

// Authenticate handles GET /api/v1/personel/authenticate/:kart_uid
// This endpoint now returns a JWT token for successful authentication
func (h *PersonelHandler) Authenticate(c *gin.Context) {
	kartUID := c.Param("kart_uid")
	if kartUID == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Card UID is required", nil)
		return
	}

	personel, err := h.service.AuthenticateByNFC(kartUID)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, constants.ERROR_NFC_AUTHENTICATION_FAILED, "NFC authentication failed", err)
		return
	}

	// Generate JWT token for the authenticated user
	// Use a simple user ID (we'll use 1 for now since this is read-only)
	token, err := utils.GenerateJWT(1, personel.PersonelGorevKodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to generate authentication token", err)
		return
	}

	// Return both the personel info and the token
	response := gin.H{
		"personel": personel,
		"token":    token,
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_NFC_AUTHENTICATION, "Authentication successful", response)
}
