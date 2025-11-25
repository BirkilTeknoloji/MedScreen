package handler

//TODO: prodda sil
import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ResetHandler struct {
	db *gorm.DB
}

func NewResetHandler(db *gorm.DB) *ResetHandler {
	return &ResetHandler{db: db}
}

func (h *ResetHandler) ResetDatabase(c *gin.Context) {
	sqlFile := "database_querytest.sql"

	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to read SQL file",
			"error":   err.Error(),
		})
		return
	}

	sqlContent := string(sqlBytes)

	if err := h.db.Exec(sqlContent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to execute reset SQL",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database reset successfully",
	})
}
