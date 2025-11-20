package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type diagnosisService struct {
	repo repository.DiagnosisRepository
}

// NewDiagnosisService creates a new instance of DiagnosisService
func NewDiagnosisService(repo repository.DiagnosisRepository) DiagnosisService {
	return &diagnosisService{repo: repo}
}

// CreateDiagnosis creates a new diagnosis with validation
func (s *diagnosisService) CreateDiagnosis(diagnosis *models.Diagnosis) error {
	if diagnosis == nil {
		return errors.New("diagnosis cannot be nil")
	}

	// Validate required fields
	if diagnosis.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if diagnosis.DoctorID == 0 {
		return errors.New("doctor_id is required")
	}
	if diagnosis.DiagnosisDate.IsZero() {
		return errors.New("diagnosis_date is required")
	}
	if diagnosis.DiagnosisName == "" {
		return errors.New("diagnosis_name is required")
	}

	// Validate severity enum
	if err := validateDiagnosisSeverity(diagnosis.Severity); err != nil {
		return err
	}

	// Validate status enum
	if err := validateDiagnosisStatus(diagnosis.Status); err != nil {
		return err
	}

	return s.repo.Create(diagnosis)
}

// GetDiagnosis retrieves a diagnosis by ID
func (s *diagnosisService) GetDiagnosis(id uint) (*models.Diagnosis, error) {
	if id == 0 {
		return nil, errors.New("invalid diagnosis id")
	}

	diagnosis, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if diagnosis == nil {
		return nil, errors.New("diagnosis not found")
	}

	return diagnosis, nil
}

// GetDiagnoses retrieves all diagnoses with pagination
func (s *diagnosisService) GetDiagnoses(page, limit int) ([]models.Diagnosis, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateDiagnosis updates an existing diagnosis with validation
func (s *diagnosisService) UpdateDiagnosis(id uint, diagnosis *models.Diagnosis) error {
	if id == 0 {
		return errors.New("invalid diagnosis id")
	}
	if diagnosis == nil {
		return errors.New("diagnosis cannot be nil")
	}

	// Check if diagnosis exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("diagnosis not found")
	}

	// Validate severity enum if being updated
	if diagnosis.Severity != "" {
		if err := validateDiagnosisSeverity(diagnosis.Severity); err != nil {
			return err
		}
	}

	// Validate status enum if being updated
	if diagnosis.Status != "" {
		if err := validateDiagnosisStatus(diagnosis.Status); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	diagnosis.ID = id

	return s.repo.Update(diagnosis)
}

// DeleteDiagnosis soft deletes a diagnosis
func (s *diagnosisService) DeleteDiagnosis(id uint) error {
	if id == 0 {
		return errors.New("invalid diagnosis id")
	}

	// Check if diagnosis exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("diagnosis not found")
	}

	return s.repo.Delete(id)
}

// GetDiagnosesByFilters retrieves diagnoses with filters
func (s *diagnosisService) GetDiagnosesByFilters(
	patientID, doctorID, appointmentID *uint,
	startDate, endDate *time.Time,
	page, limit int,
) ([]models.Diagnosis, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByFilters(patientID, doctorID, appointmentID, startDate, endDate, page, limit)
}

// validateDiagnosisSeverity validates that the severity is one of the allowed values
func validateDiagnosisSeverity(severity models.DiagnosisSeverity) error {
	switch severity {
	case models.SeverityMild, models.SeverityModerate,
		models.SeveritySevere, models.SeverityCritical:
		return nil
	default:
		return errors.New("invalid severity: must be one of mild, moderate, severe, critical")
	}
}

// validateDiagnosisStatus validates that the status is one of the allowed values
func validateDiagnosisStatus(status models.DiagnosisStatus) error {
	switch status {
	case models.DiagnosisActive, models.DiagnosisUnderObservation,
		models.DiagnosisResolved, models.DiagnosisChronic:
		return nil
	default:
		return errors.New("invalid status: must be one of active, under_observation, resolved, chronic")
	}
}
