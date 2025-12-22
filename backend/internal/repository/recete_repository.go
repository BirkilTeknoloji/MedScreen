package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// receteRepository implements ReceteRepository interface
type receteRepository struct {
	db *gorm.DB
}

// NewReceteRepository creates a new ReceteRepository instance
func NewReceteRepository(db *gorm.DB) ReceteRepository {
	return &receteRepository{db: db}
}

// FindByKodu retrieves a prescription by its code
func (r *receteRepository) FindByKodu(kodu string) (*models.Recete, error) {
	var recete models.Recete
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").Preload("Ilaclar").
		Where("recete_kodu = ?", kodu).First(&recete).Error; err != nil {
		return nil, err
	}
	return &recete, nil
}

// FindByBasvuruKodu retrieves prescriptions by visit code with pagination
func (r *receteRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Recete, int64, error) {
	var receteler []models.Recete
	var total int64

	if err := r.db.Model(&models.Recete{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").Preload("Ilaclar").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("recete_zamani DESC").
		Offset(offset).Limit(limit).Find(&receteler).Error; err != nil {
		return nil, 0, err
	}

	return receteler, total, nil
}

// FindByHekimKodu retrieves prescriptions by physician code with pagination
func (r *receteRepository) FindByHekimKodu(hekimKodu string, page, limit int) ([]models.Recete, int64, error) {
	var receteler []models.Recete
	var total int64

	if err := r.db.Model(&models.Recete{}).Where("hekim_kodu = ?", hekimKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").Preload("Ilaclar").
		Where("hekim_kodu = ?", hekimKodu).
		Order("recete_zamani DESC").
		Offset(offset).Limit(limit).Find(&receteler).Error; err != nil {
		return nil, 0, err
	}

	return receteler, total, nil
}

// FindIlacByReceteKodu retrieves prescription medications by prescription code with pagination
func (r *receteRepository) FindIlacByReceteKodu(receteKodu string, page, limit int) ([]models.ReceteIlac, int64, error) {
	var ilaclar []models.ReceteIlac
	var total int64

	if err := r.db.Model(&models.ReceteIlac{}).Where("recete_kodu = ?", receteKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Where("recete_kodu = ?", receteKodu).
		Offset(offset).Limit(limit).Find(&ilaclar).Error; err != nil {
		return nil, 0, err
	}

	return ilaclar, total, nil
}
