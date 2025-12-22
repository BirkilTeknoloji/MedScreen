package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NFCKartHandler handles HTTP requests for NFC card operations (read-only)
type NFCKartHandler struct {
	service service.NFCKartService
}

// NewNFCKartHandler creates a new NFCKartHandler instance
func NewNFCKartHandler(service service.NFCKartService) *NFCKartHandler {
	return &NFCKartHandler{service: service}
}

// GetByKodu handles GET /api/v1/nfc-kart/:kodu
func (h *NFCKartHandler) GetByKodu(c *gin.Context) {
	kodu := c.Param("kodu")
	if kodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_NFC_KART_KODU, "NFC card code is required", nil)
		return
	}

	nfcKart, err := h.service.GetByKodu(kodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_NFC_KART_NOT_FOUND, "NFC card not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_NFC_KART_RETRIEVED, "NFC card retrieved successfully", nfcKart)
}

// GetByKartUID handles GET /api/v1/nfc-kart/authenticate/:kart_uid
// This endpoint now returns a JWT token for successful authentication
func (h *NFCKartHandler) GetByKartUID(c *gin.Context) {
	kartUID := c.Param("kart_uid")
	if kartUID == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Card UID is required", nil)
		return
	}

	nfcKart, err := h.service.GetByKartUID(kartUID)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_NFC_KART_NOT_FOUND, "NFC card not found", err)
		return
	}

	// Check if card is active
	if nfcKart.AktiflikBilgisi != 1 {
		utils.SendErrorResponse(c, http.StatusUnauthorized, constants.ERROR_NFC_AUTHENTICATION_FAILED, "NFC card is inactive", nil)
		return
	}

	// Generate JWT token for the authenticated user
	// Use a simple user ID (we'll use 1 for now since this is read-only)
	token, err := utils.GenerateJWT(1, nfcKart.Personel.PersonelGorevKodu)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to generate authentication token", err)
		return
	}

	// Return both the NFC card info and the token
	response := gin.H{
		"nfc_kart": nfcKart,
		"token":    token,
		"personel": nfcKart.Personel,
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_NFC_AUTHENTICATION, "NFC authentication successful", response)
}

// GetByPersonelKodu handles GET /api/v1/nfc-kart/personel/:personel_kodu
func (h *NFCKartHandler) GetByPersonelKodu(c *gin.Context) {
	personelKodu := c.Param("personel_kodu")
	if personelKodu == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_PERSONEL_KODU, "Personnel code is required", nil)
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

	nfcKartlar, total, err := h.service.GetByPersonelKodu(personelKodu, page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve NFC cards", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_NFC_KARTLAR_RETRIEVED, "NFC cards retrieved successfully", nfcKartlar, meta)
}
