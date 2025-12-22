package repository

import (
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

// hastaVitalFizikiBulguRepository implements HastaVitalFizikiBulguRepository interface
type hastaVitalFizikiBulguRepository struct {
	db *gorm.DB
}

// NewHastaVitalFizikiBulguRepository creates a new HastaVitalFizikiBulguRepository instance
func NewHastaVitalFizikiBulguRepository(db *gorm.DB) HastaVitalFizikiBulguRepository {
	return &hastaVitalFizikiBulguRepository{db: db}
}

// FindByKodu retrieves vital signs by their code
func (r *hastaVitalFizikiBulguRepository) FindByKodu(kodu string) (*models.HastaVitalFizikiBulgu, error) {
	var bulgu models.HastaVitalFizikiBulgu
	if err := r.db.Preload("HastaBasvuru").Preload("Hemsire").
		Where("hasta_vital_fiziki_bulgu_kodu = ?", kodu).First(&bulgu).Error; err != nil {
		return nil, err
	}
	return &bulgu, nil
}

// FindByBasvuruKodu retrieves vital signs by visit code with pagination
func (r *hastaVitalFizikiBulguRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error) {
	var bulgular []models.HastaVitalFizikiBulgu
	var total int64

	// Count total records
	if err := r.db.Model(&models.HastaVitalFizikiBulgu{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Preload("HastaBasvuru").Preload("Hemsire").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("islem_zamani DESC").
		Offset(offset).Limit(limit).Find(&bulgular).Error; err != nil {
		return nil, 0, err
	}

	return bulgular, total, nil
}

// FindByDateRange retrieves vital signs within a date range with pagination
func (r *hastaVitalFizikiBulguRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error) {
	var bulgular []models.HastaVitalFizikiBulgu
	var total int64

	// Count total records within date range
	if err := r.db.Model(&models.HastaVitalFizikiBulgu{}).
		Where("islem_zamani >= ? AND islem_zamani <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Preload("HastaBasvuru").Preload("Hemsire").
		Where("islem_zamani >= ? AND islem_zamani <= ?", startDate, endDate).
		Order("islem_zamani DESC").
		Offset(offset).Limit(limit).Find(&bulgular).Error; err != nil {
		return nil, 0, err
	}

	return bulgular, total, nil
}
