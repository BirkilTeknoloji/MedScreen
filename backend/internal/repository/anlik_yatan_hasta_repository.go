package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// anlikYatanHastaRepository implements AnlikYatanHastaRepository interface
type anlikYatanHastaRepository struct {
	db *gorm.DB
}

// NewAnlikYatanHastaRepository creates a new AnlikYatanHastaRepository instance
func NewAnlikYatanHastaRepository(db *gorm.DB) AnlikYatanHastaRepository {
	return &anlikYatanHastaRepository{db: db}
}

// FindByKodu retrieves a current inpatient by its code
func (r *anlikYatanHastaRepository) FindByKodu(kodu string) (*models.AnlikYatanHasta, error) {
	var yatanHasta models.AnlikYatanHasta
	if err := r.db.Preload("Hasta").Preload("Yatak").Preload("Hekim").Preload("HastaBasvuru").
		Where("anlik_yatan_hasta_kodu = ?", kodu).First(&yatanHasta).Error; err != nil {
		return nil, err
	}
	return &yatanHasta, nil
}

// FindByYatakKodu retrieves current inpatients by bed code with pagination
func (r *anlikYatanHastaRepository) FindByYatakKodu(yatakKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error) {
	var yatanHastalar []models.AnlikYatanHasta
	var total int64

	// Count total records
	if err := r.db.Model(&models.AnlikYatanHasta{}).Where("yatak_kodu = ?", yatakKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with preloading
	if err := r.db.Preload("Hasta").Preload("Yatak").Preload("Hekim").Preload("HastaBasvuru").
		Where("yatak_kodu = ?", yatakKodu).
		Offset(offset).Limit(limit).Find(&yatanHastalar).Error; err != nil {
		return nil, 0, err
	}

	return yatanHastalar, total, nil
}

// FindByHastaKodu retrieves current inpatients by patient code with pagination
func (r *anlikYatanHastaRepository) FindByHastaKodu(hastaKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error) {
	var yatanHastalar []models.AnlikYatanHasta
	var total int64

	// Count total records
	if err := r.db.Model(&models.AnlikYatanHasta{}).Where("hasta_kodu = ?", hastaKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with preloading
	if err := r.db.Preload("Hasta").Preload("Yatak").Preload("Hekim").Preload("HastaBasvuru").
		Where("hasta_kodu = ?", hastaKodu).
		Offset(offset).Limit(limit).Find(&yatanHastalar).Error; err != nil {
		return nil, 0, err
	}

	return yatanHastalar, total, nil
}

// FindByBirimKodu retrieves current inpatients by unit code with pagination
func (r *anlikYatanHastaRepository) FindByBirimKodu(birimKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error) {
	var yatanHastalar []models.AnlikYatanHasta
	var total int64

	// Count total records
	if err := r.db.Model(&models.AnlikYatanHasta{}).Where("birim_kodu = ?", birimKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with preloading
	if err := r.db.Preload("Hasta").Preload("Yatak").Preload("Hekim").Preload("HastaBasvuru").
		Where("birim_kodu = ?", birimKodu).
		Offset(offset).Limit(limit).Find(&yatanHastalar).Error; err != nil {
		return nil, 0, err
	}

	return yatanHastalar, total, nil
}
