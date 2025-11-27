package repository

import (
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

// QRTokenRepository defines the interface for QR token data operations
type QRTokenRepository interface {
	Create(token *models.QRToken) error
	FindByToken(tokenStr string) (*models.QRToken, error)
	Update(token *models.QRToken) error
	Delete(id uint) error
	DeleteExpiredTokens() error
}

type qrTokenRepository struct {
	db *gorm.DB
}

// NewQRTokenRepository creates a new QR token repository instance
func NewQRTokenRepository(db *gorm.DB) QRTokenRepository {
	return &qrTokenRepository{db: db}
}

// Create creates a new QR token
func (r *qrTokenRepository) Create(token *models.QRToken) error {
	return r.db.Create(token).Error
}

// FindByToken finds a QR token by token string with patient and device relationships
func (r *qrTokenRepository) FindByToken(tokenStr string) (*models.QRToken, error) {
	var token models.QRToken
	err := r.db.Where("token = ?", tokenStr).
		Preload("Patient").
		Preload("Device").
		First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// Update updates a QR token
func (r *qrTokenRepository) Update(token *models.QRToken) error {
	return r.db.Save(token).Error
}

// Delete deletes a QR token by ID
func (r *qrTokenRepository) Delete(id uint) error {
	return r.db.Delete(&models.QRToken{}, id).Error
}

// DeleteExpiredTokens deletes all expired tokens
func (r *qrTokenRepository) DeleteExpiredTokens() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&models.QRToken{}).Error
}
