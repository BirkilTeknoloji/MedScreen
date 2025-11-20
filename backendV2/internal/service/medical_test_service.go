package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type medicalTestService struct {
	repo repository.MedicalTestRepository
}

// NewMedicalTestService creates a new instance of MedicalTestService
func NewMedicalTestService(repo repository.MedicalTestRepository) MedicalTestService {
	return &medicalTestService{repo: repo}
}

// CreateMedicalTest creates a new medical test with validation
func (s *medicalTestService) CreateMedicalTest(test *models.MedicalTest) error {
	if test == nil {
		return errors.New("medical test cannot be nil")
	}

	// Validate required fields
	if test.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if test.OrderedByDoctorID == 0 {
		return errors.New("ordered_by_doctor_id is required")
	}
	if test.TestName == "" {
		return errors.New("test_name is required")
	}
	if test.OrderedDate.IsZero() {
		return errors.New("ordered_date is required")
	}

	// Validate test_type enum
	if err := validateTestType(test.TestType); err != nil {
		return err
	}

	// Validate status enum
	if err := validateTestStatus(test.Status); err != nil {
		return err
	}

	return s.repo.Create(test)
}

// GetMedicalTest retrieves a medical test by ID
func (s *medicalTestService) GetMedicalTest(id uint) (*models.MedicalTest, error) {
	if id == 0 {
		return nil, errors.New("invalid medical test id")
	}

	test, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if test == nil {
		return nil, errors.New("medical test not found")
	}

	return test, nil
}

// GetMedicalTests retrieves all medical tests with pagination
func (s *medicalTestService) GetMedicalTests(page, limit int) ([]models.MedicalTest, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateMedicalTest updates an existing medical test with validation
func (s *medicalTestService) UpdateMedicalTest(id uint, test *models.MedicalTest) error {
	if id == 0 {
		return errors.New("invalid medical test id")
	}
	if test == nil {
		return errors.New("medical test cannot be nil")
	}

	// Check if medical test exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("medical test not found")
	}

	// Validate test_type enum if being updated
	if test.TestType != "" {
		if err := validateTestType(test.TestType); err != nil {
			return err
		}
	}

	// Validate status enum if being updated
	if test.Status != "" {
		if err := validateTestStatus(test.Status); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	test.ID = id

	return s.repo.Update(test)
}

// DeleteMedicalTest soft deletes a medical test
func (s *medicalTestService) DeleteMedicalTest(id uint) error {
	if id == 0 {
		return errors.New("invalid medical test id")
	}

	// Check if medical test exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("medical test not found")
	}

	return s.repo.Delete(id)
}

// GetMedicalTestsByFilters retrieves medical tests with filters
func (s *medicalTestService) GetMedicalTestsByFilters(
	patientID, doctorID *uint,
	testType *models.TestType,
	status *models.TestStatus,
	startDate, endDate *time.Time,
	page, limit int,
) ([]models.MedicalTest, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Validate test_type enum if provided
	if testType != nil {
		if err := validateTestType(*testType); err != nil {
			return nil, 0, err
		}
	}

	// Validate status enum if provided
	if status != nil {
		if err := validateTestStatus(*status); err != nil {
			return nil, 0, err
		}
	}

	return s.repo.FindByFilters(patientID, doctorID, testType, status, startDate, endDate, page, limit)
}

// validateTestType validates that the test type is one of the allowed values
func validateTestType(testType models.TestType) error {
	switch testType {
	case models.TestBlood, models.TestXRay, models.TestMRI,
		models.TestCTScan, models.TestUltrasound, models.TestECG,
		models.TestBiopsy:
		return nil
	default:
		return errors.New("invalid test_type: must be one of blood_test, x-ray, mri, ct_scan, ultrasound, ecg, biopsy")
	}
}

// validateTestStatus validates that the status is one of the allowed values
func validateTestStatus(status models.TestStatus) error {
	switch status {
	case models.TestOrdered, models.TestScheduled,
		models.TestInProgress, models.TestCompleted,
		models.TestCancelled:
		return nil
	default:
		return errors.New("invalid status: must be one of ordered, scheduled, in_progress, completed, cancelled")
	}
}
