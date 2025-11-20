package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

// NewPatientRepository creates a new instance of PatientRepository
func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db: db}
}

// Create creates a new patient in the database
func (r *patientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

// FindByID retrieves a patient by ID with preloaded primary doctor relationship
func (r *patientRepository) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Preload("PrimaryDoctor").First(&patient, id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// FindAll retrieves all patients with pagination and preloaded primary doctor relationship
func (r *patientRepository) FindAll(page, limit int) ([]models.Patient, int64, error) {
	var patients []models.Patient
	var total int64

	// Count total records
	if err := r.db.Model(&models.Patient{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded primary doctor
	err := r.db.Preload("PrimaryDoctor").Offset(offset).Limit(limit).Find(&patients).Error
	if err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}

// Update updates an existing patient
func (r *patientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

// Delete soft deletes a patient by ID
func (r *patientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}

// FindByTCNumber retrieves a patient by TC number with preloaded primary doctor relationship
func (r *patientRepository) FindByTCNumber(tcNumber string) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Preload("PrimaryDoctor").Where("tc_number = ?", tcNumber).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// SearchByName searches for patients by first name or last name with pagination
func (r *patientRepository) SearchByName(name string, page, limit int) ([]models.Patient, int64, error) {
	var patients []models.Patient
	var total int64

	// Build search query with ILIKE for case-insensitive search
	searchPattern := "%" + name + "%"
	query := r.db.Model(&models.Patient{}).Where("first_name ILIKE ? OR last_name ILIKE ?", searchPattern, searchPattern)

	// Count total records matching search
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Retrieve paginated records with preloaded primary doctor
	err := r.db.Preload("PrimaryDoctor").
		Where("first_name ILIKE ? OR last_name ILIKE ?", searchPattern, searchPattern).
		Offset(offset).Limit(limit).Find(&patients).Error
	if err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}
