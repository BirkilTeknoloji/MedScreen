package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// hastaTibbiBilgiRepository implements HastaTibbiBilgiRepository interface
type hastaTibbiBilgiRepository struct {
	db *gorm.DB
}

// NewHastaTibbiBilgiRepository creates a new HastaTibbiBilgiRepository instance
func NewHastaTibbiBilgiRepository(db *gorm.DB) HastaTibbiBilgiRepository {
	return &hastaTibbiBilgiRepository{db: db}
}

// FindByKodu retrieves medical information by its code
func (r *hastaTibbiBilgiRepository) FindByKodu(kodu string) (*models.HastaTibbiBilgi, error) {
	var bilgi models.HastaTibbiBilgi
	if err := r.db.Preload("Hasta").
		Where("hasta_tibbi_bilgi_kodu = ?", kodu).First(&bilgi).Error; err != nil {
		return nil, err
	}
	return &bilgi, nil
}

// FindByHastaKodu retrieves medical information by patient code with pagination
func (r *hastaTibbiBilgiRepository) FindByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error) {
	var bilgiler []models.HastaTibbiBilgi
	var total int64

	if err := r.db.Model(&models.HastaTibbiBilgi{}).Where("hasta_kodu = ?", hastaKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").
		Where("hasta_kodu = ?", hastaKodu).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&bilgiler).Error; err != nil {
		return nil, 0, err
	}

	return bilgiler, total, nil
}

// FindByTuru retrieves medical information by type with pagination
func (r *hastaTibbiBilgiRepository) FindByTuru(turuKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error) {
	var bilgiler []models.HastaTibbiBilgi
	var total int64

	if err := r.db.Model(&models.HastaTibbiBilgi{}).Where("tibbi_bilgi_turu_kodu = ?", turuKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").
		Where("tibbi_bilgi_turu_kodu = ?", turuKodu).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&bilgiler).Error; err != nil {
		return nil, 0, err
	}

	return bilgiler, total, nil
}
