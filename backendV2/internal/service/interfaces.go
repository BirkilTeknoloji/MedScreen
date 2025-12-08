package service

import (
	"context"
	"medscreen/internal/models"
	"time"
)

// UserService defines the interface for user business logic operations
type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(id uint) (*models.User, error)
	GetUsers(page, limit int, role *models.UserRole) ([]models.User, int64, error)
	UpdateUser(ctx context.Context, id uint, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	AuthenticateByNFC(nfcCardID string) (*models.User, error)
}

// NFCCardService defines the interface for NFC card business logic operations
type NFCCardService interface {
	CreateCard(ctx context.Context, card *models.NFCCard) error
	GetCard(id uint) (*models.NFCCard, error)
	GetCards(page, limit int) ([]models.NFCCard, int64, error)
	UpdateCard(ctx context.Context, id uint, card *models.NFCCard) error
	DeleteCard(ctx context.Context, id uint) error
	AssignCardToUser(ctx context.Context, cardID, userID uint) error
	DeactivateCard(ctx context.Context, cardID uint) error
	GetCardByUID(cardUID string) (*models.NFCCard, error)
}

// PatientService defines the interface for patient business logic operations
type PatientService interface {
	CreatePatient(ctx context.Context, patient *models.Patient) error
	GetPatient(id uint) (*models.Patient, error)
	GetPatients(page, limit int) ([]models.Patient, int64, error)
	UpdatePatient(ctx context.Context, id uint, patient *models.Patient) error
	DeletePatient(ctx context.Context, id uint) error
	GetPatientByTCNumber(tcNumber string) (*models.Patient, error)
	SearchPatientsByName(name string, page, limit int) ([]models.Patient, int64, error)
	GetPatientMedicalHistory(patientID uint) (map[string]interface{}, error)
}

// AppointmentService defines the interface for appointment business logic operations
type AppointmentService interface {
	CreateAppointment(ctx context.Context, appointment *models.Appointment) error
	GetAppointment(id uint) (*models.Appointment, error)
	GetAppointments(page, limit int) ([]models.Appointment, int64, error)
	UpdateAppointment(ctx context.Context, id uint, appointment *models.Appointment) error
	DeleteAppointment(ctx context.Context, id uint) error
	GetAppointmentsByFilters(doctorID, patientID *uint, status *models.AppointmentStatus, startDate, endDate *time.Time, page, limit int) ([]models.Appointment, int64, error)
}

// DiagnosisService defines the interface for diagnosis business logic operations
type DiagnosisService interface {
	CreateDiagnosis(ctx context.Context, diagnosis *models.Diagnosis) error
	GetDiagnosis(id uint) (*models.Diagnosis, error)
	GetDiagnoses(page, limit int) ([]models.Diagnosis, int64, error)
	UpdateDiagnosis(ctx context.Context, id uint, diagnosis *models.Diagnosis) error
	DeleteDiagnosis(ctx context.Context, id uint) error
	GetDiagnosesByFilters(patientID, doctorID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.Diagnosis, int64, error)
}

// PrescriptionService defines the interface for prescription business logic operations
type PrescriptionService interface {
	CreatePrescription(ctx context.Context, prescription *models.Prescription) error
	GetPrescription(id uint) (*models.Prescription, error)
	GetPrescriptions(page, limit int) ([]models.Prescription, int64, error)
	UpdatePrescription(ctx context.Context, id uint, prescription *models.Prescription) error
	DeletePrescription(ctx context.Context, id uint) error
	GetPrescriptionsByFilters(patientID, doctorID *uint, status *models.PrescriptionStatus, startDate, endDate *time.Time, page, limit int) ([]models.Prescription, int64, error)
}

// MedicalTestService defines the interface for medical test business logic operations
type MedicalTestService interface {
	CreateMedicalTest(ctx context.Context, test *models.MedicalTest) error
	GetMedicalTest(id uint) (*models.MedicalTest, error)
	GetMedicalTests(page, limit int) ([]models.MedicalTest, int64, error)
	UpdateMedicalTest(ctx context.Context, id uint, test *models.MedicalTest) error
	DeleteMedicalTest(ctx context.Context, id uint) error
	GetMedicalTestsByFilters(patientID, doctorID *uint, testType *models.TestType, status *models.TestStatus, startDate, endDate *time.Time, page, limit int) ([]models.MedicalTest, int64, error)
}

// MedicalHistoryService defines the interface for medical history business logic operations
type MedicalHistoryService interface {
	CreateMedicalHistory(ctx context.Context, history *models.MedicalHistory) error
	GetMedicalHistory(id uint) (*models.MedicalHistory, error)
	GetMedicalHistories(page, limit int) ([]models.MedicalHistory, int64, error)
	UpdateMedicalHistory(ctx context.Context, id uint, history *models.MedicalHistory) error
	DeleteMedicalHistory(ctx context.Context, id uint) error
	GetMedicalHistoriesByFilters(patientID *uint, status *models.MedicalHistoryStatus, page, limit int) ([]models.MedicalHistory, int64, error)
}

// SurgeryHistoryService defines the interface for surgery history business logic operations
type SurgeryHistoryService interface {
	CreateSurgeryHistory(ctx context.Context, surgery *models.SurgeryHistory) error
	GetSurgeryHistory(id uint) (*models.SurgeryHistory, error)
	GetSurgeryHistories(page, limit int) ([]models.SurgeryHistory, int64, error)
	UpdateSurgeryHistory(ctx context.Context, id uint, surgery *models.SurgeryHistory) error
	DeleteSurgeryHistory(ctx context.Context, id uint) error
	GetSurgeryHistoriesByFilters(patientID *uint, startDate, endDate *time.Time, page, limit int) ([]models.SurgeryHistory, int64, error)
}

// AllergyService defines the interface for allergy business logic operations
type AllergyService interface {
	CreateAllergy(ctx context.Context, allergy *models.Allergy) error
	GetAllergy(id uint) (*models.Allergy, error)
	GetAllergies(page, limit int) ([]models.Allergy, int64, error)
	UpdateAllergy(ctx context.Context, id uint, allergy *models.Allergy) error
	DeleteAllergy(ctx context.Context, id uint) error
	GetAllergiesByFilters(patientID *uint, severity *models.AllergySeverity, page, limit int) ([]models.Allergy, int64, error)
}

// VitalSignService defines the interface for vital sign business logic operations
type VitalSignService interface {
	CreateVitalSign(ctx context.Context, vitalSign *models.VitalSign) error
	GetVitalSign(id uint) (*models.VitalSign, error)
	GetVitalSigns(page, limit int) ([]models.VitalSign, int64, error)
	UpdateVitalSign(ctx context.Context, id uint, vitalSign *models.VitalSign) error
	DeleteVitalSign(ctx context.Context, id uint) error
	GetVitalSignsByFilters(patientID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.VitalSign, int64, error)
}

// DeviceService defines the interface for device business logic operations
type DeviceService interface {
	RegisterDevice(ctx context.Context, device *models.Device) error
	GetDeviceByMAC(mac string) (*models.Device, error)
	GetDeviceByID(id uint) (*models.Device, error)
	GetAllDevices(page, limit int) ([]models.Device, int64, error)
	AssignPatient(ctx context.Context, mac string, patientID uint) error
	UnassignPatient(ctx context.Context, mac string) error
	UpdateDevice(ctx context.Context, mac string, updates *DeviceUpdateRequest) (*models.Device, error)
	DeleteDevice(ctx context.Context, mac string) error
	GetDevicesByFilters(roomNumber *string, patientID *uint, page, limit int) ([]models.Device, int64, error)
}
