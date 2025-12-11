package repository

import (
	"context"
	"medscreen/internal/models"

	"gorm.io/gorm"
)

type medicalHistoryRepository struct {
	db *gorm.DB
}

// NewMedicalHistoryRepository creates a new instance of MedicalHistoryRepository
func NewMedicalHistoryRepository(db *gorm.DB) MedicalHistoryRepository {
	return &medicalHistoryRepository{db: db}
}

// Create creates a new medical history entry in the database
func (r *medicalHistoryRepository) Create(ctx context.Context, history *models.MedicalHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

// FindByID retrieves a medical history entry by ID with preloaded relationships
func (r *medicalHistoryRepository) FindByID(id uint) (*models.MedicalHistory, error) {
	var history models.MedicalHistory
	err := r.db.Preload("Patient").Preload("AddedByDoctor").First(&history, id).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
}

// FindAll retrieves all medical history entries with pagination and preloaded relationships
func (r *medicalHistoryRepository) FindAll(page, limit int) ([]models.MedicalHistory, int64, error) {
	var histories []models.MedicalHistory
	var total int64

	// Count total records
	if err := r.db.Model(&models.MedicalHistory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Offset(offset).Limit(limit).Find(&histories).Error
	if err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}

// Update updates an existing medical history entry
func (r *medicalHistoryRepository) Update(ctx context.Context, history *models.MedicalHistory) error {
	return r.db.WithContext(ctx).Save(history).Error
}

// Delete soft deletes a medical history entry by ID
func (r *medicalHistoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MedicalHistory{}, id).Error
}

// FindByPatientID retrieves medical history entries by patient ID with pagination
func (r *medicalHistoryRepository) FindByPatientID(patientID uint, page, limit int) ([]models.MedicalHistory, int64, error) {
	var histories []models.MedicalHistory
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.MedicalHistory{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&histories).Error
	if err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}

// FindByStatus retrieves medical history entries by status with pagination
func (r *medicalHistoryRepository) FindByStatus(status models.MedicalHistoryStatus, page, limit int) ([]models.MedicalHistory, int64, error) {
	var histories []models.MedicalHistory
	var total int64

	// Count total records with status filter
	if err := r.db.Model(&models.MedicalHistory{}).Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Where("status = ?", status).
		Offset(offset).Limit(limit).Find(&histories).Error
	if err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}

// FindByFilters retrieves medical history entries with multiple optional filters
func (r *medicalHistoryRepository) FindByFilters(patientID *uint, status *models.MedicalHistoryStatus, page, limit int) ([]models.MedicalHistory, int64, error) {
	var histories []models.MedicalHistory
	var total int64

	// Build query with filters
	query := r.db.Model(&models.MedicalHistory{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
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
	if status != nil {
		dataQuery = dataQuery.Where("status = ?", *status)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&histories).Error
	if err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}
