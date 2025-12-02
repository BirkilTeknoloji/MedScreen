package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

type vitalSignRepository struct {
	db *gorm.DB
}

// NewVitalSignRepository creates a new instance of VitalSignRepository
func NewVitalSignRepository(db *gorm.DB) VitalSignRepository {
	return &vitalSignRepository{db: db}
}

// Create creates a new vital sign entry in the database
func (r *vitalSignRepository) Create(ctx context.Context, vitalSign *models.VitalSign) error {
	return r.db.WithContext(ctx).Create(vitalSign).Error
}

// FindByID retrieves a vital sign entry by ID with preloaded relationships
func (r *vitalSignRepository) FindByID(id uint) (*models.VitalSign, error) {
	var vitalSign models.VitalSign
	err := r.db.Preload("Patient").Preload("Appointment").Preload("RecordedByUser").First(&vitalSign, id).Error
	if err != nil {
		return nil, err
	}
	return &vitalSign, nil
}

// FindAll retrieves all vital sign entries with pagination and preloaded relationships
func (r *vitalSignRepository) FindAll(page, limit int) ([]models.VitalSign, int64, error) {
	var vitalSigns []models.VitalSign
	var total int64

	// Count total records
	if err := r.db.Model(&models.VitalSign{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Appointment").Preload("RecordedByUser").
		Offset(offset).Limit(limit).Find(&vitalSigns).Error
	if err != nil {
		return nil, 0, err
	}

	return vitalSigns, total, nil
}

// Update updates an existing vital sign entry
func (r *vitalSignRepository) Update(ctx context.Context, vitalSign *models.VitalSign) error {
	return r.db.WithContext(ctx).Save(vitalSign).Error
}

// Delete deletes a vital sign entry by ID
func (r *vitalSignRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VitalSign{}, id).Error
}

// FindByPatientID retrieves vital sign entries by patient ID with pagination
func (r *vitalSignRepository) FindByPatientID(patientID uint, page, limit int) ([]models.VitalSign, int64, error) {
	var vitalSigns []models.VitalSign
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.VitalSign{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Appointment").Preload("RecordedByUser").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&vitalSigns).Error
	if err != nil {
		return nil, 0, err
	}

	return vitalSigns, total, nil
}

// FindByAppointmentID retrieves vital sign entries by appointment ID with pagination
func (r *vitalSignRepository) FindByAppointmentID(appointmentID uint, page, limit int) ([]models.VitalSign, int64, error) {
	var vitalSigns []models.VitalSign
	var total int64

	// Count total records with appointment filter
	if err := r.db.Model(&models.VitalSign{}).Where("appointment_id = ?", appointmentID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Appointment").Preload("RecordedByUser").
		Where("appointment_id = ?", appointmentID).
		Offset(offset).Limit(limit).Find(&vitalSigns).Error
	if err != nil {
		return nil, 0, err
	}

	return vitalSigns, total, nil
}

// FindByDateRange retrieves vital sign entries within a date range with pagination
func (r *vitalSignRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.VitalSign, int64, error) {
	var vitalSigns []models.VitalSign
	var total int64

	// Count total records with date range filter
	if err := r.db.Model(&models.VitalSign{}).
		Where("recorded_at >= ? AND recorded_at <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Appointment").Preload("RecordedByUser").
		Where("recorded_at >= ? AND recorded_at <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&vitalSigns).Error
	if err != nil {
		return nil, 0, err
	}

	return vitalSigns, total, nil
}

// FindByFilters retrieves vital sign entries with multiple optional filters
func (r *vitalSignRepository) FindByFilters(patientID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.VitalSign, int64, error) {
	var vitalSigns []models.VitalSign
	var total int64

	// Build query with filters
	query := r.db.Model(&models.VitalSign{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if appointmentID != nil {
		query = query.Where("appointment_id = ?", *appointmentID)
	}
	if startDate != nil && endDate != nil {
		query = query.Where("recorded_at >= ? AND recorded_at <= ?", *startDate, *endDate)
	}

	// Count total records with filters
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Rebuild query for data retrieval with preloaded relationships
	dataQuery := r.db.Preload("Patient").Preload("Appointment").Preload("RecordedByUser")

	if patientID != nil {
		dataQuery = dataQuery.Where("patient_id = ?", *patientID)
	}
	if appointmentID != nil {
		dataQuery = dataQuery.Where("appointment_id = ?", *appointmentID)
	}
	if startDate != nil && endDate != nil {
		dataQuery = dataQuery.Where("recorded_at >= ? AND recorded_at <= ?", *startDate, *endDate)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&vitalSigns).Error
	if err != nil {
		return nil, 0, err
	}

	return vitalSigns, total, nil
}
