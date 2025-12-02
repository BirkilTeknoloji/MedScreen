package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

type appointmentRepository struct {
	db *gorm.DB
}

// NewAppointmentRepository creates a new instance of AppointmentRepository
func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

// Create creates a new appointment in the database
func (r *appointmentRepository) Create(ctx context.Context, appointment *models.Appointment) error {
	return r.db.WithContext(ctx).Create(appointment).Error
}

// FindByID retrieves an appointment by ID with preloaded patient and doctor relationships
func (r *appointmentRepository) FindByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").First(&appointment, id).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

// FindAll retrieves all appointments with pagination and preloaded relationships
func (r *appointmentRepository) FindAll(page, limit int) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	// Count total records
	if err := r.db.Model(&models.Appointment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").
		Offset(offset).Limit(limit).Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

// Update updates an existing appointment
func (r *appointmentRepository) Update(ctx context.Context, appointment *models.Appointment) error {
	return r.db.WithContext(ctx).Save(appointment).Error
}

// Delete soft deletes an appointment by ID
func (r *appointmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Appointment{}, id).Error
}

// FindByDoctorID retrieves appointments by doctor ID with pagination
func (r *appointmentRepository) FindByDoctorID(doctorID uint, page, limit int) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	// Count total records for this doctor
	if err := r.db.Model(&models.Appointment{}).Where("doctor_id = ?", doctorID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").
		Where("doctor_id = ?", doctorID).
		Offset(offset).Limit(limit).Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

// FindByPatientID retrieves appointments by patient ID with pagination
func (r *appointmentRepository) FindByPatientID(patientID uint, page, limit int) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	// Count total records for this patient
	if err := r.db.Model(&models.Appointment{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

// FindByStatus retrieves appointments by status with pagination
func (r *appointmentRepository) FindByStatus(status models.AppointmentStatus, page, limit int) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	// Count total records with this status
	if err := r.db.Model(&models.Appointment{}).Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").
		Where("status = ?", status).
		Offset(offset).Limit(limit).Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

// FindByDateRange retrieves appointments within a date range with pagination
func (r *appointmentRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	// Count total records in date range
	if err := r.db.Model(&models.Appointment{}).
		Where("appointment_date >= ? AND appointment_date <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").
		Where("appointment_date >= ? AND appointment_date <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

// FindByFilters retrieves appointments with multiple optional filters and pagination
func (r *appointmentRepository) FindByFilters(doctorID, patientID *uint, status *models.AppointmentStatus, startDate, endDate *time.Time, page, limit int) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	// Build query with filters
	query := r.db.Model(&models.Appointment{})

	if doctorID != nil {
		query = query.Where("doctor_id = ?", *doctorID)
	}

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if startDate != nil && endDate != nil {
		query = query.Where("appointment_date >= ? AND appointment_date <= ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("appointment_date >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("appointment_date <= ?", *endDate)
	}

	// Count total records matching filters
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := query.Preload("Patient").Preload("Doctor").Preload("CreatedByUser").
		Offset(offset).Limit(limit).Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}
