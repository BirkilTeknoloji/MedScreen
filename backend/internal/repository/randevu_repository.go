package repository

import (
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

// randevuRepository implements RandevuRepository interface
type randevuRepository struct {
	db *gorm.DB
}

// NewRandevuRepository creates a new RandevuRepository instance
func NewRandevuRepository(db *gorm.DB) RandevuRepository {
	return &randevuRepository{db: db}
}

// FindByKodu retrieves an appointment by its code
func (r *randevuRepository) FindByKodu(kodu string) (*models.Randevu, error) {
	var randevu models.Randevu
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("randevu_kodu = ?", kodu).First(&randevu).Error; err != nil {
		return nil, err
	}
	return &randevu, nil
}

// FindByHastaKodu retrieves appointments by patient code with pagination
func (r *randevuRepository) FindByHastaKodu(hastaKodu string, page, limit int) ([]models.Randevu, int64, error) {
	var randevular []models.Randevu
	var total int64

	if err := r.db.Model(&models.Randevu{}).Where("hasta_kodu = ?", hastaKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("hasta_kodu = ?", hastaKodu).
		Order("randevu_zamani DESC").
		Offset(offset).Limit(limit).Find(&randevular).Error; err != nil {
		return nil, 0, err
	}

	return randevular, total, nil
}

// FindByBasvuruKodu retrieves appointments by visit code with pagination
func (r *randevuRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Randevu, int64, error) {
	var randevular []models.Randevu
	var total int64

	if err := r.db.Model(&models.Randevu{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("randevu_zamani DESC").
		Offset(offset).Limit(limit).Find(&randevular).Error; err != nil {
		return nil, 0, err
	}

	return randevular, total, nil
}

// FindByHekimKodu retrieves appointments by physician code with pagination
func (r *randevuRepository) FindByHekimKodu(hekimKodu string, page, limit int) ([]models.Randevu, int64, error) {
	var randevular []models.Randevu
	var total int64

	if err := r.db.Model(&models.Randevu{}).Where("hekim_kodu = ?", hekimKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("hekim_kodu = ?", hekimKodu).
		Order("randevu_zamani DESC").
		Offset(offset).Limit(limit).Find(&randevular).Error; err != nil {
		return nil, 0, err
	}

	return randevular, total, nil
}

// FindByTuru retrieves appointments by type with pagination
func (r *randevuRepository) FindByTuru(randevuTuru string, page, limit int) ([]models.Randevu, int64, error) {
	var randevular []models.Randevu
	var total int64

	if err := r.db.Model(&models.Randevu{}).Where("randevu_turu = ?", randevuTuru).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("randevu_turu = ?", randevuTuru).
		Order("randevu_zamani DESC").
		Offset(offset).Limit(limit).Find(&randevular).Error; err != nil {
		return nil, 0, err
	}

	return randevular, total, nil
}

// FindByDateRange retrieves appointments within a date range with pagination
func (r *randevuRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Randevu, int64, error) {
	var randevular []models.Randevu
	var total int64

	if err := r.db.Model(&models.Randevu{}).
		Where("randevu_zamani >= ? AND randevu_zamani <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Hasta").Preload("HastaBasvuru").Preload("Hekim").
		Where("randevu_zamani >= ? AND randevu_zamani <= ?", startDate, endDate).
		Order("randevu_zamani DESC").
		Offset(offset).Limit(limit).Find(&randevular).Error; err != nil {
		return nil, 0, err
	}

	return randevular, total, nil
}
