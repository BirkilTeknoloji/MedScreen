package handler

import (
	"medscreen/internal/constants"
	"medscreen/internal/models"
	"medscreen/internal/service"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser handles POST /api/v1/users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.CreateUser(c.Request.Context(), &user); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_USER_CREATE_FAILED, "Failed to create user", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, constants.SUCCESS_USER_CREATED, "User created successfully", user)
}

// GetUser handles GET /api/v1/users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_USER_ID, "Invalid user ID", err)
		return
	}

	user, err := h.service.GetUser(uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, constants.ERROR_USER_NOT_FOUND, "User not found", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_USER_RETRIEVED, "User retrieved successfully", user)
}

// GetUsers handles GET /api/v1/users
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	var role *models.UserRole
	if roleStr := c.Query("role"); roleStr != "" {
		r := models.UserRole(roleStr)
		role = &r
	}

	users, total, err := h.service.GetUsers(page, limit, role)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_INTERNAL_SERVER, "Failed to retrieve users", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, constants.SUCCESS_USERS_RETRIEVED, "Users retrieved successfully", users, meta)
}

// UpdateUser handles PUT /api/v1/users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_USER_ID, "Invalid user ID", err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_REQUEST, "Invalid request body", err)
		return
	}

	if err := h.service.UpdateUser(c.Request.Context(), uint(id), &user); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_USER_UPDATE_FAILED, "Failed to update user", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_USER_UPDATED, "User updated successfully", user)
}

// DeleteUser handles DELETE /api/v1/users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, constants.ERROR_INVALID_USER_ID, "Invalid user ID", err)
		return
	}

	if err := h.service.DeleteUser(c.Request.Context(), uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, constants.ERROR_USER_DELETE_FAILED, "Failed to delete user", err)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, constants.SUCCESS_USER_DELETED, "User deleted successfully", nil)
}
