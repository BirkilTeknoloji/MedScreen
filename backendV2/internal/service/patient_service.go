package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type patientService struct {
	patientRepo        repository.PatientRepository
	appointmentRepo    repository.AppointmentRepository
	diagnosisRepo      repository.DiagnosisRepository
	prescriptionRepo   repository.PrescriptionRepository
	medicalTestRepo    repository.MedicalTestRepository
	medicalHistoryRepo repository.MedicalHistoryRepository
	surgeryHistoryRepo repository.SurgeryHistoryRepository
	allergyRepo        repository.AllergyRepository
	vitalSignRepo      repository.VitalSignRepository
}

// NewPatientService creates a new instance of PatientService
func NewPatientService(
	patientRepo repository.PatientRepository,
	appointmentRepo repository.AppointmentRepository,
	diagnosisRepo repository.DiagnosisRepository,
	prescriptionRepo repository.PrescriptionRepository,
	medicalTestRepo repository.MedicalTestRepository,
	medicalHistoryRepo repository.MedicalHistoryRepository,
	surgeryHistoryRepo repository.SurgeryHistoryRepository,
	allergyRepo repository.AllergyRepository,
	vitalSignRepo repository.VitalSignRepository,
) PatientService {
	return &patientService{
		patientRepo:        patientRepo,
		appointmentRepo:    appointmentRepo,
		diagnosisRepo:      diagnosisRepo,
		prescriptionRepo:   prescriptionRepo,
		medicalTestRepo:    medicalTestRepo,
		medicalHistoryRepo: medicalHistoryRepo,
		surgeryHistoryRepo: surgeryHistoryRepo,
		allergyRepo:        allergyRepo,
		vitalSignRepo:      vitalSignRepo,
	}
}

// CreatePatient creates a new patient with validation
func (s *patientService) CreatePatient(patient *models.Patient) error {
	if patient == nil {
		return errors.New("patient cannot be nil")
	}

	// Validate required fields
	if patient.FirstName == "" {
		return errors.New("first_name is required")
	}
	if patient.LastName == "" {
		return errors.New("last_name is required")
	}
	if patient.TCNumber == "" {
		return errors.New("tc_number is required")
	}
	if patient.Phone == "" {
		return errors.New("phone is required")
	}

	// Validate TC number format (11 digits)
	if err := validateTCNumber(patient.TCNumber); err != nil {
		return err
	}

	// Validate TC number uniqueness
	existing, err := s.patientRepo.FindByTCNumber(patient.TCNumber)
	if err == nil && existing != nil {
		return errors.New("tc_number already exists")
	}

	// Validate gender enum
	if err := validateGender(patient.Gender); err != nil {
		return err
	}

	return s.patientRepo.Create(patient)
}

// GetPatient retrieves a patient by ID
func (s *patientService) GetPatient(id uint) (*models.Patient, error) {
	if id == 0 {
		return nil, errors.New("invalid patient id")
	}

	patient, err := s.patientRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, errors.New("patient not found")
	}

	return patient, nil
}

// GetPatients retrieves all patients with pagination
func (s *patientService) GetPatients(page, limit int) ([]models.Patient, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.patientRepo.FindAll(page, limit)
}

// UpdatePatient updates an existing patient with validation
func (s *patientService) UpdatePatient(id uint, patient *models.Patient) error {
	if id == 0 {
		return errors.New("invalid patient id")
	}
	if patient == nil {
		return errors.New("patient cannot be nil")
	}

	// Check if patient exists
	existing, err := s.patientRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("patient not found")
	}

	// Validate TC number format if being updated
	if patient.TCNumber != "" {
		if err := validateTCNumber(patient.TCNumber); err != nil {
			return err
		}

		// Validate TC number uniqueness if being changed
		if patient.TCNumber != existing.TCNumber {
			existingByTC, err := s.patientRepo.FindByTCNumber(patient.TCNumber)
			if err == nil && existingByTC != nil {
				return errors.New("tc_number already exists")
			}
		}
	}

	// Validate gender enum if being updated
	if patient.Gender != "" {
		if err := validateGender(patient.Gender); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	patient.ID = id

	return s.patientRepo.Update(patient)
}

// DeletePatient soft deletes a patient
func (s *patientService) DeletePatient(id uint) error {
	if id == 0 {
		return errors.New("invalid patient id")
	}

	// Check if patient exists
	existing, err := s.patientRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("patient not found")
	}

	return s.patientRepo.Delete(id)
}

// GetPatientByTCNumber retrieves a patient by TC number
func (s *patientService) GetPatientByTCNumber(tcNumber string) (*models.Patient, error) {
	if tcNumber == "" {
		return nil, errors.New("tc_number is required")
	}

	if err := validateTCNumber(tcNumber); err != nil {
		return nil, err
	}

	patient, err := s.patientRepo.FindByTCNumber(tcNumber)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, errors.New("patient not found")
	}

	return patient, nil
}

// SearchPatientsByName searches for patients by name
func (s *patientService) SearchPatientsByName(name string, page, limit int) ([]models.Patient, int64, error) {
	if name == "" {
		return nil, 0, errors.New("name is required for search")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.patientRepo.SearchByName(name, page, limit)
}

// GetPatientMedicalHistory retrieves aggregated medical history for a patient
func (s *patientService) GetPatientMedicalHistory(patientID uint) (map[string]interface{}, error) {
	if patientID == 0 {
		return nil, errors.New("invalid patient id")
	}

	// Check if patient exists
	patient, err := s.patientRepo.FindByID(patientID)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, errors.New("patient not found")
	}

	// Aggregate all medical data
	result := make(map[string]interface{})
	result["patient"] = patient

	// Get appointments
	appointments, _, _ := s.appointmentRepo.FindByPatientID(patientID, 1, 100)
	result["appointments"] = appointments

	// Get diagnoses
	diagnoses, _, _ := s.diagnosisRepo.FindByPatientID(patientID, 1, 100)
	result["diagnoses"] = diagnoses

	// Get prescriptions
	prescriptions, _, _ := s.prescriptionRepo.FindByPatientID(patientID, 1, 100)
	result["prescriptions"] = prescriptions

	// Get medical tests
	medicalTests, _, _ := s.medicalTestRepo.FindByPatientID(patientID, 1, 100)
	result["medical_tests"] = medicalTests

	// Get medical history
	medicalHistory, _, _ := s.medicalHistoryRepo.FindByPatientID(patientID, 1, 100)
	result["medical_history"] = medicalHistory

	// Get surgery history
	surgeryHistory, _, _ := s.surgeryHistoryRepo.FindByPatientID(patientID, 1, 100)
	result["surgery_history"] = surgeryHistory

	// Get allergies
	allergies, _, _ := s.allergyRepo.FindByPatientID(patientID, 1, 100)
	result["allergies"] = allergies

	// Get vital signs
	vitalSigns, _, _ := s.vitalSignRepo.FindByPatientID(patientID, 1, 100)
	result["vital_signs"] = vitalSigns

	return result, nil
}

// validateTCNumber validates that the TC number is exactly 11 digits
func validateTCNumber(tcNumber string) error {
	if len(tcNumber) != 11 {
		return errors.New("tc_number must be exactly 11 digits")
	}

	for _, char := range tcNumber {
		if char < '0' || char > '9' {
			return errors.New("tc_number must contain only digits")
		}
	}

	return nil
}

// validateGender validates that the gender is one of the allowed values
func validateGender(gender models.Gender) error {
	switch gender {
	case models.GenderMale, models.GenderFemale:
		return nil
	default:
		return errors.New("invalid gender: must be either 'male' or 'female'")
	}
}
