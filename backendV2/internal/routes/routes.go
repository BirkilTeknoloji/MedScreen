package routes

import (
	"medscreen/internal/handler"
	"medscreen/internal/middleware"

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
	Reset          *handler.ResetHandler //TODO: prodda sil
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

	// User routes
	users := protected.Group("/users")
	{
		users.POST("", handlers.User.CreateUser)
		users.GET("", handlers.User.GetUsers)
		users.GET("/:id", handlers.User.GetUser)
		users.PUT("/:id", handlers.User.UpdateUser)
		users.DELETE("/:id", handlers.User.DeleteUser)
	}

	// Patient routes
	patients := protected.Group("/patients")
	{
		patients.POST("", handlers.Patient.CreatePatient)
		patients.GET("", handlers.Patient.GetPatients)
		patients.GET("/:id", handlers.Patient.GetPatient)
		patients.PUT("/:id", handlers.Patient.UpdatePatient)
		patients.DELETE("/:id", handlers.Patient.DeletePatient)
		patients.GET("/:id/medical-history", handlers.Patient.GetPatientMedicalHistory)
	}

	// Appointment routes
	appointments := protected.Group("/appointments")
	{
		appointments.POST("", handlers.Appointment.CreateAppointment)
		appointments.GET("", handlers.Appointment.GetAppointments)
		appointments.GET("/:id", handlers.Appointment.GetAppointment)
		appointments.PUT("/:id", handlers.Appointment.UpdateAppointment)
		appointments.DELETE("/:id", handlers.Appointment.DeleteAppointment)
	}

	// Diagnosis routes
	diagnoses := protected.Group("/diagnoses")
	{
		diagnoses.POST("", handlers.Diagnosis.CreateDiagnosis)
		diagnoses.GET("", handlers.Diagnosis.GetDiagnoses)
		diagnoses.GET("/:id", handlers.Diagnosis.GetDiagnosis)
		diagnoses.PUT("/:id", handlers.Diagnosis.UpdateDiagnosis)
		diagnoses.DELETE("/:id", handlers.Diagnosis.DeleteDiagnosis)
	}

	// Prescription routes
	prescriptions := protected.Group("/prescriptions")
	{
		prescriptions.POST("", handlers.Prescription.CreatePrescription)
		prescriptions.GET("", handlers.Prescription.GetPrescriptions)
		prescriptions.GET("/:id", handlers.Prescription.GetPrescription)
		prescriptions.PUT("/:id", handlers.Prescription.UpdatePrescription)
		prescriptions.DELETE("/:id", handlers.Prescription.DeletePrescription)
	}

	// Medical test routes
	medicalTests := protected.Group("/medical-tests")
	{
		medicalTests.POST("", handlers.MedicalTest.CreateMedicalTest)
		medicalTests.GET("", handlers.MedicalTest.GetMedicalTests)
		medicalTests.GET("/:id", handlers.MedicalTest.GetMedicalTest)
		medicalTests.PUT("/:id", handlers.MedicalTest.UpdateMedicalTest)
		medicalTests.DELETE("/:id", handlers.MedicalTest.DeleteMedicalTest)
	}

	// Medical history routes
	medicalHistory := protected.Group("/medical-history")
	{
		medicalHistory.POST("", handlers.MedicalHistory.CreateMedicalHistory)
		medicalHistory.GET("", handlers.MedicalHistory.GetMedicalHistories)
		medicalHistory.GET("/:id", handlers.MedicalHistory.GetMedicalHistory)
		medicalHistory.PUT("/:id", handlers.MedicalHistory.UpdateMedicalHistory)
		medicalHistory.DELETE("/:id", handlers.MedicalHistory.DeleteMedicalHistory)
	}

	// Surgery history routes
	surgeryHistory := protected.Group("/surgery-history")
	{
		surgeryHistory.POST("", handlers.SurgeryHistory.CreateSurgeryHistory)
		surgeryHistory.GET("", handlers.SurgeryHistory.GetSurgeryHistories)
		surgeryHistory.GET("/:id", handlers.SurgeryHistory.GetSurgeryHistory)
		surgeryHistory.PUT("/:id", handlers.SurgeryHistory.UpdateSurgeryHistory)
		surgeryHistory.DELETE("/:id", handlers.SurgeryHistory.DeleteSurgeryHistory)
	}

	// Allergy routes
	allergies := protected.Group("/allergies")
	{
		allergies.POST("", handlers.Allergy.CreateAllergy)
		allergies.GET("", handlers.Allergy.GetAllergies)
		allergies.GET("/:id", handlers.Allergy.GetAllergy)
		allergies.PUT("/:id", handlers.Allergy.UpdateAllergy)
		allergies.DELETE("/:id", handlers.Allergy.DeleteAllergy)
	}

	// Vital sign routes
	vitalSigns := protected.Group("/vital-signs")
	{
		vitalSigns.POST("", handlers.VitalSign.CreateVitalSign)
		vitalSigns.GET("", handlers.VitalSign.GetVitalSigns)
		vitalSigns.GET("/:id", handlers.VitalSign.GetVitalSign)
		vitalSigns.PUT("/:id", handlers.VitalSign.UpdateVitalSign)
		vitalSigns.DELETE("/:id", handlers.VitalSign.DeleteVitalSign)
	}

	// NFC card routes
	nfcCards := protected.Group("/nfc-cards")
	{
		nfcCards.POST("", handlers.NFCCard.CreateCard)
		nfcCards.GET("", handlers.NFCCard.GetCards)
		nfcCards.GET("/:id", handlers.NFCCard.GetCard)
		nfcCards.PUT("/:id", handlers.NFCCard.UpdateCard)
		nfcCards.DELETE("/:id", handlers.NFCCard.DeleteCard)
		nfcCards.POST("/:id/assign", handlers.NFCCard.AssignCard)
		nfcCards.POST("/:id/deactivate", handlers.NFCCard.DeactivateCard)
	}

	// Device routes
	devices := protected.Group("/devices")
	{
		devices.POST("", handlers.Device.RegisterDevice)
		devices.GET("", handlers.Device.GetDevices)
		devices.GET("/id/:id", handlers.Device.GetDeviceByID)
		devices.GET("/:mac_address", handlers.Device.GetDeviceByMAC)
		devices.PUT("/:mac_address", handlers.Device.UpdateDevice)
		devices.DELETE("/:mac_address", handlers.Device.DeleteDevice)
		devices.POST("/:mac_address/assign", handlers.Device.AssignPatient)
		devices.POST("/:mac_address/unassign", handlers.Device.UnassignPatient)
	}

	// Reset route (prodda sil)
	api.POST("/reset", handlers.Reset.ResetDatabase)
}
