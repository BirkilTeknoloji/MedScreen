package service

import (
	"context"
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type vitalSignService struct {
	repo repository.VitalSignRepository
}

// NewVitalSignService creates a new instance of VitalSignService
func NewVitalSignService(repo repository.VitalSignRepository) VitalSignService {
	return &vitalSignService{repo: repo}
}

// CreateVitalSign creates a new vital sign entry with validation
func (s *vitalSignService) CreateVitalSign(ctx context.Context, vitalSign *models.VitalSign) error {
	if vitalSign == nil {
		return errors.New("vital sign cannot be nil")
	}

	// Validate required fields
	if vitalSign.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if vitalSign.RecordedByUserID == 0 {
		return errors.New("recorded_by_user_id is required")
	}
	if vitalSign.RecordedAt.IsZero() {
		return errors.New("recorded_at is required")
	}

	// Validate blood pressure (systolic > diastolic)
	if vitalSign.BloodPressureSystolic != nil && vitalSign.BloodPressureDiastolic != nil {
		if *vitalSign.BloodPressureSystolic <= *vitalSign.BloodPressureDiastolic {
			return errors.New("blood_pressure_systolic must be greater than blood_pressure_diastolic")
		}
	}

	// Calculate BMI if height and weight are provided
	if vitalSign.Height != nil && vitalSign.Weight != nil && *vitalSign.Height > 0 {
		bmi := *vitalSign.Weight / ((*vitalSign.Height / 100) * (*vitalSign.Height / 100))
		vitalSign.BMI = &bmi
	}

	return s.repo.Create(ctx, vitalSign)
}

// GetVitalSign retrieves a vital sign entry by ID
func (s *vitalSignService) GetVitalSign(id uint) (*models.VitalSign, error) {
	if id == 0 {
		return nil, errors.New("invalid vital sign id")
	}

	vitalSign, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if vitalSign == nil {
		return nil, errors.New("vital sign not found")
	}

	return vitalSign, nil
}

// GetVitalSigns retrieves all vital sign entries with pagination
func (s *vitalSignService) GetVitalSigns(page, limit int) ([]models.VitalSign, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateVitalSign updates an existing vital sign entry with validation
func (s *vitalSignService) UpdateVitalSign(ctx context.Context, id uint, vitalSign *models.VitalSign) error {
	if id == 0 {
		return errors.New("invalid vital sign id")
	}
	if vitalSign == nil {
		return errors.New("vital sign cannot be nil")
	}

	// Check if vital sign exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("vital sign not found")
	}

	// Validate blood pressure (systolic > diastolic)
	if vitalSign.BloodPressureSystolic != nil && vitalSign.BloodPressureDiastolic != nil {
		if *vitalSign.BloodPressureSystolic <= *vitalSign.BloodPressureDiastolic {
			return errors.New("blood_pressure_systolic must be greater than blood_pressure_diastolic")
		}
	}

	// Calculate BMI if height and weight are provided
	height := vitalSign.Height
	weight := vitalSign.Weight
	if height == nil {
		height = existing.Height
	}
	if weight == nil {
		weight = existing.Weight
	}
	if height != nil && weight != nil && *height > 0 {
		bmi := *weight / ((*height / 100) * (*height / 100))
		vitalSign.BMI = &bmi
	}

	// Set the ID to ensure we're updating the correct record
	vitalSign.ID = id

	return s.repo.Update(ctx, vitalSign)
}

// DeleteVitalSign soft deletes a vital sign entry
func (s *vitalSignService) DeleteVitalSign(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid vital sign id")
	}

	// Check if vital sign exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("vital sign not found")
	}

	return s.repo.Delete(ctx, id)
}

// GetVitalSignsByFilters retrieves vital sign entries with filters
func (s *vitalSignService) GetVitalSignsByFilters(
	patientID, appointmentID *uint,
	startDate, endDate *time.Time,
	page, limit int,
) ([]models.VitalSign, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByFilters(patientID, appointmentID, startDate, endDate, page, limit)
}
