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
}

// SetupRoutes registers all API endpoints
func SetupRoutes(router *gin.Engine, handlers *Handlers, corsOrigins, corsMethods, corsHeaders []string) {
	// Apply middleware
	router.Use(middleware.CORSMiddleware(corsOrigins, corsMethods, corsHeaders))
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())

	// API v1 group
	api := router.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	{
		users.POST("", handlers.User.CreateUser)
		users.GET("", handlers.User.GetUsers)
		users.GET("/:id", handlers.User.GetUser)
		users.PUT("/:id", handlers.User.UpdateUser)
		users.DELETE("/:id", handlers.User.DeleteUser)
	}

	// Patient routes
	patients := api.Group("/patients")
	{
		patients.POST("", handlers.Patient.CreatePatient)
		patients.GET("", handlers.Patient.GetPatients)
		patients.GET("/:id", handlers.Patient.GetPatient)
		patients.PUT("/:id", handlers.Patient.UpdatePatient)
		patients.DELETE("/:id", handlers.Patient.DeletePatient)
		patients.GET("/:id/medical-history", handlers.Patient.GetPatientMedicalHistory)
	}

	// Appointment routes
	appointments := api.Group("/appointments")
	{
		appointments.POST("", handlers.Appointment.CreateAppointment)
		appointments.GET("", handlers.Appointment.GetAppointments)
		appointments.GET("/:id", handlers.Appointment.GetAppointment)
		appointments.PUT("/:id", handlers.Appointment.UpdateAppointment)
		appointments.DELETE("/:id", handlers.Appointment.DeleteAppointment)
	}

	// Diagnosis routes
	diagnoses := api.Group("/diagnoses")
	{
		diagnoses.POST("", handlers.Diagnosis.CreateDiagnosis)
		diagnoses.GET("", handlers.Diagnosis.GetDiagnoses)
		diagnoses.GET("/:id", handlers.Diagnosis.GetDiagnosis)
		diagnoses.PUT("/:id", handlers.Diagnosis.UpdateDiagnosis)
		diagnoses.DELETE("/:id", handlers.Diagnosis.DeleteDiagnosis)
	}

	// Prescription routes
	prescriptions := api.Group("/prescriptions")
	{
		prescriptions.POST("", handlers.Prescription.CreatePrescription)
		prescriptions.GET("", handlers.Prescription.GetPrescriptions)
		prescriptions.GET("/:id", handlers.Prescription.GetPrescription)
		prescriptions.PUT("/:id", handlers.Prescription.UpdatePrescription)
		prescriptions.DELETE("/:id", handlers.Prescription.DeletePrescription)
	}

	// Medical test routes
	medicalTests := api.Group("/medical-tests")
	{
		medicalTests.POST("", handlers.MedicalTest.CreateMedicalTest)
		medicalTests.GET("", handlers.MedicalTest.GetMedicalTests)
		medicalTests.GET("/:id", handlers.MedicalTest.GetMedicalTest)
		medicalTests.PUT("/:id", handlers.MedicalTest.UpdateMedicalTest)
		medicalTests.DELETE("/:id", handlers.MedicalTest.DeleteMedicalTest)
	}

	// Medical history routes
	medicalHistory := api.Group("/medical-history")
	{
		medicalHistory.POST("", handlers.MedicalHistory.CreateMedicalHistory)
		medicalHistory.GET("", handlers.MedicalHistory.GetMedicalHistories)
		medicalHistory.GET("/:id", handlers.MedicalHistory.GetMedicalHistory)
		medicalHistory.PUT("/:id", handlers.MedicalHistory.UpdateMedicalHistory)
		medicalHistory.DELETE("/:id", handlers.MedicalHistory.DeleteMedicalHistory)
	}

	// Surgery history routes
	surgeryHistory := api.Group("/surgery-history")
	{
		surgeryHistory.POST("", handlers.SurgeryHistory.CreateSurgeryHistory)
		surgeryHistory.GET("", handlers.SurgeryHistory.GetSurgeryHistories)
		surgeryHistory.GET("/:id", handlers.SurgeryHistory.GetSurgeryHistory)
		surgeryHistory.PUT("/:id", handlers.SurgeryHistory.UpdateSurgeryHistory)
		surgeryHistory.DELETE("/:id", handlers.SurgeryHistory.DeleteSurgeryHistory)
	}

	// Allergy routes
	allergies := api.Group("/allergies")
	{
		allergies.POST("", handlers.Allergy.CreateAllergy)
		allergies.GET("", handlers.Allergy.GetAllergies)
		allergies.GET("/:id", handlers.Allergy.GetAllergy)
		allergies.PUT("/:id", handlers.Allergy.UpdateAllergy)
		allergies.DELETE("/:id", handlers.Allergy.DeleteAllergy)
	}

	// Vital sign routes
	vitalSigns := api.Group("/vital-signs")
	{
		vitalSigns.POST("", handlers.VitalSign.CreateVitalSign)
		vitalSigns.GET("", handlers.VitalSign.GetVitalSigns)
		vitalSigns.GET("/:id", handlers.VitalSign.GetVitalSign)
		vitalSigns.PUT("/:id", handlers.VitalSign.UpdateVitalSign)
		vitalSigns.DELETE("/:id", handlers.VitalSign.DeleteVitalSign)
	}

	// NFC card routes
	nfcCards := api.Group("/nfc-cards")
	{
		nfcCards.POST("", handlers.NFCCard.CreateCard)
		nfcCards.GET("", handlers.NFCCard.GetCards)
		nfcCards.GET("/:id", handlers.NFCCard.GetCard)
		nfcCards.PUT("/:id", handlers.NFCCard.UpdateCard)
		nfcCards.DELETE("/:id", handlers.NFCCard.DeleteCard)
		nfcCards.POST("/:id/assign", handlers.NFCCard.AssignCard)
		nfcCards.POST("/:id/deactivate", handlers.NFCCard.DeactivateCard)
		nfcCards.POST("/authenticate", handlers.NFCCard.AuthenticateByNFC)
	}
}
