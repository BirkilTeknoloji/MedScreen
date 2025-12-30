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

	// Initialize database connection (read-only mode for VEM 2.0)
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Note: No migrations run - VEM 2.0 tables already exist in the database
	// The database.RunMigrations call has been removed for read-only mode

	// Initialize VEM 2.0 repositories (read-only)
	personelRepo := repository.NewPersonelRepository(db)
	nfcKartRepo := repository.NewNFCKartRepository(db)
	hastaRepo := repository.NewHastaRepository(db)
	hastaBasvuruRepo := repository.NewHastaBasvuruRepository(db)
	yatakRepo := repository.NewYatakRepository(db)
	tabletCihazRepo := repository.NewTabletCihazRepository(db)
	anlikYatanHastaRepo := repository.NewAnlikYatanHastaRepository(db)
	hastaVitalFizikiBulguRepo := repository.NewHastaVitalFizikiBulguRepository(db)
	klinikSeyirRepo := repository.NewKlinikSeyirRepository(db)
	tibbiOrderRepo := repository.NewTibbiOrderRepository(db)
	tetkikSonucRepo := repository.NewTetkikSonucRepository(db)
	receteRepo := repository.NewReceteRepository(db)
	basvuruTaniRepo := repository.NewBasvuruTaniRepository(db)
	hastaTibbiBilgiRepo := repository.NewHastaTibbiBilgiRepository(db)
	hastaUyariRepo := repository.NewHastaUyariRepository(db)
	riskSkorlamaRepo := repository.NewRiskSkorlamaRepository(db)
	basvuruYemekRepo := repository.NewBasvuruYemekRepository(db)
	randevuRepo := repository.NewRandevuRepository(db)

	// Initialize VEM 2.0 services (read-only)
	personelService := service.NewPersonelService(personelRepo, nfcKartRepo)
	nfcKartService := service.NewNFCKartService(nfcKartRepo)
	hastaService := service.NewHastaService(hastaRepo)
	hastaBasvuruService := service.NewHastaBasvuruService(hastaBasvuruRepo)
	yatakService := service.NewYatakService(yatakRepo)
	tabletCihazService := service.NewTabletCihazService(tabletCihazRepo)
	anlikYatanHastaService := service.NewAnlikYatanHastaService(anlikYatanHastaRepo)
	hastaVitalFizikiBulguService := service.NewHastaVitalFizikiBulguService(hastaVitalFizikiBulguRepo)
	klinikSeyirService := service.NewKlinikSeyirService(klinikSeyirRepo)
	tibbiOrderService := service.NewTibbiOrderService(tibbiOrderRepo)
	tetkikSonucService := service.NewTetkikSonucService(tetkikSonucRepo)
	receteService := service.NewReceteService(receteRepo)
	basvuruTaniService := service.NewBasvuruTaniService(basvuruTaniRepo)
	hastaTibbiBilgiService := service.NewHastaTibbiBilgiService(hastaTibbiBilgiRepo)
	hastaUyariService := service.NewHastaUyariService(hastaUyariRepo)
	riskSkorlamaService := service.NewRiskSkorlamaService(riskSkorlamaRepo)
	basvuruYemekService := service.NewBasvuruYemekService(basvuruYemekRepo)
	randevuService := service.NewRandevuService(randevuRepo)

	// Initialize VEM 2.0 handlers (read-only, GET endpoints only)
	handlers := &routes.Handlers{
		Personel:              handler.NewPersonelHandler(personelService),
		NFCKart:               handler.NewNFCKartHandler(nfcKartService),
		Hasta:                 handler.NewHastaHandler(hastaService),
		HastaBasvuru:          handler.NewHastaBasvuruHandler(hastaBasvuruService),
		Yatak:                 handler.NewYatakHandler(yatakService),
		TabletCihaz:           handler.NewTabletCihazHandler(tabletCihazService),
		AnlikYatanHasta:       handler.NewAnlikYatanHastaHandler(anlikYatanHastaService),
		HastaVitalFizikiBulgu: handler.NewHastaVitalFizikiBulguHandler(hastaVitalFizikiBulguService),
		KlinikSeyir:           handler.NewKlinikSeyirHandler(klinikSeyirService),
		TibbiOrder:            handler.NewTibbiOrderHandler(tibbiOrderService),
		TetkikSonuc:           handler.NewTetkikSonucHandler(tetkikSonucService),
		Recete:                handler.NewReceteHandler(receteService),
		BasvuruTani:           handler.NewBasvuruTaniHandler(basvuruTaniService),
		HastaTibbiBilgi:       handler.NewHastaTibbiBilgiHandler(hastaTibbiBilgiService),
		HastaUyari:            handler.NewHastaUyariHandler(hastaUyariService),
		RiskSkorlama:          handler.NewRiskSkorlamaHandler(riskSkorlamaService),
		BasvuruYemek:          handler.NewBasvuruYemekHandler(basvuruYemekService),
		Randevu:               handler.NewRandevuHandler(randevuService),
	}

	// Set up Gin router
	router := gin.Default()

	// Register all VEM 2.0 routes with middleware (GET only)
	routes.SetupRoutes(router, handlers, cfg.CORS.AllowedOrigins, cfg.CORS.AllowedMethods, cfg.CORS.AllowedHeaders)

	// Create HTTP server
	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting MedScreen VEM 2.0 server on %s (read-only mode)", serverAddr)
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
