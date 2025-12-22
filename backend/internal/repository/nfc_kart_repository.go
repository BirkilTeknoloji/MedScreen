package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// nfcKartRepository implements NFCKartRepository interface
type nfcKartRepository struct {
	db *gorm.DB
}

// NewNFCKartRepository creates a new NFCKartRepository instance
func NewNFCKartRepository(db *gorm.DB) NFCKartRepository {
	return &nfcKartRepository{db: db}
}

// FindByKodu retrieves an NFC card by its code
func (r *nfcKartRepository) FindByKodu(kodu string) (*models.NFCKart, error) {
	var nfcKart models.NFCKart
	if err := r.db.Preload("Personel").Where("nfc_kart_kodu = ?", kodu).First(&nfcKart).Error; err != nil {
		return nil, err
	}
	return &nfcKart, nil
}

// FindByKartUID retrieves an NFC card by its UID
func (r *nfcKartRepository) FindByKartUID(kartUID string) (*models.NFCKart, error) {
	var nfcKart models.NFCKart
	if err := r.db.Preload("Personel").Where("kart_uid = ?", kartUID).First(&nfcKart).Error; err != nil {
		return nil, err
	}
	return &nfcKart, nil
}

// FindByPersonelKodu retrieves NFC cards by personnel code with pagination
func (r *nfcKartRepository) FindByPersonelKodu(personelKodu string, page, limit int) ([]models.NFCKart, int64, error) {
	var nfcKartlar []models.NFCKart
	var total int64

	// Count total records
	if err := r.db.Model(&models.NFCKart{}).Where("personel_kodu = ?", personelKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with Personel preloading
	if err := r.db.Preload("Personel").Where("personel_kodu = ?", personelKodu).Offset(offset).Limit(limit).Find(&nfcKartlar).Error; err != nil {
		return nil, 0, err
	}

	return nfcKartlar, total, nil
}

// FindAll retrieves all NFC cards with pagination
func (r *nfcKartRepository) FindAll(page, limit int) ([]models.NFCKart, int64, error) {
	var nfcKartlar []models.NFCKart
	var total int64

	// Count total records
	if err := r.db.Model(&models.NFCKart{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with Personel preloading
	if err := r.db.Preload("Personel").Offset(offset).Limit(limit).Find(&nfcKartlar).Error; err != nil {
		return nil, 0, err
	}

	return nfcKartlar, total, nil
}
