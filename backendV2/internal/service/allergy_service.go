package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type allergyService struct {
	repo repository.AllergyRepository
}

// NewAllergyService creates a new instance of AllergyService
func NewAllergyService(repo repository.AllergyRepository) AllergyService {
	return &allergyService{repo: repo}
}

// CreateAllergy creates a new allergy entry with validation
func (s *allergyService) CreateAllergy(allergy *models.Allergy) error {
	if allergy == nil {
		return errors.New("allergy cannot be nil")
	}

	// Validate required fields
	if allergy.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if allergy.Allergen == "" {
		return errors.New("allergen is required")
	}
	if allergy.Reaction == "" {
		return errors.New("reaction is required")
	}
	if allergy.DiagnosedDate.IsZero() {
		return errors.New("diagnosed_date is required")
	}
	if allergy.AddedByDoctorID == 0 {
		return errors.New("added_by_doctor_id is required")
	}

	// Validate allergy_type enum
	if err := validateAllergyType(allergy.AllergyType); err != nil {
		return err
	}

	// Validate severity enum
	if err := validateAllergySeverity(allergy.Severity); err != nil {
		return err
	}

	return s.repo.Create(allergy)
}

// GetAllergy retrieves an allergy entry by ID
func (s *allergyService) GetAllergy(id uint) (*models.Allergy, error) {
	if id == 0 {
		return nil, errors.New("invalid allergy id")
	}

	allergy, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if allergy == nil {
		return nil, errors.New("allergy not found")
	}

	return allergy, nil
}

// GetAllergies retrieves all allergy entries with pagination
func (s *allergyService) GetAllergies(page, limit int) ([]models.Allergy, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateAllergy updates an existing allergy entry with validation
func (s *allergyService) UpdateAllergy(id uint, allergy *models.Allergy) error {
	if id == 0 {
		return errors.New("invalid allergy id")
	}
	if allergy == nil {
		return errors.New("allergy cannot be nil")
	}

	// Check if allergy exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("allergy not found")
	}

	// Validate allergy_type enum if being updated
	if allergy.AllergyType != "" {
		if err := validateAllergyType(allergy.AllergyType); err != nil {
			return err
		}
	}

	// Validate severity enum if being updated
	if allergy.Severity != "" {
		if err := validateAllergySeverity(allergy.Severity); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	allergy.ID = id

	// Preserve immutable fields - don't allow changing who added it or the patient
	allergy.AddedByDoctorID = existing.AddedByDoctorID
	allergy.PatientID = existing.PatientID

	return s.repo.Update(allergy)
}

// DeleteAllergy soft deletes an allergy entry
func (s *allergyService) DeleteAllergy(id uint) error {
	if id == 0 {
		return errors.New("invalid allergy id")
	}

	// Check if allergy exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("allergy not found")
	}

	return s.repo.Delete(id)
}

// GetAllergiesByFilters retrieves allergy entries with filters
func (s *allergyService) GetAllergiesByFilters(
	patientID *uint,
	severity *models.AllergySeverity,
	page, limit int,
) ([]models.Allergy, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Validate severity enum if provided
	if severity != nil {
		if err := validateAllergySeverity(*severity); err != nil {
			return nil, 0, err
		}
	}

	return s.repo.FindByFilters(patientID, severity, page, limit)
}

// validateAllergyType validates that the allergy type is one of the allowed values
func validateAllergyType(allergyType models.AllergyType) error {
	switch allergyType {
	case models.AllergyMedication, models.AllergyFood,
		models.AllergyEnvironmental, models.AllergyOther:
		return nil
	default:
		return errors.New("invalid allergy_type: must be one of medication, food, environmental, other")
	}
}

// validateAllergySeverity validates that the severity is one of the allowed values
func validateAllergySeverity(severity models.AllergySeverity) error {
	switch severity {
	case models.AllergyMild, models.AllergyModerate,
		models.AllergySevere, models.AllergyLifeThreatening:
		return nil
	default:
		return errors.New("invalid severity: must be one of mild, moderate, severe, life-threatening")
	}
}
