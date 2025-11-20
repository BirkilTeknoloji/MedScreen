package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type prescriptionService struct {
	repo repository.PrescriptionRepository
}

// NewPrescriptionService creates a new instance of PrescriptionService
func NewPrescriptionService(repo repository.PrescriptionRepository) PrescriptionService {
	return &prescriptionService{repo: repo}
}

// CreatePrescription creates a new prescription with validation
func (s *prescriptionService) CreatePrescription(prescription *models.Prescription) error {
	if prescription == nil {
		return errors.New("prescription cannot be nil")
	}

	// Validate required fields
	if prescription.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if prescription.DoctorID == 0 {
		return errors.New("doctor_id is required")
	}
	if prescription.PrescribedDate.IsZero() {
		return errors.New("prescribed_date is required")
	}
	if prescription.MedicationName == "" {
		return errors.New("medication_name is required")
	}
	if prescription.Dosage == "" {
		return errors.New("dosage is required")
	}
	if prescription.Frequency == "" {
		return errors.New("frequency is required")
	}
	if prescription.Duration == "" {
		return errors.New("duration is required")
	}

	// Validate quantity is positive
	if prescription.Quantity <= 0 {
		return errors.New("quantity must be a positive integer")
	}

	// Validate status enum
	if err := validatePrescriptionStatus(prescription.Status); err != nil {
		return err
	}

	return s.repo.Create(prescription)
}

// GetPrescription retrieves a prescription by ID
func (s *prescriptionService) GetPrescription(id uint) (*models.Prescription, error) {
	if id == 0 {
		return nil, errors.New("invalid prescription id")
	}

	prescription, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if prescription == nil {
		return nil, errors.New("prescription not found")
	}

	return prescription, nil
}

// GetPrescriptions retrieves all prescriptions with pagination
func (s *prescriptionService) GetPrescriptions(page, limit int) ([]models.Prescription, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdatePrescription updates an existing prescription with validation
func (s *prescriptionService) UpdatePrescription(id uint, prescription *models.Prescription) error {
	if id == 0 {
		return errors.New("invalid prescription id")
	}
	if prescription == nil {
		return errors.New("prescription cannot be nil")
	}

	// Check if prescription exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("prescription not found")
	}

	// Validate quantity is positive if being updated
	if prescription.Quantity != 0 && prescription.Quantity <= 0 {
		return errors.New("quantity must be a positive integer")
	}

	// Validate status enum if being updated
	if prescription.Status != "" {
		if err := validatePrescriptionStatus(prescription.Status); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	prescription.ID = id

	return s.repo.Update(prescription)
}

// DeletePrescription soft deletes a prescription
func (s *prescriptionService) DeletePrescription(id uint) error {
	if id == 0 {
		return errors.New("invalid prescription id")
	}

	// Check if prescription exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("prescription not found")
	}

	return s.repo.Delete(id)
}

// GetPrescriptionsByFilters retrieves prescriptions with filters
func (s *prescriptionService) GetPrescriptionsByFilters(
	patientID, doctorID *uint,
	status *models.PrescriptionStatus,
	startDate, endDate *time.Time,
	page, limit int,
) ([]models.Prescription, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Validate status enum if provided
	if status != nil {
		if err := validatePrescriptionStatus(*status); err != nil {
			return nil, 0, err
		}
	}

	return s.repo.FindByFilters(patientID, doctorID, status, startDate, endDate, page, limit)
}

// validatePrescriptionStatus validates that the status is one of the allowed values
func validatePrescriptionStatus(status models.PrescriptionStatus) error {
	switch status {
	case models.PrescriptionActive, models.PrescriptionCompleted,
		models.PrescriptionCancelled, models.PrescriptionExpired:
		return nil
	default:
		return errors.New("invalid status: must be one of active, completed, cancelled, expired")
	}
}
