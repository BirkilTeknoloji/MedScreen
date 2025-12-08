package repository

import (
	"context"
	"medscreen/internal/models"
	"time"

	"gorm.io/gorm"
)

type medicalTestRepository struct {
	db *gorm.DB
}

// NewMedicalTestRepository creates a new instance of MedicalTestRepository
func NewMedicalTestRepository(db *gorm.DB) MedicalTestRepository {
	return &medicalTestRepository{db: db}
}

// Create creates a new medical test in the database
func (r *medicalTestRepository) Create(ctx context.Context, test *models.MedicalTest) error {
	return r.db.WithContext(ctx).Create(test).Error
}

// FindByID retrieves a medical test by ID with preloaded relationships
func (r *medicalTestRepository) FindByID(id uint) (*models.MedicalTest, error) {
	var test models.MedicalTest
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").First(&test, id).Error
	if err != nil {
		return nil, err
	}
	return &test, nil
}

// FindAll retrieves all medical tests with pagination and preloaded relationships
func (r *medicalTestRepository) FindAll(page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Count total records
	if err := r.db.Model(&models.MedicalTest{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").
		Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

// Update updates an existing medical test
func (r *medicalTestRepository) Update(ctx context.Context, test *models.MedicalTest) error {
	return r.db.WithContext(ctx).Save(test).Error
}

// Delete soft deletes a medical test by ID
func (r *medicalTestRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MedicalTest{}, id).Error
}

// FindByPatientID retrieves medical tests by patient ID with pagination
func (r *medicalTestRepository) FindByPatientID(patientID uint, page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Count total records with patient filter
	if err := r.db.Model(&models.MedicalTest{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").
		Where("patient_id = ?", patientID).
		Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

// FindByDoctorID retrieves medical tests by doctor ID with pagination
func (r *medicalTestRepository) FindByDoctorID(doctorID uint, page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Count total records with doctor filter
	if err := r.db.Model(&models.MedicalTest{}).Where("ordered_by_doctor_id = ?", doctorID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").
		Where("ordered_by_doctor_id = ?", doctorID).
		Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

// FindByTestType retrieves medical tests by test type with pagination
func (r *medicalTestRepository) FindByTestType(testType models.TestType, page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Count total records with test type filter
	if err := r.db.Model(&models.MedicalTest{}).Where("test_type = ?", testType).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").
		Where("test_type = ?", testType).
		Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

// FindByStatus retrieves medical tests by status with pagination
func (r *medicalTestRepository) FindByStatus(status models.TestStatus, page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Count total records with status filter
	if err := r.db.Model(&models.MedicalTest{}).Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").
		Where("status = ?", status).
		Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

// FindByDateRange retrieves medical tests within a date range with pagination
func (r *medicalTestRepository) FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Count total records with date range filter
	if err := r.db.Model(&models.MedicalTest{}).
		Where("ordered_date >= ? AND ordered_date <= ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded relationships
	err := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment").
		Where("ordered_date >= ? AND ordered_date <= ?", startDate, endDate).
		Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

// FindByFilters retrieves medical tests with multiple optional filters
func (r *medicalTestRepository) FindByFilters(patientID, doctorID *uint, testType *models.TestType, status *models.TestStatus, startDate, endDate *time.Time, page, limit int) ([]models.MedicalTest, int64, error) {
	var tests []models.MedicalTest
	var total int64

	// Build query with filters
	query := r.db.Model(&models.MedicalTest{})

	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if doctorID != nil {
		query = query.Where("ordered_by_doctor_id = ?", *doctorID)
	}
	if testType != nil {
		query = query.Where("test_type = ?", *testType)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if startDate != nil && endDate != nil {
		query = query.Where("ordered_date >= ? AND ordered_date <= ?", *startDate, *endDate)
	}

	// Count total records with filters
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Rebuild query for data retrieval with preloaded relationships
	dataQuery := r.db.Preload("Patient").Preload("OrderedByDoctor").Preload("Appointment")

	if patientID != nil {
		dataQuery = dataQuery.Where("patient_id = ?", *patientID)
	}
	if doctorID != nil {
		dataQuery = dataQuery.Where("ordered_by_doctor_id = ?", *doctorID)
	}
	if testType != nil {
		dataQuery = dataQuery.Where("test_type = ?", *testType)
	}
	if status != nil {
		dataQuery = dataQuery.Where("status = ?", *status)
	}
	if startDate != nil && endDate != nil {
		dataQuery = dataQuery.Where("ordered_date >= ? AND ordered_date <= ?", *startDate, *endDate)
	}

	// Retrieve paginated records
	err := dataQuery.Offset(offset).Limit(limit).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}
