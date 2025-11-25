package repository

import (
	"medscreen/internal/models"
	"time"
)

// UserRepository defines the interface for user data access operations
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindAll(page, limit int) ([]models.User, int64, error)
	Update(user *models.User) error
	Delete(id uint) error
	FindByRole(role models.UserRole, page, limit int) ([]models.User, int64, error)
	FindByNFCCardID(nfcCardID uint) (*models.User, error)
	FindByNFCCardUID(cardUID string) (*models.User, error)
}

// NFCCardRepository defines the interface for NFC card data access operations
type NFCCardRepository interface {
	Create(card *models.NFCCard) error
	FindByID(id uint) (*models.NFCCard, error)
	FindAll(page, limit int) ([]models.NFCCard, int64, error)
	Update(card *models.NFCCard) error
	Delete(id uint) error
	FindByCardUID(cardUID string) (*models.NFCCard, error)
	FindByAssignedUserID(userID uint, page, limit int) ([]models.NFCCard, int64, error)
	FindByActiveStatus(isActive bool, page, limit int) ([]models.NFCCard, int64, error)
}

// PatientRepository defines the interface for patient data access operations
type PatientRepository interface {
	Create(patient *models.Patient) error
	FindByID(id uint) (*models.Patient, error)
	FindAll(page, limit int) ([]models.Patient, int64, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
	FindByTCNumber(tcNumber string) (*models.Patient, error)
	SearchByName(name string, page, limit int) ([]models.Patient, int64, error)
}

// AppointmentRepository defines the interface for appointment data access operations
type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
	FindByID(id uint) (*models.Appointment, error)
	FindAll(page, limit int) ([]models.Appointment, int64, error)
	Update(appointment *models.Appointment) error
	Delete(id uint) error
	FindByDoctorID(doctorID uint, page, limit int) ([]models.Appointment, int64, error)
	FindByPatientID(patientID uint, page, limit int) ([]models.Appointment, int64, error)
	FindByStatus(status models.AppointmentStatus, page, limit int) ([]models.Appointment, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Appointment, int64, error)
	FindByFilters(doctorID, patientID *uint, status *models.AppointmentStatus, startDate, endDate *time.Time, page, limit int) ([]models.Appointment, int64, error)
}

// DiagnosisRepository defines the interface for diagnosis data access operations
type DiagnosisRepository interface {
	Create(diagnosis *models.Diagnosis) error
	FindByID(id uint) (*models.Diagnosis, error)
	FindAll(page, limit int) ([]models.Diagnosis, int64, error)
	Update(diagnosis *models.Diagnosis) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.Diagnosis, int64, error)
	FindByDoctorID(doctorID uint, page, limit int) ([]models.Diagnosis, int64, error)
	FindByAppointmentID(appointmentID uint, page, limit int) ([]models.Diagnosis, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Diagnosis, int64, error)
	FindByFilters(patientID, doctorID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.Diagnosis, int64, error)
}

// PrescriptionRepository defines the interface for prescription data access operations
type PrescriptionRepository interface {
	Create(prescription *models.Prescription) error
	FindByID(id uint) (*models.Prescription, error)
	FindAll(page, limit int) ([]models.Prescription, int64, error)
	Update(prescription *models.Prescription) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.Prescription, int64, error)
	FindByDoctorID(doctorID uint, page, limit int) ([]models.Prescription, int64, error)
	FindByStatus(status models.PrescriptionStatus, page, limit int) ([]models.Prescription, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Prescription, int64, error)
	FindByFilters(patientID, doctorID *uint, status *models.PrescriptionStatus, startDate, endDate *time.Time, page, limit int) ([]models.Prescription, int64, error)
}

// MedicalTestRepository defines the interface for medical test data access operations
type MedicalTestRepository interface {
	Create(test *models.MedicalTest) error
	FindByID(id uint) (*models.MedicalTest, error)
	FindAll(page, limit int) ([]models.MedicalTest, int64, error)
	Update(test *models.MedicalTest) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.MedicalTest, int64, error)
	FindByDoctorID(doctorID uint, page, limit int) ([]models.MedicalTest, int64, error)
	FindByTestType(testType models.TestType, page, limit int) ([]models.MedicalTest, int64, error)
	FindByStatus(status models.TestStatus, page, limit int) ([]models.MedicalTest, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.MedicalTest, int64, error)
	FindByFilters(patientID, doctorID *uint, testType *models.TestType, status *models.TestStatus, startDate, endDate *time.Time, page, limit int) ([]models.MedicalTest, int64, error)
}

// MedicalHistoryRepository defines the interface for medical history data access operations
type MedicalHistoryRepository interface {
	Create(history *models.MedicalHistory) error
	FindByID(id uint) (*models.MedicalHistory, error)
	FindAll(page, limit int) ([]models.MedicalHistory, int64, error)
	Update(history *models.MedicalHistory) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.MedicalHistory, int64, error)
	FindByStatus(status models.MedicalHistoryStatus, page, limit int) ([]models.MedicalHistory, int64, error)
	FindByFilters(patientID *uint, status *models.MedicalHistoryStatus, page, limit int) ([]models.MedicalHistory, int64, error)
}

// SurgeryHistoryRepository defines the interface for surgery history data access operations
type SurgeryHistoryRepository interface {
	Create(surgery *models.SurgeryHistory) error
	FindByID(id uint) (*models.SurgeryHistory, error)
	FindAll(page, limit int) ([]models.SurgeryHistory, int64, error)
	Update(surgery *models.SurgeryHistory) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.SurgeryHistory, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.SurgeryHistory, int64, error)
	FindByFilters(patientID *uint, startDate, endDate *time.Time, page, limit int) ([]models.SurgeryHistory, int64, error)
}

// AllergyRepository defines the interface for allergy data access operations
type AllergyRepository interface {
	Create(allergy *models.Allergy) error
	FindByID(id uint) (*models.Allergy, error)
	FindAll(page, limit int) ([]models.Allergy, int64, error)
	Update(allergy *models.Allergy) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.Allergy, int64, error)
	FindBySeverity(severity models.AllergySeverity, page, limit int) ([]models.Allergy, int64, error)
	FindByFilters(patientID *uint, severity *models.AllergySeverity, page, limit int) ([]models.Allergy, int64, error)
}

// VitalSignRepository defines the interface for vital sign data access operations
type VitalSignRepository interface {
	Create(vitalSign *models.VitalSign) error
	FindByID(id uint) (*models.VitalSign, error)
	FindAll(page, limit int) ([]models.VitalSign, int64, error)
	Update(vitalSign *models.VitalSign) error
	Delete(id uint) error
	FindByPatientID(patientID uint, page, limit int) ([]models.VitalSign, int64, error)
	FindByAppointmentID(appointmentID uint, page, limit int) ([]models.VitalSign, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.VitalSign, int64, error)
	FindByFilters(patientID, appointmentID *uint, startDate, endDate *time.Time, page, limit int) ([]models.VitalSign, int64, error)
}

// DeviceRepository defines the interface for device data access operations
type DeviceRepository interface {
	Create(device *models.Device) error
	GetByMAC(mac string) (*models.Device, error)
	GetByID(id uint) (*models.Device, error)
	Update(device *models.Device) error
	Delete(id uint) error
	FindAll(page, limit int) ([]models.Device, int64, error)
	DeleteByMAC(mac string) error
	FindByFilters(roomNumber *string, patientID *uint, page, limit int) ([]models.Device, int64, error)
}
