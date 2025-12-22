package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// tabletCihazRepository implements TabletCihazRepository interface
type tabletCihazRepository struct {
	db *gorm.DB
}

// NewTabletCihazRepository creates a new TabletCihazRepository instance
func NewTabletCihazRepository(db *gorm.DB) TabletCihazRepository {
	return &tabletCihazRepository{db: db}
}

// FindByKodu retrieves a tablet device by its code
func (r *tabletCihazRepository) FindByKodu(kodu string) (*models.TabletCihaz, error) {
	var cihaz models.TabletCihaz
	if err := r.db.Preload("Yatak").Where("tablet_cihaz_kodu = ?", kodu).First(&cihaz).Error; err != nil {
		return nil, err
	}
	return &cihaz, nil
}

// FindByYatakKodu retrieves tablet devices by bed code with pagination
func (r *tabletCihazRepository) FindByYatakKodu(yatakKodu string, page, limit int) ([]models.TabletCihaz, int64, error) {
	var cihazlar []models.TabletCihaz
	var total int64

	// Count total records
	if err := r.db.Model(&models.TabletCihaz{}).Where("yatak_kodu = ?", yatakKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with Yatak preloading
	if err := r.db.Preload("Yatak").Where("yatak_kodu = ?", yatakKodu).
		Offset(offset).Limit(limit).Find(&cihazlar).Error; err != nil {
		return nil, 0, err
	}

	return cihazlar, total, nil
}

// FindAll retrieves all tablet devices with pagination
func (r *tabletCihazRepository) FindAll(page, limit int) ([]models.TabletCihaz, int64, error) {
	var cihazlar []models.TabletCihaz
	var total int64

	// Count total records
	if err := r.db.Model(&models.TabletCihaz{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results with Yatak preloading
	if err := r.db.Preload("Yatak").Offset(offset).Limit(limit).Find(&cihazlar).Error; err != nil {
		return nil, 0, err
	}

	return cihazlar, total, nil
}
