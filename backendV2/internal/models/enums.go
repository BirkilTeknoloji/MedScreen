package models

// UserRole represents the role of a user in the system
type UserRole string

const (
	RoleDoctor       UserRole = "doctor"
	RoleNurse        UserRole = "nurse"
	RoleReceptionist UserRole = "receptionist"
	RoleAdmin        UserRole = "admin"
)

// AppointmentType represents the type of appointment
type AppointmentType string

const (
	TypeConsultation AppointmentType = "consultation"
	TypeCheckup      AppointmentType = "checkup"
	TypeFollowUp     AppointmentType = "follow-up"
	TypeEmergency    AppointmentType = "emergency"
)

// AppointmentStatus represents the status of an appointment
type AppointmentStatus string

const (
	StatusScheduled AppointmentStatus = "scheduled"
	StatusConfirmed AppointmentStatus = "confirmed"
	StatusCompleted AppointmentStatus = "completed"
	StatusCancelled AppointmentStatus = "cancelled"
	StatusNoShow    AppointmentStatus = "no-show"
)

// DiagnosisSeverity represents the severity level of a diagnosis
type DiagnosisSeverity string

const (
	SeverityMild     DiagnosisSeverity = "mild"
	SeverityModerate DiagnosisSeverity = "moderate"
	SeveritySevere   DiagnosisSeverity = "severe"
	SeverityCritical DiagnosisSeverity = "critical"
)

// DiagnosisStatus represents the status of a diagnosis
type DiagnosisStatus string

const (
	DiagnosisActive           DiagnosisStatus = "active"
	DiagnosisUnderObservation DiagnosisStatus = "under_observation"
	DiagnosisResolved         DiagnosisStatus = "resolved"
	DiagnosisChronic          DiagnosisStatus = "chronic"
)

// TestType represents the type of medical test
type TestType string

const (
	TestBlood      TestType = "blood_test"
	TestXRay       TestType = "x-ray"
	TestMRI        TestType = "mri"
	TestCTScan     TestType = "ct_scan"
	TestUltrasound TestType = "ultrasound"
	TestECG        TestType = "ecg"
	TestBiopsy     TestType = "biopsy"
)

// TestStatus represents the status of a medical test
type TestStatus string

const (
	TestOrdered    TestStatus = "ordered"
	TestScheduled  TestStatus = "scheduled"
	TestInProgress TestStatus = "in_progress"
	TestCompleted  TestStatus = "completed"
	TestCancelled  TestStatus = "cancelled"
)

// AllergyType represents the type of allergy
type AllergyType string

const (
	AllergyMedication    AllergyType = "medication"
	AllergyFood          AllergyType = "food"
	AllergyEnvironmental AllergyType = "environmental"
	AllergyOther         AllergyType = "other"
)

// AllergySeverity represents the severity of an allergy
type AllergySeverity string

const (
	AllergyMild            AllergySeverity = "mild"
	AllergyModerate        AllergySeverity = "moderate"
	AllergySevere          AllergySeverity = "severe"
	AllergyLifeThreatening AllergySeverity = "life-threatening"
)

// MedicalHistoryStatus represents the status of a medical history entry
type MedicalHistoryStatus string

const (
	HistoryActive     MedicalHistoryStatus = "active"
	HistoryResolved   MedicalHistoryStatus = "resolved"
	HistoryChronic    MedicalHistoryStatus = "chronic"
	HistoryMonitoring MedicalHistoryStatus = "monitoring"
)

// PrescriptionStatus represents the status of a prescription
type PrescriptionStatus string

const (
	PrescriptionActive    PrescriptionStatus = "active"
	PrescriptionCompleted PrescriptionStatus = "completed"
	PrescriptionCancelled PrescriptionStatus = "cancelled"
	PrescriptionExpired   PrescriptionStatus = "expired"
)

// Gender represents the gender of a patient
type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

// QRType represents the type of QR code token
type QRType string

const (
	QRTypePatientAssignment QRType = "patient_assignment"
	QRTypePrescriptionInfo  QRType = "prescription_info"
)
