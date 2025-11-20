package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type medicalHistoryService struct {
	repo repository.MedicalHistoryRepository
}

// NewMedicalHistoryService creates a new instance of MedicalHistoryService
func NewMedicalHistoryService(repo repository.MedicalHistoryRepository) MedicalHistoryService {
	return &medicalHistoryService{repo: repo}
}

// CreateMedicalHistory creates a new medical history entry with validation
func (s *medicalHistoryService) CreateMedicalHistory(history *models.MedicalHistory) error {
	if history == nil {
		return errors.New("medical history cannot be nil")
	}

	// Validate required fields
	if history.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if history.ConditionName == "" {
		return errors.New("condition_name is required")
	}
	if history.DiagnosedDate.IsZero() {
		return errors.New("diagnosed_date is required")
	}
	if history.AddedByDoctorID == 0 {
		return errors.New("added_by_doctor_id is required")
	}

	// Validate status enum
	if err := validateMedicalHistoryStatus(history.Status); err != nil {
		return err
	}

	return s.repo.Create(history)
}

// GetMedicalHistory retrieves a medical history entry by ID
func (s *medicalHistoryService) GetMedicalHistory(id uint) (*models.MedicalHistory, error) {
	if id == 0 {
		return nil, errors.New("invalid medical history id")
	}

	history, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if history == nil {
		return nil, errors.New("medical history not found")
	}

	return history, nil
}

// GetMedicalHistories retrieves all medical history entries with pagination
func (s *medicalHistoryService) GetMedicalHistories(page, limit int) ([]models.MedicalHistory, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateMedicalHistory updates an existing medical history entry with validation
func (s *medicalHistoryService) UpdateMedicalHistory(id uint, history *models.MedicalHistory) error {
	if id == 0 {
		return errors.New("invalid medical history id")
	}
	if history == nil {
		return errors.New("medical history cannot be nil")
	}

	// Check if medical history exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("medical history not found")
	}

	// Validate status enum if being updated
	if history.Status != "" {
		if err := validateMedicalHistoryStatus(history.Status); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	history.ID = id

	return s.repo.Update(history)
}

// DeleteMedicalHistory soft deletes a medical history entry
func (s *medicalHistoryService) DeleteMedicalHistory(id uint) error {
	if id == 0 {
		return errors.New("invalid medical history id")
	}

	// Check if medical history exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("medical history not found")
	}

	return s.repo.Delete(id)
}

// GetMedicalHistoriesByFilters retrieves medical history entries with filters
func (s *medicalHistoryService) GetMedicalHistoriesByFilters(
	patientID *uint,
	status *models.MedicalHistoryStatus,
	page, limit int,
) ([]models.MedicalHistory, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Validate status enum if provided
	if status != nil {
		if err := validateMedicalHistoryStatus(*status); err != nil {
			return nil, 0, err
		}
	}

	return s.repo.FindByFilters(patientID, status, page, limit)
}

// validateMedicalHistoryStatus validates that the status is one of the allowed values
func validateMedicalHistoryStatus(status models.MedicalHistoryStatus) error {
	switch status {
	case models.HistoryActive, models.HistoryResolved,
		models.HistoryChronic, models.HistoryMonitoring:
		return nil
	default:
		return errors.New("invalid status: must be one of active, resolved, chronic, monitoring")
	}
}
