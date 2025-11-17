package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type surgeryHistoryService struct {
	repo repository.SurgeryHistoryRepository
}

// NewSurgeryHistoryService creates a new instance of SurgeryHistoryService
func NewSurgeryHistoryService(repo repository.SurgeryHistoryRepository) SurgeryHistoryService {
	return &surgeryHistoryService{repo: repo}
}

// CreateSurgeryHistory creates a new surgery history entry with validation
func (s *surgeryHistoryService) CreateSurgeryHistory(surgery *models.SurgeryHistory) error {
	if surgery == nil {
		return errors.New("surgery history cannot be nil")
	}

	// Validate required fields
	if surgery.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if surgery.ProcedureName == "" {
		return errors.New("procedure_name is required")
	}
	if surgery.SurgeryDate.IsZero() {
		return errors.New("surgery_date is required")
	}
	if surgery.AddedByDoctorID == 0 {
		return errors.New("added_by_doctor_id is required")
	}

	// Validate surgery_date is not in the future
	if surgery.SurgeryDate.After(time.Now()) {
		return errors.New("surgery_date cannot be in the future")
	}

	return s.repo.Create(surgery)
}

// GetSurgeryHistory retrieves a surgery history entry by ID
func (s *surgeryHistoryService) GetSurgeryHistory(id uint) (*models.SurgeryHistory, error) {
	if id == 0 {
		return nil, errors.New("invalid surgery history id")
	}

	surgery, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if surgery == nil {
		return nil, errors.New("surgery history not found")
	}

	return surgery, nil
}

// GetSurgeryHistories retrieves all surgery history entries with pagination
func (s *surgeryHistoryService) GetSurgeryHistories(page, limit int) ([]models.SurgeryHistory, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateSurgeryHistory updates an existing surgery history entry with validation
func (s *surgeryHistoryService) UpdateSurgeryHistory(id uint, surgery *models.SurgeryHistory) error {
	if id == 0 {
		return errors.New("invalid surgery history id")
	}
	if surgery == nil {
		return errors.New("surgery history cannot be nil")
	}

	// Check if surgery history exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("surgery history not found")
	}

	// Validate surgery_date is not in the future if being updated
	if !surgery.SurgeryDate.IsZero() && surgery.SurgeryDate.After(time.Now()) {
		return errors.New("surgery_date cannot be in the future")
	}

	// Set the ID to ensure we're updating the correct record
	surgery.ID = id

	return s.repo.Update(surgery)
}

// DeleteSurgeryHistory soft deletes a surgery history entry
func (s *surgeryHistoryService) DeleteSurgeryHistory(id uint) error {
	if id == 0 {
		return errors.New("invalid surgery history id")
	}

	// Check if surgery history exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("surgery history not found")
	}

	return s.repo.Delete(id)
}

// GetSurgeryHistoriesByFilters retrieves surgery history entries with filters
func (s *surgeryHistoryService) GetSurgeryHistoriesByFilters(
	patientID *uint,
	startDate, endDate *time.Time,
	page, limit int,
) ([]models.SurgeryHistory, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByFilters(patientID, startDate, endDate, page, limit)
}
