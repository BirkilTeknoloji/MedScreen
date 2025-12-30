package repository

import (
	"medscreen/internal/models"
	"medscreen/internal/utils"
	"time"

	"gorm.io/gorm"
)

// klinikSeyirRepository implements KlinikSeyirRepository interface
type klinikSeyirRepository struct {
	db *gorm.DB
}

// NewKlinikSeyirRepository creates a new KlinikSeyirRepository instance
func NewKlinikSeyirRepository(db *gorm.DB) KlinikSeyirRepository {
	return &klinikSeyirRepository{db: db}
}

// FindByKodu retrieves clinical progress notes by their code
func (r *klinikSeyirRepository) FindByKodu(kodu string) (*models.KlinikSeyir, error) {
	var seyir models.KlinikSeyir
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").
		Where("klinik_seyir_kodu = ?", kodu).First(&seyir).Error; err != nil {
		return nil, err
	}
	sanitizeKlinikSeyir(&seyir)
	return &seyir, nil
}

// FindByBasvuruKodu retrieves clinical notes by visit code with pagination
func (r *klinikSeyirRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.KlinikSeyir, int64, error) {
	var seyirler []models.KlinikSeyir
	var total int64

	if err := r.db.Model(&models.KlinikSeyir{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("seyir_zamani DESC").
		Offset(offset).Limit(limit).Find(&seyirler).Error; err != nil {
		return nil, 0, err
	}

	for i := range seyirler {
		sanitizeKlinikSeyir(&seyirler[i])
	}
	return seyirler, total, nil
}

// FindBySeyirTipi retrieves clinical notes by type with pagination
func (r *klinikSeyirRepository) FindBySeyirTipi(seyirTipi string, page, limit int) ([]models.KlinikSeyir, int64, error) {
	var seyirler []models.KlinikSeyir
	var total int64

	if err := r.db.Model(&models.KlinikSeyir{}).Where("seyir_tipi = ?", seyirTipi).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").
		Where("seyir_tipi = ?", seyirTipi).
		Order("seyir_zamani DESC").
		Offset(offset).Limit(limit).Find(&seyirler).Error; err != nil {
		return nil, 0, err
	}

	for i := range seyirler {
		sanitizeKlinikSeyir(&seyirler[i])
	}
	return seyirler, total, nil
}

// FindBySepsisDurumu retrieves clinical notes by sepsis status with pagination
func (r *klinikSeyirRepository) FindBySepsisDurumu(sepsisDurumu int, page, limit int) ([]models.KlinikSeyir, int64, error) {
	var seyirler []models.KlinikSeyir
	var total int64

	if err := r.db.Model(&models.KlinikSeyir{}).Where("sepsis_durumu = ?", sepsisDurumu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").
		Where("sepsis_durumu = ?", sepsisDurumu).
		Order("seyir_zamani DESC").
		Offset(offset).Limit(limit).Find(&seyirler).Error; err != nil {
		return nil, 0, err
	}

	for i := range seyirler {
		sanitizeKlinikSeyir(&seyirler[i])
	}
	return seyirler, total, nil
}

// FindBySeyirTipiAndSepsisDurumu retrieves clinical notes by type and sepsis status with pagination
func (r *klinikSeyirRepository) FindBySeyirTipiAndSepsisDurumu(seyirTipi string, sepsisDurumu int, page, limit int) ([]models.KlinikSeyir, int64, error) {
	var seyirler []models.KlinikSeyir
	var total int64

	if err := r.db.Model(&models.KlinikSeyir{}).
		Where("seyir_tipi = ? AND sepsis_durumu = ?", seyirTipi, sepsisDurumu).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").
		Where("seyir_tipi = ? AND sepsis_durumu = ?", seyirTipi, sepsisDurumu).
		Order("seyir_zamani DESC").
		Offset(offset).Limit(limit).Find(&seyirler).Error; err != nil {
		return nil, 0, err
	}

	for i := range seyirler {
		sanitizeKlinikSeyir(&seyirler[i])
	}
	return seyirler, total, nil
}

// FindByDateRange retrieves clinical notes within a date range with pagination
func (r *klinikSeyirRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.KlinikSeyir, int64, error) {
	var seyirler []models.KlinikSeyir
	var total int64

	if err := r.db.Model(&models.KlinikSeyir{}).
		Where("seyir_zamani >= ? AND seyir_zamani <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").
		Where("seyir_zamani >= ? AND seyir_zamani <= ?", startDate, endDate).
		Order("seyir_zamani DESC").
		Offset(offset).Limit(limit).Find(&seyirler).Error; err != nil {
		return nil, 0, err
	}

	for i := range seyirler {
		sanitizeKlinikSeyir(&seyirler[i])
	}
	return seyirler, total, nil
}

func sanitizeKlinikSeyir(seyir *models.KlinikSeyir) {
	seyir.SeyirBilgisi = utils.NormalizeUTF8(seyir.SeyirBilgisi)
}
