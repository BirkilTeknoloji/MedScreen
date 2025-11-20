package service

import (
	"medscreen/internal/models"
	"time"
)

// UserService defines the interface for user business logic operations
type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id uint) (*models.User, error)
	GetUsers(page, limit int, role *models.UserRole) ([]models.User, int64, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
	AuthenticateByNFC(nfcCardID string) (*models.User, error)
}

// NFCCardService defines the interface for NFC card business logic operations
type NFCCardService interface {
	CreateCard(card *models.NFCCard) error
	GetCard(id uint) (*models.NFCCard, error)
	GetCards(page, limit int) ([]models.NFCCard, int64, error)
	UpdateCard(id uint, card *models.NFCCard) error
	DeleteCard(id uint) error
	AssignCardToUser(cardID, userID uint) error
	DeactivateCard(cardID uint) error
	GetCardByUID(cardUID string) (*models.NFCCard, error)
}

// PatientService defines the interface for patient business logic operations
type PatientService interface {
	CreatePatient(patient *models.Patient) error
	GetPatient(id uint) (*models.Patient, error)
	GetPatients(page, limit int) ([]models.Patient, int64, error)
	UpdatePatient(id uint, patient *models.Patient) error
	DeletePatient(id uint) error
	GetPatientByTCNumber(tcNumber string) (*models.Patient, error)
	SearchPatientsByName(name string, page, limit int) ([]models.Patient, int64, error)
	GetPatientMedicalHistory(patientID uint) (map[string]interface{}, error)
}

// AppointmentService defines the interface for appointment business logic operations
type AppointmentService interface {
	CreateAppointment(appointment *models.Appointment) error
	GetAppointment(id uint) (*models.Appointment, error)
	GetAppointments(page, limit int) ([]models.Appointment, int64, error)
	UpdateAppointment(id uint, appointment *models.Appointment) error
	DeleteAppointment(id uint) error
	GetAppointmentsByFilters(doctorID, patientID *uint, status *models.AppointmentStatus, startDate, endDate *time.Time, page, limit int) ([]models.Appointment, int64, error)
}

// DiagnosisService defines the interface for diagnosis business logic operations
type DiagnosisService interface {
	CreateDiagnosis(diagnosis *models.Diagnosis) error
	GetDiagnosis(id uint) (*models.Diagnosis, error)
	GetDiagnoses(page, limit int) ([]models.Diagnosis, int64, error)
	UpdateDiagnosis(id uint, diagnosis *models.Diagnosis) error
	DeleteDiagnosis(id uint) error
	GetDiagnosesByFilters(patientID, doctorID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.Diagnosis, int64, error)
}

// PrescriptionService defines the interface for prescription business logic operations
type PrescriptionService interface {
	CreatePrescription(prescription *models.Prescription) error
	GetPrescription(id uint) (*models.Prescription, error)
	GetPrescriptions(page, limit int) ([]models.Prescription, int64, error)
	UpdatePrescription(id uint, prescription *models.Prescription) error
	DeletePrescription(id uint) error
	GetPrescriptionsByFilters(patientID, doctorID *uint, status *models.PrescriptionStatus, startDate, endDate *time.Time, page, limit int) ([]models.Prescription, int64, error)
}

// MedicalTestService defines the interface for medical test business logic operations
type MedicalTestService interface {
	CreateMedicalTest(test *models.MedicalTest) error
	GetMedicalTest(id uint) (*models.MedicalTest, error)
	GetMedicalTests(page, limit int) ([]models.MedicalTest, int64, error)
	UpdateMedicalTest(id uint, test *models.MedicalTest) error
	DeleteMedicalTest(id uint) error
	GetMedicalTestsByFilters(patientID, doctorID *uint, testType *models.TestType, status *models.TestStatus, startDate, endDate *time.Time, page, limit int) ([]models.MedicalTest, int64, error)
}

// MedicalHistoryService defines the interface for medical history business logic operations
type MedicalHistoryService interface {
	CreateMedicalHistory(history *models.MedicalHistory) error
	GetMedicalHistory(id uint) (*models.MedicalHistory, error)
	GetMedicalHistories(page, limit int) ([]models.MedicalHistory, int64, error)
	UpdateMedicalHistory(id uint, history *models.MedicalHistory) error
	DeleteMedicalHistory(id uint) error
	GetMedicalHistoriesByFilters(patientID *uint, status *models.MedicalHistoryStatus, page, limit int) ([]models.MedicalHistory, int64, error)
}

// SurgeryHistoryService defines the interface for surgery history business logic operations
type SurgeryHistoryService interface {
	CreateSurgeryHistory(surgery *models.SurgeryHistory) error
	GetSurgeryHistory(id uint) (*models.SurgeryHistory, error)
	GetSurgeryHistories(page, limit int) ([]models.SurgeryHistory, int64, error)
	UpdateSurgeryHistory(id uint, surgery *models.SurgeryHistory) error
	DeleteSurgeryHistory(id uint) error
	GetSurgeryHistoriesByFilters(patientID *uint, startDate, endDate *time.Time, page, limit int) ([]models.SurgeryHistory, int64, error)
}

// AllergyService defines the interface for allergy business logic operations
type AllergyService interface {
	CreateAllergy(allergy *models.Allergy) error
	GetAllergy(id uint) (*models.Allergy, error)
	GetAllergies(page, limit int) ([]models.Allergy, int64, error)
	UpdateAllergy(id uint, allergy *models.Allergy) error
	DeleteAllergy(id uint) error
	GetAllergiesByFilters(patientID *uint, severity *models.AllergySeverity, page, limit int) ([]models.Allergy, int64, error)
}

// VitalSignService defines the interface for vital sign business logic operations
type VitalSignService interface {
	CreateVitalSign(vitalSign *models.VitalSign) error
	GetVitalSign(id uint) (*models.VitalSign, error)
	GetVitalSigns(page, limit int) ([]models.VitalSign, int64, error)
	UpdateVitalSign(id uint, vitalSign *models.VitalSign) error
	DeleteVitalSign(id uint) error
	GetVitalSignsByFilters(patientID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.VitalSign, int64, error)
}
