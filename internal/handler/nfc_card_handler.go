package handler

import (
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NFCCardHandler handles HTTP requests for NFC card operations
type NFCCardHandler struct {
	service     service.NFCCardService
	userService service.UserService
}

// NewNFCCardHandler creates a new NFCCardHandler instance
func NewNFCCardHandler(service service.NFCCardService, userService service.UserService) *NFCCardHandler {
	return &NFCCardHandler{
		service:     service,
		userService: userService,
	}
}

// CreateCard handles POST /api/v1/nfc-cards
func (h *NFCCardHandler) CreateCard(c *gin.Context) {
	var card models.NFCCard
	if err := c.ShouldBindJSON(&card); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.CreateCard(&card); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create NFC card", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, "NFC card created successfully", card)
}

// GetCard handles GET /api/v1/nfc-cards/:id
func (h *NFCCardHandler) GetCard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid card ID", err)
		return
	}

	card, err := h.service.GetCard(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "NFC card not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "NFC card retrieved successfully", card)
}

// GetCards handles GET /api/v1/nfc-cards
func (h *NFCCardHandler) GetCards(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	cards, total, err := h.service.GetCards(page, limit)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve NFC cards", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "NFC cards retrieved successfully", cards, meta)
}

// UpdateCard handles PUT /api/v1/nfc-cards/:id
func (h *NFCCardHandler) UpdateCard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid card ID", err)
		return
	}

	var card models.NFCCard
	if err := c.ShouldBindJSON(&card); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateCard(uint(id), &card); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update NFC card", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "NFC card updated successfully", card)
}

// DeleteCard handles DELETE /api/v1/nfc-cards/:id
func (h *NFCCardHandler) DeleteCard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid card ID", err)
		return
	}

	if err := h.service.DeleteCard(uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete NFC card", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "NFC card deleted successfully", nil)
}

// AssignCard handles POST /api/v1/nfc-cards/:id/assign
func (h *NFCCardHandler) AssignCard(c *gin.Context) {
	cardID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid card ID", err)
		return
	}

	var request struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.service.AssignCardToUser(uint(cardID), request.UserID); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to assign card to user", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Card assigned to user successfully", nil)
}

// DeactivateCard handles POST /api/v1/nfc-cards/:id/deactivate
func (h *NFCCardHandler) DeactivateCard(c *gin.Context) {
	cardID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid card ID", err)
		return
	}

	if err := h.service.DeactivateCard(uint(cardID)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to deactivate card", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Card deactivated successfully", nil)
}

// AuthenticateByNFC handles POST /api/v1/nfc-cards/authenticate
func (h *NFCCardHandler) AuthenticateByNFC(c *gin.Context) {
	var request struct {
		CardUID string `json:"card_uid" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := h.userService.AuthenticateByNFC(request.CardUID)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Authentication successful", user)
}
