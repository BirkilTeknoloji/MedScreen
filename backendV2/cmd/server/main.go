package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"medscreen/internal/config"
	"medscreen/internal/database"
	"medscreen/internal/handler"
	"medscreen/internal/repository"
	"medscreen/internal/routes"
	"medscreen/internal/service"
	"medscreen/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set JWT secret key
	utils.SetJWTSecretKey()

	// Set Gin mode
	gin.SetMode(cfg.Server.GinMode)

	// Initialize database connection
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run database migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories with dependency injection
	userRepo := repository.NewUserRepository(db)
	nfcCardRepo := repository.NewNFCCardRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)
	diagnosisRepo := repository.NewDiagnosisRepository(db)
	prescriptionRepo := repository.NewPrescriptionRepository(db)
	medicalTestRepo := repository.NewMedicalTestRepository(db)
	medicalHistoryRepo := repository.NewMedicalHistoryRepository(db)
	surgeryHistoryRepo := repository.NewSurgeryHistoryRepository(db)
	allergyRepo := repository.NewAllergyRepository(db)
	vitalSignRepo := repository.NewVitalSignRepository(db)
	deviceRepo := repository.NewDeviceRepository(db)
	qrTokenRepo := repository.NewQRTokenRepository(db)

	// Initialize services with repository dependencies
	userService := service.NewUserService(userRepo)
	nfcCardService := service.NewNFCCardService(nfcCardRepo)
	patientService := service.NewPatientService(
		patientRepo,
		appointmentRepo,
		diagnosisRepo,
		prescriptionRepo,
		medicalTestRepo,
		medicalHistoryRepo,
		surgeryHistoryRepo,
		allergyRepo,
		vitalSignRepo,
	)
	appointmentService := service.NewAppointmentService(appointmentRepo)
	diagnosisService := service.NewDiagnosisService(diagnosisRepo)
	prescriptionService := service.NewPrescriptionService(prescriptionRepo)
	medicalTestService := service.NewMedicalTestService(medicalTestRepo)
	medicalHistoryService := service.NewMedicalHistoryService(medicalHistoryRepo)
	surgeryHistoryService := service.NewSurgeryHistoryService(surgeryHistoryRepo)
	allergyService := service.NewAllergyService(allergyRepo)
	vitalSignService := service.NewVitalSignService(vitalSignRepo)
	deviceService := service.NewDeviceService(deviceRepo, patientRepo)
	qrService := service.NewQRService(qrTokenRepo, patientRepo, deviceRepo)

	// Initialize handlers with service dependencies
	handlers := &routes.Handlers{
		User:           handler.NewUserHandler(userService),
		Patient:        handler.NewPatientHandler(patientService),
		Appointment:    handler.NewAppointmentHandler(appointmentService),
		Diagnosis:      handler.NewDiagnosisHandler(diagnosisService),
		Prescription:   handler.NewPrescriptionHandler(prescriptionService),
		MedicalTest:    handler.NewMedicalTestHandler(medicalTestService),
		MedicalHistory: handler.NewMedicalHistoryHandler(medicalHistoryService),
		SurgeryHistory: handler.NewSurgeryHistoryHandler(surgeryHistoryService),
		Allergy:        handler.NewAllergyHandler(allergyService),
		VitalSign:      handler.NewVitalSignHandler(vitalSignService),
		NFCCard:        handler.NewNFCCardHandler(nfcCardService, userService),
		Device:         handler.NewDeviceHandler(deviceService),
		QR:             handler.NewQRHandler(qrService, deviceService),
		Reset:          handler.NewResetHandler(db), //TODO: prodda sil
	}

	// Set up Gin router
	router := gin.Default()

	// Register all routes with middleware
	routes.SetupRoutes(router, handlers, cfg.CORS.AllowedOrigins, cfg.CORS.AllowedMethods, cfg.CORS.AllowedHeaders)

	// Create HTTP server
	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting MedScreen server on %s", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	if err := database.CloseDatabase(db); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	log.Println("Server exited successfully")
}
