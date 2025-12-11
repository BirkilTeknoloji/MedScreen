package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

type diagnosisRepository struct {
	db *gorm.DB
}

// NewDiagnosisRepository creates a new instance of DiagnosisRepository
func NewDiagnosisRepository(db *gorm.DB) DiagnosisRepository {
	return &diagnosisRepository{db: db}
}

// Create creates a new diagnosis in the database
func (r *diagnosisRepository) Create(ctx context.Context, diagnosis *models.Diagnosis) error {
	return r.db.WithContext(ctx).Create(diagnosis).Error
}

// FindByID retrieves a diagnosis by ID with preloaded relationships
func (r *diagnosisRepository) FindByID(id uint) (*models.Diagnosis, error) {
	var diagnosis models.Diagnosis
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").First(&diagnosis, id).Error
	if err != nil {
		return nil, err
	}
	return &diagnosis, nil
}

// FindAll retrieves all diagnoses with pagination and preloaded relationships
func (r *diagnosisRepository) FindAll(page, limit int) ([]models.Diagnosis, int64, error) {
	var diagnoses []models.Diagnosis
	var total int64

	// Count total records
	if err := r.db.Model(&models.Diagnosis{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Offset(offset).Limit(limit).Find(&diagnoses).Error
	if err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}

// Update updates an existing diagnosis
func (r *diagnosisRepository) Update(ctx context.Context, diagnosis *models.Diagnosis) error {
	return r.db.WithContext(ctx).Save(diagnosis).Error
}

// Delete soft deletes a diagnosis by ID
func (r *diagnosisRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Diagnosis{}, id).Error
}

// FindByPatientID retrieves diagnoses by patient ID with pagination
func (r *diagnosisRepository) FindByPatientID(patientID uint, page, limit int) ([]models.Diagnosis, int64, error) {
	var diagnoses []models.Diagnosis
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.Diagnosis{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&diagnoses).Error
	if err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}

// FindByDoctorID retrieves diagnoses by doctor ID with pagination
func (r *diagnosisRepository) FindByDoctorID(doctorID uint, page, limit int) ([]models.Diagnosis, int64, error) {
	var diagnoses []models.Diagnosis
	var total int64

	// Count total records with doctor filter
	if err := r.db.Model(&models.Diagnosis{}).Where("doctor_id = ?", doctorID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("doctor_id = ?", doctorID).
		Offset(offset).Limit(limit).Find(&diagnoses).Error
	if err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}

// FindByAppointmentID retrieves diagnoses by appointment ID with pagination
func (r *diagnosisRepository) FindByAppointmentID(appointmentID uint, page, limit int) ([]models.Diagnosis, int64, error) {
	var diagnoses []models.Diagnosis
	var total int64

	// Count total records with appointment filter
	if err := r.db.Model(&models.Diagnosis{}).Where("appointment_id = ?", appointmentID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("appointment_id = ?", appointmentID).
		Offset(offset).Limit(limit).Find(&diagnoses).Error
	if err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}

// FindByDateRange retrieves diagnoses within a date range with pagination
func (r *diagnosisRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Diagnosis, int64, error) {
	var diagnoses []models.Diagnosis
	var total int64

	// Count total records with date range filter
	if err := r.db.Model(&models.Diagnosis{}).
		Where("diagnosis_date >= ? AND diagnosis_date <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Appointment").
		Where("diagnosis_date >= ? AND diagnosis_date <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&diagnoses).Error
	if err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}

// FindByFilters retrieves diagnoses with multiple optional filters
func (r *diagnosisRepository) FindByFilters(patientID, doctorID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.Diagnosis, int64, error) {
	var diagnoses []models.Diagnosis
	var total int64

	// Build query with filters
	query := r.db.Model(&models.Diagnosis{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if doctorID != nil {
		query = query.Where("doctor_id = ?", *doctorID)
	}
	if appointmentID != nil {
		query = query.Where("appointment_id = ?", *appointmentID)
	}
	if startDate != nil && endDate != nil {
		query = query.Where("diagnosis_date >= ? AND diagnosis_date <= ?", *startDate, *endDate)
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
	if appointmentID != nil {
		dataQuery = dataQuery.Where("appointment_id = ?", *appointmentID)
	}
	if startDate != nil && endDate != nil {
		dataQuery = dataQuery.Where("diagnosis_date >= ? AND diagnosis_date <= ?", *startDate, *endDate)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&diagnoses).Error
	if err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}
