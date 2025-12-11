package handler

import (
	"medscreen/internal/repository"
	"medscreen/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuditLogHandler struct {
	repo repository.AuditLogRepository
}

func NewAuditLogHandler(repo repository.AuditLogRepository) *AuditLogHandler {
	return &AuditLogHandler{repo: repo}
}

// GetAuditLogs retrieves audit logs with pagination and filtering
func (h *AuditLogHandler) GetAuditLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	entityName := c.Query("entity_name")
	entityIDStr := c.Query("entity_id")
	userIDStr := c.Query("user_id")

	var logs interface{}
	var total int64
	var err error

	if entityName != "" && entityIDStr != "" {
		entityID, _ := strconv.Atoi(entityIDStr)
		logs, total, err = h.repo.FindByEntity(entityName, uint(entityID), page, limit)
	} else if userIDStr != "" {
		userID, _ := strconv.Atoi(userIDStr)
		logs, total, err = h.repo.FindByUser(uint(userID), page, limit)
	} else {
		logs, total, err = h.repo.FindAll(page, limit)
	}

	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to retrieve audit logs", err)
		return
	}

	meta := utils.CalculateMeta(page, limit, total)
	utils.SendSuccessResponseWithMeta(c, http.StatusOK, "AUDIT_LOGS_RETRIEVED", "Audit logs retrieved successfully", logs, meta)
}
