package repository

import (
	"context"
	"medscreen/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user in the database
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByID retrieves a user by ID
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll retrieves all users with pagination
func (r *userRepository) FindAll(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Count total records
	if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update updates an existing user
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete soft deletes a user by ID
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

// FindByRole retrieves users by role with pagination
func (r *userRepository) FindByRole(role models.UserRole, page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Count total records with role filter
	if err := r.db.Model(&models.User{}).Where("role = ?", role).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with role filter
	err := r.db.Where("role = ?", role).Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// FindByNFCCardID retrieves a user by NFC card ID (foreign key)
func (r *userRepository) FindByNFCCardID(nfcCardID uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("nfc_card_id = ?", nfcCardID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByNFCCardUID retrieves a user by NFC card UID from the nfc_cards table
func (r *userRepository) FindByNFCCardUID(cardUID string) (*models.User, error) {
	var user models.User
	err := r.db.Table("users").
		Joins("JOIN nfc_cards ON nfc_cards.assigned_user_id = users.id").
		Where("nfc_cards.card_uid = ? AND nfc_cards.is_active = ?", cardUID, true).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
