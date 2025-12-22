package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// yatakRepository implements YatakRepository interface
type yatakRepository struct {
	db *gorm.DB
}

// NewYatakRepository creates a new YatakRepository instance
func NewYatakRepository(db *gorm.DB) YatakRepository {
	return &yatakRepository{db: db}
}

// FindByKodu retrieves a bed by its code
func (r *yatakRepository) FindByKodu(kodu string) (*models.Yatak, error) {
	var yatak models.Yatak
	if err := r.db.Where("yatak_kodu = ?", kodu).First(&yatak).Error; err != nil {
		return nil, err
	}
	return &yatak, nil
}

// FindByBirimAndOda retrieves beds by unit and room codes with pagination
func (r *yatakRepository) FindByBirimAndOda(birimKodu, odaKodu string, page, limit int) ([]models.Yatak, int64, error) {
	var yataklar []models.Yatak
	var total int64

	// Count total records
	if err := r.db.Model(&models.Yatak{}).
		Where("birim_kodu = ? AND oda_kodu = ?", birimKodu, odaKodu).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Where("birim_kodu = ? AND oda_kodu = ?", birimKodu, odaKodu).
		Offset(offset).Limit(limit).Find(&yataklar).Error; err != nil {
		return nil, 0, err
	}

	return yataklar, total, nil
}

// FindAll retrieves all beds with pagination
func (r *yatakRepository) FindAll(page, limit int) ([]models.Yatak, int64, error) {
	var yataklar []models.Yatak
	var total int64

	// Count total records
	if err := r.db.Model(&models.Yatak{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Offset(offset).Limit(limit).Find(&yataklar).Error; err != nil {
		return nil, 0, err
	}

	return yataklar, total, nil
}
