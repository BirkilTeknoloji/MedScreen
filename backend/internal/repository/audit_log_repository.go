package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// AuditLogRepository defines the interface for audit log operations
type AuditLogRepository interface {
	Create(log *models.AuditLog) error
	FindAll(page, limit int) ([]models.AuditLog, int64, error)
	FindByEntity(entityName string, entityID uint, page, limit int) ([]models.AuditLog, int64, error)
	FindByUser(userID uint, page, limit int) ([]models.AuditLog, int64, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository creates a new instance of AuditLogRepository
func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

// Create creates a new audit log entry
func (r *auditLogRepository) Create(log *models.AuditLog) error {
	return r.db.Create(log).Error
}

// FindAll retrieves all audit logs with pagination
func (r *auditLogRepository) FindAll(page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	if err := r.db.Model(&models.AuditLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := r.db.Preload("User").Order("created_at desc").Offset(offset).Limit(limit).Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// FindByEntity retrieves audit logs for a specific entity
func (r *auditLogRepository) FindByEntity(entityName string, entityID uint, page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	query := r.db.Model(&models.AuditLog{}).Where("entity_name = ? AND entity_id = ?", entityName, entityID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Preload("User").Order("created_at desc").Offset(offset).Limit(limit).Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// FindByUser retrieves audit logs for a specific user
func (r *auditLogRepository) FindByUser(userID uint, page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	query := r.db.Model(&models.AuditLog{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Preload("User").Order("created_at desc").Offset(offset).Limit(limit).Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
