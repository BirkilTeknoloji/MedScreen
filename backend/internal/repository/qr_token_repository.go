package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

// QRTokenRepository defines the interface for QR token data operations
type QRTokenRepository interface {
	Create(ctx context.Context, token *models.QRToken) error
	FindByToken(tokenStr string) (*models.QRToken, error)
	Update(ctx context.Context, token *models.QRToken) error
	Delete(ctx context.Context, id uint) error
	DeleteExpiredTokens(ctx context.Context) error
}

type qrTokenRepository struct {
	db *gorm.DB
}

// NewQRTokenRepository creates a new QR token repository instance
func NewQRTokenRepository(db *gorm.DB) QRTokenRepository {
	return &qrTokenRepository{db: db}
}

// Create creates a new QR token
func (r *qrTokenRepository) Create(ctx context.Context, token *models.QRToken) error {
	return r.db.WithContext(ctx).Create(token).Error
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
func (r *qrTokenRepository) Update(ctx context.Context, token *models.QRToken) error {
	return r.db.WithContext(ctx).Save(token).Error
}

// Delete deletes a QR token by ID
func (r *qrTokenRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.QRToken{}, id).Error
}

// DeleteExpiredTokens deletes all expired tokens
func (r *qrTokenRepository) DeleteExpiredTokens(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&models.QRToken{}).Error
}
