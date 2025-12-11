package service

import (
	"context"
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type appointmentService struct {
	repo repository.AppointmentRepository
}

// NewAppointmentService creates a new instance of AppointmentService
func NewAppointmentService(repo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{repo: repo}
}

// CreateAppointment creates a new appointment with validation
func (s *appointmentService) CreateAppointment(ctx context.Context, appointment *models.Appointment) error {
	if appointment == nil {
		return errors.New("appointment cannot be nil")
	}

	// Validate required fields
	if appointment.PatientID == 0 {
		return errors.New("patient_id is required")
	}
	if appointment.DoctorID == 0 {
		return errors.New("doctor_id is required")
	}
	if appointment.AppointmentDate.IsZero() {
		return errors.New("appointment_date is required")
	}
	if appointment.DurationMinutes <= 0 {
		return errors.New("duration_minutes must be positive")
	}
	if appointment.CreatedByUserID == 0 {
		return errors.New("created_by_user_id is required")
	}

	// Validate appointment_type enum
	if err := validateAppointmentType(appointment.AppointmentType); err != nil {
		return err
	}

	// Validate status enum
	if err := validateAppointmentStatus(appointment.Status); err != nil {
		return err
	}

	// Check for scheduling conflicts
	if err := s.checkSchedulingConflicts(appointment); err != nil {
		return err
	}

	return s.repo.Create(ctx, appointment)
}

// GetAppointment retrieves an appointment by ID
func (s *appointmentService) GetAppointment(id uint) (*models.Appointment, error) {
	if id == 0 {
		return nil, errors.New("invalid appointment id")
	}

	appointment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if appointment == nil {
		return nil, errors.New("appointment not found")
	}

	return appointment, nil
}

// GetAppointments retrieves all appointments with pagination
func (s *appointmentService) GetAppointments(page, limit int) ([]models.Appointment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateAppointment updates an existing appointment with validation
func (s *appointmentService) UpdateAppointment(ctx context.Context, id uint, appointment *models.Appointment) error {
	if id == 0 {
		return errors.New("invalid appointment id")
	}
	if appointment == nil {
		return errors.New("appointment cannot be nil")
	}

	// Check if appointment exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("appointment not found")
	}

	// Validate appointment_type enum if being updated
	if appointment.AppointmentType != "" {
		if err := validateAppointmentType(appointment.AppointmentType); err != nil {
			return err
		}
	}

	// Validate status enum if being updated
	if appointment.Status != "" {
		if err := validateAppointmentStatus(appointment.Status); err != nil {
			return err
		}
	}

	// Check for scheduling conflicts if date/time is being changed
	if !appointment.AppointmentDate.IsZero() && !appointment.AppointmentDate.Equal(existing.AppointmentDate) {
		if err := s.checkSchedulingConflicts(appointment); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	appointment.ID = id

	return s.repo.Update(ctx, appointment)
}

// DeleteAppointment soft deletes an appointment
func (s *appointmentService) DeleteAppointment(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid appointment id")
	}

	// Check if appointment exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("appointment not found")
	}

	return s.repo.Delete(ctx, id)
}

// GetAppointmentsByFilters retrieves appointments with filters
func (s *appointmentService) GetAppointmentsByFilters(
	doctorID, patientID *uint,
	status *models.AppointmentStatus,
	startDate, endDate *time.Time,
	page, limit int,
) ([]models.Appointment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Validate status enum if provided
	if status != nil {
		if err := validateAppointmentStatus(*status); err != nil {
			return nil, 0, err
		}
	}

	return s.repo.FindByFilters(doctorID, patientID, status, startDate, endDate, page, limit)
}

// checkSchedulingConflicts checks if the appointment conflicts with existing appointments
func (s *appointmentService) checkSchedulingConflicts(appointment *models.Appointment) error {
	// Calculate appointment end time
	endTime := appointment.AppointmentDate.Add(time.Duration(appointment.DurationMinutes) * time.Minute)

	// Check for overlapping appointments for the same doctor
	startDate := appointment.AppointmentDate.Add(-24 * time.Hour)
	endDate := endTime.Add(24 * time.Hour)

	appointments, _, err := s.repo.FindByFilters(
		&appointment.DoctorID,
		nil,
		nil,
		&startDate,
		&endDate,
		1,
		100,
	)
	if err != nil {
		return err
	}

	for _, existing := range appointments {
		// Skip if it's the same appointment (for updates)
		if existing.ID == appointment.ID {
			continue
		}

		// Skip cancelled or no-show appointments
		if existing.Status == models.StatusCancelled || existing.Status == models.StatusNoShow {
			continue
		}

		existingEnd := existing.AppointmentDate.Add(time.Duration(existing.DurationMinutes) * time.Minute)

		// Check for overlap
		if appointment.AppointmentDate.Before(existingEnd) && endTime.After(existing.AppointmentDate) {
			return errors.New("appointment conflicts with existing appointment for this doctor")
		}
	}

	return nil
}

// validateAppointmentType validates that the appointment type is one of the allowed values
func validateAppointmentType(appointmentType models.AppointmentType) error {
	switch appointmentType {
	case models.TypeConsultation, models.TypeCheckup,
		models.TypeFollowUp, models.TypeEmergency:
		return nil
	default:
		return errors.New("invalid appointment_type: must be one of consultation, checkup, follow-up, emergency")
	}
}

// validateAppointmentStatus validates that the status is one of the allowed values
func validateAppointmentStatus(status models.AppointmentStatus) error {
	switch status {
	case models.StatusScheduled, models.StatusConfirmed,
		models.StatusCompleted, models.StatusCancelled,
		models.StatusNoShow:
		return nil
	default:
		return errors.New("invalid status: must be one of scheduled, confirmed, completed, cancelled, no-show")
	}
}
