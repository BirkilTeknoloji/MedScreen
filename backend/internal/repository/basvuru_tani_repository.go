package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// basvuruTaniRepository implements BasvuruTaniRepository interface
type basvuruTaniRepository struct {
	db *gorm.DB
}

// NewBasvuruTaniRepository creates a new BasvuruTaniRepository instance
func NewBasvuruTaniRepository(db *gorm.DB) BasvuruTaniRepository {
	return &basvuruTaniRepository{db: db}
}

// FindByKodu retrieves a diagnosis by its code
func (r *basvuruTaniRepository) FindByKodu(kodu string) (*models.BasvuruTani, error) {
	var tani models.BasvuruTani
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("basvuru_tani_kodu = ?", kodu).First(&tani).Error; err != nil {
		return nil, err
	}
	return &tani, nil
}

// FindByHastaKodu retrieves diagnoses by patient code with pagination
func (r *basvuruTaniRepository) FindByHastaKodu(hastaKodu string, page, limit int) ([]models.BasvuruTani, int64, error) {
	var tanilar []models.BasvuruTani
	var total int64

	if err := r.db.Model(&models.BasvuruTani{}).Where("hasta_kodu = ?", hastaKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("hasta_kodu = ?", hastaKodu).
		Order("tani_zamani DESC").
		Offset(offset).Limit(limit).Find(&tanilar).Error; err != nil {
		return nil, 0, err
	}

	return tanilar, total, nil
}

// FindByBasvuruKodu retrieves diagnoses by visit code with pagination
func (r *basvuruTaniRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruTani, int64, error) {
	var tanilar []models.BasvuruTani
	var total int64

	if err := r.db.Model(&models.BasvuruTani{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("birincil_tani DESC, tani_zamani DESC").
		Offset(offset).Limit(limit).Find(&tanilar).Error; err != nil {
		return nil, 0, err
	}

	return tanilar, total, nil
}

// FindByTaniKodu retrieves diagnoses by ICD code with pagination
func (r *basvuruTaniRepository) FindByTaniKodu(taniKodu string, page, limit int) ([]models.BasvuruTani, int64, error) {
	var tanilar []models.BasvuruTani
	var total int64

	if err := r.db.Model(&models.BasvuruTani{}).Where("tani_kodu = ?", taniKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("tani_kodu = ?", taniKodu).
		Order("tani_zamani DESC").
		Offset(offset).Limit(limit).Find(&tanilar).Error; err != nil {
		return nil, 0, err
	}

	return tanilar, total, nil
}
