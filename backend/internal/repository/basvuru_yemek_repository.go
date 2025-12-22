package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// basvuruYemekRepository implements BasvuruYemekRepository interface
type basvuruYemekRepository struct {
	db *gorm.DB
}

// NewBasvuruYemekRepository creates a new BasvuruYemekRepository instance
func NewBasvuruYemekRepository(db *gorm.DB) BasvuruYemekRepository {
	return &basvuruYemekRepository{db: db}
}

// FindByKodu retrieves meal information by its code
func (r *basvuruYemekRepository) FindByKodu(kodu string) (*models.BasvuruYemek, error) {
	var yemek models.BasvuruYemek
	if err := r.db.Preload("HastaBasvuru").
		Where("basvuru_yemek_kodu = ?", kodu).First(&yemek).Error; err != nil {
		return nil, err
	}
	return &yemek, nil
}

// FindByBasvuruKodu retrieves meal information by visit code with pagination
func (r *basvuruYemekRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruYemek, int64, error) {
	var yemekler []models.BasvuruYemek
	var total int64

	if err := r.db.Model(&models.BasvuruYemek{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&yemekler).Error; err != nil {
		return nil, 0, err
	}

	return yemekler, total, nil
}

// FindByTuru retrieves meal information by type with pagination
func (r *basvuruYemekRepository) FindByTuru(yemekTuru string, page, limit int) ([]models.BasvuruYemek, int64, error) {
	var yemekler []models.BasvuruYemek
	var total int64

	if err := r.db.Model(&models.BasvuruYemek{}).Where("yemek_turu = ?", yemekTuru).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("yemek_turu = ?", yemekTuru).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&yemekler).Error; err != nil {
		return nil, 0, err
	}

	return yemekler, total, nil
}
