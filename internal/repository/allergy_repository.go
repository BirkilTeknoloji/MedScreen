package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

type allergyRepository struct {
	db *gorm.DB
}

// NewAllergyRepository creates a new instance of AllergyRepository
func NewAllergyRepository(db *gorm.DB) AllergyRepository {
	return &allergyRepository{db: db}
}

// Create creates a new allergy entry in the database
func (r *allergyRepository) Create(allergy *models.Allergy) error {
	return r.db.Create(allergy).Error
}

// FindByID retrieves an allergy entry by ID with preloaded relationships
func (r *allergyRepository) FindByID(id uint) (*models.Allergy, error) {
	var allergy models.Allergy
	err := r.db.Preload("Patient").Preload("AddedByDoctor").First(&allergy, id).Error
	if err != nil {
		return nil, err
	}
	return &allergy, nil
}

// FindAll retrieves all allergy entries with pagination and preloaded relationships
func (r *allergyRepository) FindAll(page, limit int) ([]models.Allergy, int64, error) {
	var allergies []models.Allergy
	var total int64

	// Count total records
	if err := r.db.Model(&models.Allergy{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Offset(offset).Limit(limit).Find(&allergies).Error
	if err != nil {
		return nil, 0, err
	}

	return allergies, total, nil
}

// Update updates an existing allergy entry
func (r *allergyRepository) Update(allergy *models.Allergy) error {
	return r.db.Save(allergy).Error
}

// Delete soft deletes an allergy entry by ID
func (r *allergyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Allergy{}, id).Error
}

// FindByPatientID retrieves allergy entries by patient ID with pagination
func (r *allergyRepository) FindByPatientID(patientID uint, page, limit int) ([]models.Allergy, int64, error) {
	var allergies []models.Allergy
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.Allergy{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&allergies).Error
	if err != nil {
		return nil, 0, err
	}

	return allergies, total, nil
}

// FindBySeverity retrieves allergy entries by severity with pagination
func (r *allergyRepository) FindBySeverity(severity models.AllergySeverity, page, limit int) ([]models.Allergy, int64, error) {
	var allergies []models.Allergy
	var total int64

	// Count total records with severity filter
	if err := r.db.Model(&models.Allergy{}).Where("severity = ?", severity).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Where("severity = ?", severity).
		Offset(offset).Limit(limit).Find(&allergies).Error
	if err != nil {
		return nil, 0, err
	}

	return allergies, total, nil
}

// FindByFilters retrieves allergy entries with multiple optional filters
func (r *allergyRepository) FindByFilters(patientID *uint, severity *models.AllergySeverity, page, limit int) ([]models.Allergy, int64, error) {
	var allergies []models.Allergy
	var total int64

	// Build query with filters
	query := r.db.Model(&models.Allergy{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if severity != nil {
		query = query.Where("severity = ?", *severity)
	}

	// Count total records with filters
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Rebuild query for data retrieval with preloaded relationships
	dataQuery := r.db.Preload("Patient").Preload("AddedByDoctor")

	if patientID != nil {
		dataQuery = dataQuery.Where("patient_id = ?", *patientID)
	}
	if severity != nil {
		dataQuery = dataQuery.Where("severity = ?", *severity)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&allergies).Error
	if err != nil {
		return nil, 0, err
	}

	return allergies, total, nil
}
