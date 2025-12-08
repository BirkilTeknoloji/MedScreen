package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

type surgeryHistoryRepository struct {
	db *gorm.DB
}

// NewSurgeryHistoryRepository creates a new instance of SurgeryHistoryRepository
func NewSurgeryHistoryRepository(db *gorm.DB) SurgeryHistoryRepository {
	return &surgeryHistoryRepository{db: db}
}

// Create creates a new surgery history entry in the database
func (r *surgeryHistoryRepository) Create(ctx context.Context, surgery *models.SurgeryHistory) error {
	return r.db.WithContext(ctx).Create(surgery).Error
}

// FindByID retrieves a surgery history entry by ID with preloaded relationships
func (r *surgeryHistoryRepository) FindByID(id uint) (*models.SurgeryHistory, error) {
	var surgery models.SurgeryHistory
	err := r.db.Preload("Patient").Preload("AddedByDoctor").First(&surgery, id).Error
	if err != nil {
		return nil, err
	}
	return &surgery, nil
}

// FindAll retrieves all surgery history entries with pagination and preloaded relationships
func (r *surgeryHistoryRepository) FindAll(page, limit int) ([]models.SurgeryHistory, int64, error) {
	var surgeries []models.SurgeryHistory
	var total int64

	// Count total records
	if err := r.db.Model(&models.SurgeryHistory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Offset(offset).Limit(limit).Find(&surgeries).Error
	if err != nil {
		return nil, 0, err
	}

	return surgeries, total, nil
}

// Update updates an existing surgery history entry
func (r *surgeryHistoryRepository) Update(ctx context.Context, surgery *models.SurgeryHistory) error {
	return r.db.WithContext(ctx).Save(surgery).Error
}

// Delete soft deletes a surgery history entry by ID
func (r *surgeryHistoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.SurgeryHistory{}, id).Error
}

// FindByPatientID retrieves surgery history entries by patient ID with pagination
func (r *surgeryHistoryRepository) FindByPatientID(patientID uint, page, limit int) ([]models.SurgeryHistory, int64, error) {
	var surgeries []models.SurgeryHistory
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.SurgeryHistory{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&surgeries).Error
	if err != nil {
		return nil, 0, err
	}

	return surgeries, total, nil
}

// FindByDateRange retrieves surgery history entries within a date range with pagination
func (r *surgeryHistoryRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.SurgeryHistory, int64, error) {
	var surgeries []models.SurgeryHistory
	var total int64

	// Count total records with date range filter
	if err := r.db.Model(&models.SurgeryHistory{}).
		Where("surgery_date >= ? AND surgery_date <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("AddedByDoctor").
		Where("surgery_date >= ? AND surgery_date <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&surgeries).Error
	if err != nil {
		return nil, 0, err
	}

	return surgeries, total, nil
}

// FindByFilters retrieves surgery history entries with multiple optional filters
func (r *surgeryHistoryRepository) FindByFilters(patientID *uint, startDate, endDate *time.Time, page, limit int) ([]models.SurgeryHistory, int64, error) {
	var surgeries []models.SurgeryHistory
	var total int64

	// Build query with filters
	query := r.db.Model(&models.SurgeryHistory{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if startDate != nil && endDate != nil {
		query = query.Where("surgery_date >= ? AND surgery_date <= ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("surgery_date >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("surgery_date <= ?", *endDate)
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
	if startDate != nil && endDate != nil {
		dataQuery = dataQuery.Where("surgery_date >= ? AND surgery_date <= ?", *startDate, *endDate)
	} else if startDate != nil {
		dataQuery = dataQuery.Where("surgery_date >= ?", *startDate)
	} else if endDate != nil {
		dataQuery = dataQuery.Where("surgery_date <= ?", *endDate)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&surgeries).Error
	if err != nil {
		return nil, 0, err
	}

	return surgeries, total, nil
}
