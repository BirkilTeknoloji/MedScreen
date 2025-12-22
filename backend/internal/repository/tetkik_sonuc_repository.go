package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// tetkikSonucRepository implements TetkikSonucRepository interface
type tetkikSonucRepository struct {
	db *gorm.DB
}

// NewTetkikSonucRepository creates a new TetkikSonucRepository instance
func NewTetkikSonucRepository(db *gorm.DB) TetkikSonucRepository {
	return &tetkikSonucRepository{db: db}
}

// FindByKodu retrieves a test result by its code
func (r *tetkikSonucRepository) FindByKodu(kodu string) (*models.TetkikSonuc, error) {
	var sonuc models.TetkikSonuc
	if err := r.db.Preload("HastaBasvuru").
		Where("tetkik_sonuc_kodu = ?", kodu).First(&sonuc).Error; err != nil {
		return nil, err
	}
	return &sonuc, nil
}

// FindByBasvuruKodu retrieves test results by visit code with pagination
func (r *tetkikSonucRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TetkikSonuc, int64, error) {
	var sonuclar []models.TetkikSonuc
	var total int64

	if err := r.db.Model(&models.TetkikSonuc{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&sonuclar).Error; err != nil {
		return nil, 0, err
	}

	return sonuclar, total, nil
}
