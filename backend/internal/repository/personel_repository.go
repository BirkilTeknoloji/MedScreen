package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// personelRepository implements PersonelRepository interface
type personelRepository struct {
	db *gorm.DB
}

// NewPersonelRepository creates a new PersonelRepository instance
func NewPersonelRepository(db *gorm.DB) PersonelRepository {
	return &personelRepository{db: db}
}

// FindByKodu retrieves a personnel by their code
func (r *personelRepository) FindByKodu(kodu string) (*models.Personel, error) {
	var personel models.Personel
	if err := r.db.Where("personel_kodu = ?", kodu).First(&personel).Error; err != nil {
		return nil, err
	}
	return &personel, nil
}

// FindAll retrieves all personnel with pagination
func (r *personelRepository) FindAll(page, limit int) ([]models.Personel, int64, error) {
	var personeller []models.Personel
	var total int64

	// Count total records
	if err := r.db.Model(&models.Personel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Offset(offset).Limit(limit).Find(&personeller).Error; err != nil {
		return nil, 0, err
	}

	return personeller, total, nil
}

// FindByGorevKodu retrieves personnel by their role code with pagination
func (r *personelRepository) FindByGorevKodu(gorevKodu string, page, limit int) ([]models.Personel, int64, error) {
	var personeller []models.Personel
	var total int64

	// Count total records matching the role
	if err := r.db.Model(&models.Personel{}).Where("personel_gorev_kodu = ?", gorevKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	if err := r.db.Where("personel_gorev_kodu = ?", gorevKodu).Offset(offset).Limit(limit).Find(&personeller).Error; err != nil {
		return nil, 0, err
	}

	return personeller, total, nil
}
