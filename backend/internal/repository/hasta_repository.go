package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// hastaRepository implements HastaRepository interface
type hastaRepository struct {
	db *gorm.DB
}

// NewHastaRepository creates a new HastaRepository instance
func NewHastaRepository(db *gorm.DB) HastaRepository {
	return &hastaRepository{db: db}
}

// FindByKodu retrieves a patient by their code
func (r *hastaRepository) FindByKodu(kodu string) (*models.Hasta, error) {
	var hasta models.Hasta
	if err := r.db.Where("hasta_kodu = ?", kodu).First(&hasta).Error; err != nil {
		return nil, err
	}
	return &hasta, nil
}

// FindByTCKimlik retrieves a patient by their Turkish ID number
func (r *hastaRepository) FindByTCKimlik(tcKimlik string) (*models.Hasta, error) {
	var hasta models.Hasta
	if err := r.db.Where("tc_kimlik_numarasi = ?", tcKimlik).First(&hasta).Error; err != nil {
		return nil, err
	}
	return &hasta, nil
}

// FindAll retrieves all patients with pagination
func (r *hastaRepository) FindAll(page, limit int) ([]models.Hasta, int64, error) {
	var hastalar []models.Hasta
	var total int64

	// Count total records
	if err := r.db.Model(&models.Hasta{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Offset(offset).Limit(limit).Find(&hastalar).Error; err != nil {
		return nil, 0, err
	}

	return hastalar, total, nil
}

// SearchByAdSoyadi searches patients by first name and/or last name using ILIKE for case-insensitive search
func (r *hastaRepository) SearchByAdSoyadi(ad, soyadi string, page, limit int) ([]models.Hasta, int64, error) {
	var hastalar []models.Hasta
	var total int64

	query := r.db.Model(&models.Hasta{})
	if ad != "" {
		query = query.Where("ad ILIKE ?", "%"+ad+"%")
	}
	if soyadi != "" {
		query = query.Where("soyadi ILIKE ?", "%"+soyadi+"%")
	}

	// Count total records matching the search
	if err := query.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := query.Session(&gorm.Session{}).Offset(offset).Limit(limit).Find(&hastalar).Error; err != nil {
		return nil, 0, err
	}

	return hastalar, total, nil
}
