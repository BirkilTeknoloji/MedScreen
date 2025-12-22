package repository

import (
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

// hastaBasvuruRepository implements HastaBasvuruRepository interface
type hastaBasvuruRepository struct {
	db *gorm.DB
}

// NewHastaBasvuruRepository creates a new HastaBasvuruRepository instance
func NewHastaBasvuruRepository(db *gorm.DB) HastaBasvuruRepository {
	return &hastaBasvuruRepository{db: db}
}

// FindByKodu retrieves a patient visit by its code
func (r *hastaBasvuruRepository) FindByKodu(kodu string) (*models.HastaBasvuru, error) {
	var basvuru models.HastaBasvuru
	if err := r.db.Preload("Hasta").Preload("Hekim").
		Where("hasta_basvuru_kodu = ?", kodu).First(&basvuru).Error; err != nil {
		return nil, err
	}
	return &basvuru, nil
}

// FindByHastaKodu retrieves patient visits by patient code with pagination
func (r *hastaBasvuruRepository) FindByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaBasvuru, int64, error) {
	var basvurular []models.HastaBasvuru
	var total int64

	// Count total records
	if err := r.db.Model(&models.HastaBasvuru{}).Where("hasta_kodu = ?", hastaKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Preload("Hasta").Preload("Hekim").
		Where("hasta_kodu = ?", hastaKodu).
		Offset(offset).Limit(limit).Find(&basvurular).Error; err != nil {
		return nil, 0, err
	}

	return basvurular, total, nil
}

// FindByHekimKodu retrieves patient visits by physician code with pagination
func (r *hastaBasvuruRepository) FindByHekimKodu(hekimKodu string, page, limit int) ([]models.HastaBasvuru, int64, error) {
	var basvurular []models.HastaBasvuru
	var total int64

	// Count total records
	if err := r.db.Model(&models.HastaBasvuru{}).Where("hekim_kodu = ?", hekimKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Preload("Hasta").Preload("Hekim").
		Where("hekim_kodu = ?", hekimKodu).
		Offset(offset).Limit(limit).Find(&basvurular).Error; err != nil {
		return nil, 0, err
	}

	return basvurular, total, nil
}

// FindByDurum retrieves patient visits by status with pagination
func (r *hastaBasvuruRepository) FindByDurum(durum string, page, limit int) ([]models.HastaBasvuru, int64, error) {
	var basvurular []models.HastaBasvuru
	var total int64

	// Count total records
	if err := r.db.Model(&models.HastaBasvuru{}).Where("basvuru_durumu = ?", durum).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Preload("Hasta").Preload("Hekim").
		Where("basvuru_durumu = ?", durum).
		Offset(offset).Limit(limit).Find(&basvurular).Error; err != nil {
		return nil, 0, err
	}

	return basvurular, total, nil
}

// FindByDateRange retrieves patient visits within a date range with pagination
func (r *hastaBasvuruRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.HastaBasvuru, int64, error) {
	var basvurular []models.HastaBasvuru
	var total int64

	// Count total records within date range
	if err := r.db.Model(&models.HastaBasvuru{}).
		Where("hasta_kabul_zamani >= ? AND hasta_kabul_zamani <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Preload("Hasta").Preload("Hekim").
		Where("hasta_kabul_zamani >= ? AND hasta_kabul_zamani <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&basvurular).Error; err != nil {
		return nil, 0, err
	}

	return basvurular, total, nil
}
