package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// riskSkorlamaRepository implements RiskSkorlamaRepository interface
type riskSkorlamaRepository struct {
	db *gorm.DB
}

// NewRiskSkorlamaRepository creates a new RiskSkorlamaRepository instance
func NewRiskSkorlamaRepository(db *gorm.DB) RiskSkorlamaRepository {
	return &riskSkorlamaRepository{db: db}
}

// FindByKodu retrieves a risk score by its code
func (r *riskSkorlamaRepository) FindByKodu(kodu string) (*models.RiskSkorlama, error) {
	var skor models.RiskSkorlama
	if err := r.db.Preload("HastaBasvuru").
		Where("risk_skorlama_kodu = ?", kodu).First(&skor).Error; err != nil {
		return nil, err
	}
	return &skor, nil
}

// FindByBasvuruKodu retrieves risk scores by visit code with pagination
func (r *riskSkorlamaRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.RiskSkorlama, int64, error) {
	var skorlar []models.RiskSkorlama
	var total int64

	if err := r.db.Model(&models.RiskSkorlama{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("islem_zamani DESC").
		Offset(offset).Limit(limit).Find(&skorlar).Error; err != nil {
		return nil, 0, err
	}

	return skorlar, total, nil
}

// FindByTuru retrieves risk scores by type with pagination
func (r *riskSkorlamaRepository) FindByTuru(turu string, page, limit int) ([]models.RiskSkorlama, int64, error) {
	var skorlar []models.RiskSkorlama
	var total int64

	if err := r.db.Model(&models.RiskSkorlama{}).Where("risk_skorlama_turu = ?", turu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("risk_skorlama_turu = ?", turu).
		Order("islem_zamani DESC").
		Offset(offset).Limit(limit).Find(&skorlar).Error; err != nil {
		return nil, 0, err
	}

	return skorlar, total, nil
}
