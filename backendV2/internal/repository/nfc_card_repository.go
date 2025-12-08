package repository

import (
	"context"
	"medscreen/internal/models"

	"gorm.io/gorm"
)

type nfcCardRepository struct {
	db *gorm.DB
}

// NewNFCCardRepository creates a new instance of NFCCardRepository
func NewNFCCardRepository(db *gorm.DB) NFCCardRepository {
	return &nfcCardRepository{db: db}
}

// Create creates a new NFC card in the database
func (r *nfcCardRepository) Create(ctx context.Context, card *models.NFCCard) error {
	return r.db.WithContext(ctx).Create(card).Error
}

// FindByID retrieves an NFC card by ID with preloaded relationships
func (r *nfcCardRepository) FindByID(id uint) (*models.NFCCard, error) {
	var card models.NFCCard
	err := r.db.Preload("AssignedUser").Preload("CreatedByUser").First(&card, id).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

// FindAll retrieves all NFC cards with pagination and preloaded relationships
func (r *nfcCardRepository) FindAll(page, limit int) ([]models.NFCCard, int64, error) {
	var cards []models.NFCCard
	var total int64

	// Count total records
	if err := r.db.Model(&models.NFCCard{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("AssignedUser").Preload("CreatedByUser").
		Offset(offset).Limit(limit).Find(&cards).Error
	if err != nil {
		return nil, 0, err
	}

	return cards, total, nil
}

// Update updates an existing NFC card
func (r *nfcCardRepository) Update(ctx context.Context, card *models.NFCCard) error {
	return r.db.WithContext(ctx).Save(card).Error
}

// Delete deletes an NFC card by ID (hard delete since NFCCard doesn't use gorm.Model)
func (r *nfcCardRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.NFCCard{}, id).Error
}

// FindByCardUID retrieves an NFC card by card UID with preloaded relationships
func (r *nfcCardRepository) FindByCardUID(cardUID string) (*models.NFCCard, error) {
	var card models.NFCCard
	err := r.db.Preload("AssignedUser").Preload("CreatedByUser").
		Where("card_uid = ?", cardUID).First(&card).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

// FindByAssignedUserID retrieves NFC cards by assigned user ID with pagination
func (r *nfcCardRepository) FindByAssignedUserID(userID uint, page, limit int) ([]models.NFCCard, int64, error) {
	var cards []models.NFCCard
	var total int64

	// Count total records with user filter
	if err := r.db.Model(&models.NFCCard{}).Where("assigned_user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with user filter and preloaded relationships
	err := r.db.Preload("AssignedUser").Preload("CreatedByUser").
		Where("assigned_user_id = ?", userID).
		Offset(offset).Limit(limit).Find(&cards).Error
	if err != nil {
		return nil, 0, err
	}

	return cards, total, nil
}

// FindByActiveStatus retrieves NFC cards by active status with pagination
func (r *nfcCardRepository) FindByActiveStatus(isActive bool, page, limit int) ([]models.NFCCard, int64, error) {
	var cards []models.NFCCard
	var total int64

	// Count total records with active status filter
	if err := r.db.Model(&models.NFCCard{}).Where("is_active = ?", isActive).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with active status filter and preloaded relationships
	err := r.db.Preload("AssignedUser").Preload("CreatedByUser").
		Where("is_active = ?", isActive).
		Offset(offset).Limit(limit).Find(&cards).Error
	if err != nil {
		return nil, 0, err
	}

	return cards, total, nil
}
