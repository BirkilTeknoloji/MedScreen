package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

type prescriptionRepository struct {
	db *gorm.DB
}

// NewPrescriptionRepository creates a new instance of PrescriptionRepository
func NewPrescriptionRepository(db *gorm.DB) PrescriptionRepository {
	return &prescriptionRepository{db: db}
}

// Create creates a new prescription in the database
func (r *prescriptionRepository) Create(ctx context.Context, prescription *models.Prescription) error {
	return r.db.WithContext(ctx).Create(prescription).Error
}

// FindByID retrieves a prescription by ID with preloaded relationships
func (r *prescriptionRepository) FindByID(id uint) (*models.Prescription, error) {
	var prescription models.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").First(&prescription, id).Error
	if err != nil {
		return nil, err
	}
	return &prescription, nil
}

// FindAll retrieves all prescriptions with pagination and preloaded relationships
func (r *prescriptionRepository) FindAll(page, limit int) ([]models.Prescription, int64, error) {
	var prescriptions []models.Prescription
	var total int64

	// Count total records
	if err := r.db.Model(&models.Prescription{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Offset(offset).Limit(limit).Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// Update updates an existing prescription
func (r *prescriptionRepository) Update(ctx context.Context, prescription *models.Prescription) error {
	return r.db.WithContext(ctx).Save(prescription).Error
}

// Delete soft deletes a prescription by ID
func (r *prescriptionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Prescription{}, id).Error
}

// FindByPatientID retrieves prescriptions by patient ID with pagination
func (r *prescriptionRepository) FindByPatientID(patientID uint, page, limit int) ([]models.Prescription, int64, error) {
	var prescriptions []models.Prescription
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.Prescription{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// FindByDoctorID retrieves prescriptions by doctor ID with pagination
func (r *prescriptionRepository) FindByDoctorID(doctorID uint, page, limit int) ([]models.Prescription, int64, error) {
	var prescriptions []models.Prescription
	var total int64

	// Count total records with doctor filter
	if err := r.db.Model(&models.Prescription{}).Where("doctor_id = ?", doctorID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("doctor_id = ?", doctorID).
		Offset(offset).Limit(limit).Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// FindByStatus retrieves prescriptions by status with pagination
func (r *prescriptionRepository) FindByStatus(status models.PrescriptionStatus, page, limit int) ([]models.Prescription, int64, error) {
	var prescriptions []models.Prescription
	var total int64

	// Count total records with status filter
	if err := r.db.Model(&models.Prescription{}).Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("status = ?", status).
		Offset(offset).Limit(limit).Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// FindByDateRange retrieves prescriptions within a date range with pagination
func (r *prescriptionRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Prescription, int64, error) {
	var prescriptions []models.Prescription
	var total int64

	// Count total records with date range filter
	if err := r.db.Model(&models.Prescription{}).
		Where("prescribed_date >= ? AND prescribed_date <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("prescribed_date >= ? AND prescribed_date <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// FindByFilters retrieves prescriptions with multiple optional filters
func (r *prescriptionRepository) FindByFilters(patientID, doctorID *uint, status *models.PrescriptionStatus, startDate, endDate *time.Time, page, limit int) ([]models.Prescription, int64, error) {
	var prescriptions []models.Prescription
	var total int64

	// Build query with filters
	query := r.db.Model(&models.Prescription{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if doctorID != nil {
		query = query.Where("doctor_id = ?", *doctorID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if startDate != nil && endDate != nil {
		query = query.Where("prescribed_date >= ? AND prescribed_date <= ?", *startDate, *endDate)
	}

	// Count total records with filters
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Rebuild query for data retrieval with preloaded relationships
	dataQuery := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment")

	if patientID != nil {
		dataQuery = dataQuery.Where("patient_id = ?", *patientID)
	}
	if doctorID != nil {
		dataQuery = dataQuery.Where("doctor_id = ?", *doctorID)
	}
	if status != nil {
		dataQuery = dataQuery.Where("status = ?", *status)
	}
	if startDate != nil && endDate != nil {
		dataQuery = dataQuery.Where("prescribed_date >= ? AND prescribed_date <= ?", *startDate, *endDate)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&prescriptions).Error
	if err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}
