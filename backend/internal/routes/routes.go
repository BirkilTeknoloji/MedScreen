package routes

import (
	"medscreen/internal/handler"
	"medscreen/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers holds all VEM 2.0 HTTP handlers (read-only)
type Handlers struct {
	Personel              *handler.PersonelHandler
	NFCKart               *handler.NFCKartHandler
	Hasta                 *handler.HastaHandler
	HastaBasvuru          *handler.HastaBasvuruHandler
	Yatak                 *handler.YatakHandler
	TabletCihaz           *handler.TabletCihazHandler
	AnlikYatanHasta       *handler.AnlikYatanHastaHandler
	HastaVitalFizikiBulgu *handler.HastaVitalFizikiBulguHandler
	KlinikSeyir           *handler.KlinikSeyirHandler
	TibbiOrder            *handler.TibbiOrderHandler
	TetkikSonuc           *handler.TetkikSonucHandler
	Recete                *handler.ReceteHandler
	BasvuruTani           *handler.BasvuruTaniHandler
	HastaTibbiBilgi       *handler.HastaTibbiBilgiHandler
	HastaUyari            *handler.HastaUyariHandler
	RiskSkorlama          *handler.RiskSkorlamaHandler
	BasvuruYemek          *handler.BasvuruYemekHandler
}

// MethodNotAllowedMiddleware rejects write operations (POST, PUT, PATCH, DELETE)
// This middleware ensures the API is read-only as per VEM 2.0 requirements
func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPut ||
			method == http.MethodPatch || method == http.MethodDelete {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"success": false,
				"code":    "METHOD_NOT_ALLOWED",
				"message": "This API is read-only. Write operations are not permitted.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// SetupRoutes registers all VEM 2.0 API endpoints (GET only)
func SetupRoutes(router *gin.Engine, handlers *Handlers, corsOrigins, corsMethods, corsHeaders []string) {
	// Apply global middleware
	router.Use(middleware.CORSMiddleware(corsOrigins, corsMethods, corsHeaders))
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())
	router.Use(MethodNotAllowedMiddleware())

	// API v1 group
	api := router.Group("/api/v1")

	// NFC Authentication endpoint (public, GET only for read-only system)
	api.GET("/nfc-kart/authenticate/:kart_uid", handlers.NFCKart.GetByKartUID)

	// Protected routes (require authentication)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Personel routes (GET only)
	personel := protected.Group("/personel")
	{
		personel.GET("", handlers.Personel.GetAll)
		personel.GET("/:kodu", handlers.Personel.GetByKodu)
		personel.GET("/gorev/:gorev_kodu", handlers.Personel.GetByGorev)
		personel.GET("/authenticate/:kart_uid", handlers.Personel.Authenticate)
	}

	// NFC Kart routes (GET only)
	nfcKart := protected.Group("/nfc-kart")
	{
		nfcKart.GET("/:kodu", handlers.NFCKart.GetByKodu)
		nfcKart.GET("/uid/:kart_uid", handlers.NFCKart.GetByKartUID)
		nfcKart.GET("/personel/:personel_kodu", handlers.NFCKart.GetByPersonelKodu)
	}

	// Hasta routes (GET only)
	hasta := protected.Group("/hasta")
	{
		hasta.GET("", handlers.Hasta.GetAll)
		hasta.GET("/search", handlers.Hasta.Search)
		hasta.GET("/:kodu", handlers.Hasta.GetByKodu)
		hasta.GET("/tc/:tc_kimlik", handlers.Hasta.GetByTCKimlik)
	}

	// Hasta Basvuru routes (GET only)
	hastaBasvuru := protected.Group("/hasta-basvuru")
	{
		hastaBasvuru.GET("/filter", handlers.HastaBasvuru.GetByFilters)
		hastaBasvuru.GET("/:kodu", handlers.HastaBasvuru.GetByKodu)
		hastaBasvuru.GET("/hasta/:hasta_kodu", handlers.HastaBasvuru.GetByHasta)
		hastaBasvuru.GET("/hekim/:hekim_kodu", handlers.HastaBasvuru.GetByHekim)
	}

	// Yatak routes (GET only)
	yatak := protected.Group("/yatak")
	{
		yatak.GET("", handlers.Yatak.GetAll)
		yatak.GET("/:kodu", handlers.Yatak.GetByKodu)
		yatak.GET("/birim/:birim_kodu/oda/:oda_kodu", handlers.Yatak.GetByBirimOda)
	}

	// Tablet Cihaz routes (GET only)
	tabletCihaz := protected.Group("/tablet-cihaz")
	{
		tabletCihaz.GET("", handlers.TabletCihaz.GetAll)
		tabletCihaz.GET("/:kodu", handlers.TabletCihaz.GetByKodu)
		tabletCihaz.GET("/yatak/:yatak_kodu", handlers.TabletCihaz.GetByYatak)
	}

	// Anlik Yatan Hasta routes (GET only)
	anlikYatanHasta := protected.Group("/anlik-yatan-hasta")
	{
		anlikYatanHasta.GET("/:kodu", handlers.AnlikYatanHasta.GetByKodu)
		anlikYatanHasta.GET("/yatak/:yatak_kodu", handlers.AnlikYatanHasta.GetByYatak)
		anlikYatanHasta.GET("/hasta/:hasta_kodu", handlers.AnlikYatanHasta.GetByHasta)
		anlikYatanHasta.GET("/birim/:birim_kodu", handlers.AnlikYatanHasta.GetByBirim)
	}

	// Vital Bulgu routes (GET only)
	vitalBulgu := protected.Group("/vital-bulgu")
	{
		vitalBulgu.GET("/date-range", handlers.HastaVitalFizikiBulgu.GetByDateRange)
		vitalBulgu.GET("/:kodu", handlers.HastaVitalFizikiBulgu.GetByKodu)
		vitalBulgu.GET("/basvuru/:basvuru_kodu", handlers.HastaVitalFizikiBulgu.GetByBasvuru)
	}

	// Klinik Seyir routes (GET only)
	klinikSeyir := protected.Group("/klinik-seyir")
	{
		klinikSeyir.GET("/filter", handlers.KlinikSeyir.GetByFilters)
		klinikSeyir.GET("/:kodu", handlers.KlinikSeyir.GetByKodu)
		klinikSeyir.GET("/basvuru/:basvuru_kodu", handlers.KlinikSeyir.GetByBasvuru)
	}

	// Tibbi Order routes (GET only)
	tibbiOrder := protected.Group("/tibbi-order")
	{
		tibbiOrder.GET("/:kodu", handlers.TibbiOrder.GetByKodu)
		tibbiOrder.GET("/:kodu/detay", handlers.TibbiOrder.GetDetay)
		tibbiOrder.GET("/basvuru/:basvuru_kodu", handlers.TibbiOrder.GetByBasvuru)
	}

	// Tetkik Sonuc routes (GET only)
	tetkikSonuc := protected.Group("/tetkik-sonuc")
	{
		tetkikSonuc.GET("/:kodu", handlers.TetkikSonuc.GetByKodu)
		tetkikSonuc.GET("/basvuru/:basvuru_kodu", handlers.TetkikSonuc.GetByBasvuru)
	}

	// Recete routes (GET only)
	recete := protected.Group("/recete")
	{
		recete.GET("/:kodu", handlers.Recete.GetByKodu)
		recete.GET("/:kodu/ilaclar", handlers.Recete.GetIlaclar)
		recete.GET("/basvuru/:basvuru_kodu", handlers.Recete.GetByBasvuru)
		recete.GET("/hekim/:hekim_kodu", handlers.Recete.GetByHekim)
	}

	// Basvuru Tani routes (GET only)
	basvuruTani := protected.Group("/basvuru-tani")
	{
		basvuruTani.GET("/:kodu", handlers.BasvuruTani.GetByKodu)
		basvuruTani.GET("/hasta/:hasta_kodu", handlers.BasvuruTani.GetByHasta)
		basvuruTani.GET("/basvuru/:basvuru_kodu", handlers.BasvuruTani.GetByBasvuru)
	}

	// Hasta Tibbi Bilgi routes (GET only)
	hastaTibbiBilgi := protected.Group("/hasta-tibbi-bilgi")
	{
		hastaTibbiBilgi.GET("/:kodu", handlers.HastaTibbiBilgi.GetByKodu)
		hastaTibbiBilgi.GET("/hasta/:hasta_kodu", handlers.HastaTibbiBilgi.GetByHasta)
		hastaTibbiBilgi.GET("/turu/:turu_kodu", handlers.HastaTibbiBilgi.GetByTuru)
	}

	// Hasta Uyari routes (GET only)
	hastaUyari := protected.Group("/hasta-uyari")
	{
		hastaUyari.GET("/filter", handlers.HastaUyari.GetByFilters)
		hastaUyari.GET("/:kodu", handlers.HastaUyari.GetByKodu)
		hastaUyari.GET("/basvuru/:basvuru_kodu", handlers.HastaUyari.GetByBasvuru)
	}

	// Risk Skorlama routes (GET only)
	riskSkorlama := protected.Group("/risk-skorlama")
	{
		riskSkorlama.GET("/:kodu", handlers.RiskSkorlama.GetByKodu)
		riskSkorlama.GET("/basvuru/:basvuru_kodu", handlers.RiskSkorlama.GetByBasvuru)
		riskSkorlama.GET("/turu/:turu", handlers.RiskSkorlama.GetByTuru)
	}

	// Basvuru Yemek routes (GET only)
	basvuruYemek := protected.Group("/basvuru-yemek")
	{
		basvuruYemek.GET("/:kodu", handlers.BasvuruYemek.GetByKodu)
		basvuruYemek.GET("/basvuru/:basvuru_kodu", handlers.BasvuruYemek.GetByBasvuru)
		basvuruYemek.GET("/turu/:yemek_turu", handlers.BasvuruYemek.GetByTuru)
	}
}
