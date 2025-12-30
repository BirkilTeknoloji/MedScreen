package repository

import (
	"medscreen/internal/models"
	"medscreen/internal/utils"

	"gorm.io/gorm"
)

// hastaUyariRepository implements HastaUyariRepository interface
type hastaUyariRepository struct {
	db *gorm.DB
}

// NewHastaUyariRepository creates a new HastaUyariRepository instance
func NewHastaUyariRepository(db *gorm.DB) HastaUyariRepository {
	return &hastaUyariRepository{db: db}
}

// FindByKodu retrieves a patient warning by its code
func (r *hastaUyariRepository) FindByKodu(kodu string) (*models.HastaUyari, error) {
	var uyari models.HastaUyari
	if err := r.db.Preload("HastaBasvuru").
		Where("hasta_uyari_kodu = ?", kodu).First(&uyari).Error; err != nil {
		return nil, err
	}
	sanitizeHastaUyari(&uyari)
	return &uyari, nil
}

// FindByBasvuruKodu retrieves warnings by visit code with pagination
func (r *hastaUyariRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaUyari, int64, error) {
	var uyarilar []models.HastaUyari
	var total int64

	if err := r.db.Model(&models.HastaUyari{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&uyarilar).Error; err != nil {
		return nil, 0, err
	}

	for i := range uyarilar {
		sanitizeHastaUyari(&uyarilar[i])
	}
	return uyarilar, total, nil
}

// FindByTuru retrieves warnings by type with pagination
func (r *hastaUyariRepository) FindByTuru(uyariTuru string, page, limit int) ([]models.HastaUyari, int64, error) {
	var uyarilar []models.HastaUyari
	var total int64

	if err := r.db.Model(&models.HastaUyari{}).Where("uyari_turu = ?", uyariTuru).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("uyari_turu = ?", uyariTuru).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&uyarilar).Error; err != nil {
		return nil, 0, err
	}

	for i := range uyarilar {
		sanitizeHastaUyari(&uyarilar[i])
	}
	return uyarilar, total, nil
}

// FindByAktiflik retrieves warnings by active status with pagination
func (r *hastaUyariRepository) FindByAktiflik(aktiflik int, page, limit int) ([]models.HastaUyari, int64, error) {
	var uyarilar []models.HastaUyari
	var total int64

	if err := r.db.Model(&models.HastaUyari{}).Where("aktiflik_bilgisi = ?", aktiflik).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").
		Where("aktiflik_bilgisi = ?", aktiflik).
		Order("kayit_zamani DESC").
		Offset(offset).Limit(limit).Find(&uyarilar).Error; err != nil {
		return nil, 0, err
	}

	for i := range uyarilar {
		sanitizeHastaUyari(&uyarilar[i])
	}
	return uyarilar, total, nil
}

func sanitizeHastaUyari(uyari *models.HastaUyari) {
	uyari.UyariAciklama = utils.NormalizeUTF8Ptr(uyari.UyariAciklama)
}
