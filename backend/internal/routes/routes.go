package routes

import (
	"medscreen/internal/handler"
	"medscreen/internal/middleware"
	"medscreen/internal/models"

	"github.com/gin-gonic/gin"
)

// Handlers holds all HTTP handlers
type Handlers struct {
	User           *handler.UserHandler
	Patient        *handler.PatientHandler
	Appointment    *handler.AppointmentHandler
	Diagnosis      *handler.DiagnosisHandler
	Prescription   *handler.PrescriptionHandler
	MedicalTest    *handler.MedicalTestHandler
	MedicalHistory *handler.MedicalHistoryHandler
	SurgeryHistory *handler.SurgeryHistoryHandler
	Allergy        *handler.AllergyHandler
	VitalSign      *handler.VitalSignHandler
	NFCCard        *handler.NFCCardHandler
	Device         *handler.DeviceHandler
	QR             *handler.QRHandler
	Reset          *handler.ResetHandler //TODO: prodda sil
	AuditLog       *handler.AuditLogHandler
}

// SetupRoutes registers all API endpoints
func SetupRoutes(router *gin.Engine, handlers *Handlers, corsOrigins, corsMethods, corsHeaders []string) {
	// Apply middleware
	router.Use(middleware.CORSMiddleware(corsOrigins, corsMethods, corsHeaders))
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())

	// API v1 group
	api := router.Group("/api/v1")
	api.POST("/nfc-cards/authenticate", handlers.NFCCard.AuthenticateByNFC)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// User routes - Admin only
	users := protected.Group("/users")
	users.Use(middleware.RoleMiddleware(models.RoleAdmin))
	{
		users.POST("", handlers.User.CreateUser)
		users.GET("", handlers.User.GetUsers)
		users.GET("/:id", handlers.User.GetUser)
		users.PUT("/:id", handlers.User.UpdateUser)
		users.DELETE("/:id", handlers.User.DeleteUser)
	}

	// Patient routes
	// Read: Admin, Doctor, Nurse, Receptionist
	// Create/Update: Admin, Receptionist
	// Delete: Admin
	patients := protected.Group("/patients")
	{
		patients.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleReceptionist), handlers.Patient.CreatePatient)
		patients.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse, models.RoleReceptionist), handlers.Patient.GetPatients)
		patients.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse, models.RoleReceptionist), handlers.Patient.GetPatient)
		patients.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleReceptionist), handlers.Patient.UpdatePatient)
		patients.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.Patient.DeletePatient)
		patients.GET("/:id/medical-history", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Patient.GetPatientMedicalHistory)
		patients.POST("/:id/generate-qr", middleware.RoleMiddleware(models.RoleAdmin, models.RoleReceptionist), handlers.QR.GeneratePatientQR)
	}

	// Appointment routes
	// Read: Admin, Doctor, Nurse, Receptionist
	// Create/Update: Admin, Doctor, Receptionist
	// Delete: Admin, Receptionist
	appointments := protected.Group("/appointments")
	{
		appointments.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleReceptionist), handlers.Appointment.CreateAppointment)
		appointments.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse, models.RoleReceptionist), handlers.Appointment.GetAppointments)
		appointments.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse, models.RoleReceptionist), handlers.Appointment.GetAppointment)
		appointments.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleReceptionist), handlers.Appointment.UpdateAppointment)
		appointments.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleReceptionist), handlers.Appointment.DeleteAppointment)
	}

	// Diagnosis routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor
	// Delete: Admin
	diagnoses := protected.Group("/diagnoses")
	{
		diagnoses.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor), handlers.Diagnosis.CreateDiagnosis)
		diagnoses.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Diagnosis.GetDiagnoses)
		diagnoses.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Diagnosis.GetDiagnosis)
		diagnoses.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor), handlers.Diagnosis.UpdateDiagnosis)
		diagnoses.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.Diagnosis.DeleteDiagnosis)
	}

	// Prescription routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor
	// Delete: Admin
	prescriptions := protected.Group("/prescriptions")
	{
		prescriptions.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor), handlers.Prescription.CreatePrescription)
		prescriptions.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Prescription.GetPrescriptions)
		prescriptions.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Prescription.GetPrescription)
		prescriptions.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor), handlers.Prescription.UpdatePrescription)
		prescriptions.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.Prescription.DeletePrescription)
	}

	// Medical test routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor, Nurse
	// Delete: Admin
	medicalTests := protected.Group("/medical-tests")
	{
		medicalTests.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalTest.CreateMedicalTest)
		medicalTests.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalTest.GetMedicalTests)
		medicalTests.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalTest.GetMedicalTest)
		medicalTests.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalTest.UpdateMedicalTest)
		medicalTests.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.MedicalTest.DeleteMedicalTest)
	}

	// Medical history routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor, Nurse
	// Delete: Admin
	medicalHistory := protected.Group("/medical-history")
	{
		medicalHistory.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalHistory.CreateMedicalHistory)
		medicalHistory.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalHistory.GetMedicalHistories)
		medicalHistory.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalHistory.GetMedicalHistory)
		medicalHistory.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.MedicalHistory.UpdateMedicalHistory)
		medicalHistory.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.MedicalHistory.DeleteMedicalHistory)
	}

	// Surgery history routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor, Nurse
	// Delete: Admin
	surgeryHistory := protected.Group("/surgery-history")
	{
		surgeryHistory.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.SurgeryHistory.CreateSurgeryHistory)
		surgeryHistory.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.SurgeryHistory.GetSurgeryHistories)
		surgeryHistory.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.SurgeryHistory.GetSurgeryHistory)
		surgeryHistory.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.SurgeryHistory.UpdateSurgeryHistory)
		surgeryHistory.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.SurgeryHistory.DeleteSurgeryHistory)
	}

	// Allergy routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor, Nurse
	// Delete: Admin
	allergies := protected.Group("/allergies")
	{
		allergies.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Allergy.CreateAllergy)
		allergies.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Allergy.GetAllergies)
		allergies.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Allergy.GetAllergy)
		allergies.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.Allergy.UpdateAllergy)
		allergies.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.Allergy.DeleteAllergy)
	}

	// Vital sign routes
	// Read: Admin, Doctor, Nurse
	// Create/Update: Admin, Doctor, Nurse
	// Delete: Admin
	vitalSigns := protected.Group("/vital-signs")
	{
		vitalSigns.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.VitalSign.CreateVitalSign)
		vitalSigns.GET("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.VitalSign.GetVitalSigns)
		vitalSigns.GET("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.VitalSign.GetVitalSign)
		vitalSigns.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleDoctor, models.RoleNurse), handlers.VitalSign.UpdateVitalSign)
		vitalSigns.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), handlers.VitalSign.DeleteVitalSign)
	}

	// NFC card routes - Admin only
	nfcCards := protected.Group("/nfc-cards")
	nfcCards.Use(middleware.RoleMiddleware(models.RoleAdmin))
	{
		nfcCards.POST("", handlers.NFCCard.CreateCard)
		nfcCards.GET("", handlers.NFCCard.GetCards)
		nfcCards.GET("/:id", handlers.NFCCard.GetCard)
		nfcCards.PUT("/:id", handlers.NFCCard.UpdateCard)
		nfcCards.DELETE("/:id", handlers.NFCCard.DeleteCard)
		nfcCards.POST("/:id/assign", handlers.NFCCard.AssignCard)
		nfcCards.POST("/:id/deactivate", handlers.NFCCard.DeactivateCard)
	}

	// Device routes - Admin only
	devices := protected.Group("/devices")
	devices.Use(middleware.RoleMiddleware(models.RoleAdmin))
	{
		devices.POST("", handlers.Device.RegisterDevice)
		devices.GET("", handlers.Device.GetDevices)
		devices.GET("/id/:id", handlers.Device.GetDeviceByID)
		devices.GET("/:mac_address", handlers.Device.GetDeviceByMAC)
		devices.PUT("/:mac_address", handlers.Device.UpdateDevice)
		devices.DELETE("/:mac_address", handlers.Device.DeleteDevice)
		devices.POST("/:mac_address/assign", handlers.Device.AssignPatient)
		devices.POST("/:mac_address/unassign", handlers.Device.UnassignPatient)
		devices.POST("/:mac_address/scan-patient-qr", handlers.QR.ScanPatientQR)
		devices.POST("/:mac_address/generate-prescription-qr/:patient_id", handlers.QR.GeneratePrescriptionQR)
	}

	// QR Token routes
	qrTokens := protected.Group("/qr-tokens")
	{
		qrTokens.GET("/:token/validate", handlers.QR.ValidateQRToken)
	}

	// Audit Log routes - Admin only
	auditLogs := protected.Group("/audit-logs")
	auditLogs.Use(middleware.RoleMiddleware(models.RoleAdmin))
	{
		auditLogs.GET("", handlers.AuditLog.GetAuditLogs)
	}

	// Reset route (prodda sil)
	api.POST("/reset", handlers.Reset.ResetDatabase)
}
