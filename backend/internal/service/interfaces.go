package service

import (
	"medscreen/internal/models"
	"time"
)

// PersonelService defines the read-only interface for personnel business logic operations
type PersonelService interface {
	GetByKodu(kodu string) (*models.Personel, error)
	GetAll(page, limit int) ([]models.Personel, int64, error)
	GetByGorevKodu(gorevKodu string, page, limit int) ([]models.Personel, int64, error)
	AuthenticateByNFC(kartUID string) (*models.Personel, error)
}

// NFCKartService defines the read-only interface for NFC card business logic operations
type NFCKartService interface {
	GetByKodu(kodu string) (*models.NFCKart, error)
	GetByKartUID(kartUID string) (*models.NFCKart, error)
	GetByPersonelKodu(personelKodu string, page, limit int) ([]models.NFCKart, int64, error)
}

// HastaService defines the read-only interface for patient business logic operations
type HastaService interface {
	GetByKodu(kodu string) (*models.Hasta, error)
	GetByTCKimlik(tcKimlik string) (*models.Hasta, error)
	GetAll(page, limit int) ([]models.Hasta, int64, error)
	SearchByAdSoyadi(ad, soyadi string, page, limit int) ([]models.Hasta, int64, error)
}

// HastaBasvuruService defines the read-only interface for patient visit business logic operations
type HastaBasvuruService interface {
	GetByKodu(kodu string) (*models.HastaBasvuru, error)
	GetByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaBasvuru, int64, error)
	GetByHekimKodu(hekimKodu string, page, limit int) ([]models.HastaBasvuru, int64, error)
	GetByFilters(durum *string, startDate, endDate *time.Time, page, limit int) ([]models.HastaBasvuru, int64, error)
}

// YatakService defines the read-only interface for bed business logic operations
type YatakService interface {
	GetByKodu(kodu string) (*models.Yatak, error)
	GetByBirimAndOda(birimKodu, odaKodu string, page, limit int) ([]models.Yatak, int64, error)
	GetAll(page, limit int) ([]models.Yatak, int64, error)
}

// TabletCihazService defines the read-only interface for tablet device business logic operations
type TabletCihazService interface {
	GetByKodu(kodu string) (*models.TabletCihaz, error)
	GetByYatakKodu(yatakKodu string, page, limit int) ([]models.TabletCihaz, int64, error)
	GetAll(page, limit int) ([]models.TabletCihaz, int64, error)
}

// AnlikYatanHastaService defines the read-only interface for current inpatient business logic operations
type AnlikYatanHastaService interface {
	GetByKodu(kodu string) (*models.AnlikYatanHasta, error)
	GetByYatakKodu(yatakKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error)
	GetByHastaKodu(hastaKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error)
	GetByBirimKodu(birimKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error)
}

// HastaVitalFizikiBulguService defines the read-only interface for vital signs business logic operations
type HastaVitalFizikiBulguService interface {
	GetByKodu(kodu string) (*models.HastaVitalFizikiBulgu, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error)
	GetByDateRange(startDate, endDate time.Time, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error)
}

// KlinikSeyirService defines the read-only interface for clinical progress notes business logic operations
type KlinikSeyirService interface {
	GetByKodu(kodu string) (*models.KlinikSeyir, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.KlinikSeyir, int64, error)
	GetByFilters(seyirTipi *string, startDate, endDate *time.Time, page, limit int) ([]models.KlinikSeyir, int64, error)
}

// TibbiOrderService defines the read-only interface for medical orders business logic operations
type TibbiOrderService interface {
	GetByKodu(kodu string) (*models.TibbiOrder, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TibbiOrder, int64, error)
	GetDetayByOrderKodu(orderKodu string, page, limit int) ([]models.TibbiOrderDetay, int64, error)
}

// TetkikSonucService defines the read-only interface for test results business logic operations
type TetkikSonucService interface {
	GetByKodu(kodu string) (*models.TetkikSonuc, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TetkikSonuc, int64, error)
}

// ReceteService defines the read-only interface for prescription business logic operations
type ReceteService interface {
	GetByKodu(kodu string) (*models.Recete, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Recete, int64, error)
	GetByHekimKodu(hekimKodu string, page, limit int) ([]models.Recete, int64, error)
	GetIlaclar(receteKodu string, page, limit int) ([]models.ReceteIlac, int64, error)
}

// BasvuruTaniService defines the read-only interface for diagnosis business logic operations
type BasvuruTaniService interface {
	GetByKodu(kodu string) (*models.BasvuruTani, error)
	GetByHastaKodu(hastaKodu string, page, limit int) ([]models.BasvuruTani, int64, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruTani, int64, error)
}

// HastaTibbiBilgiService defines the read-only interface for patient medical information business logic operations
type HastaTibbiBilgiService interface {
	GetByKodu(kodu string) (*models.HastaTibbiBilgi, error)
	GetByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error)
	GetByTuru(turuKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error)
}

// HastaUyariService defines the read-only interface for patient warnings business logic operations
type HastaUyariService interface {
	GetByKodu(kodu string) (*models.HastaUyari, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaUyari, int64, error)
	GetByFilters(uyariTuru *string, aktiflik *int, page, limit int) ([]models.HastaUyari, int64, error)
}

// RiskSkorlamaService defines the read-only interface for risk scoring business logic operations
type RiskSkorlamaService interface {
	GetByKodu(kodu string) (*models.RiskSkorlama, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.RiskSkorlama, int64, error)
	GetByTuru(turu string, page, limit int) ([]models.RiskSkorlama, int64, error)
}

// BasvuruYemekService defines the read-only interface for diet/meal business logic operations
type BasvuruYemekService interface {
	GetByKodu(kodu string) (*models.BasvuruYemek, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruYemek, int64, error)
	GetByTuru(yemekTuru string, page, limit int) ([]models.BasvuruYemek, int64, error)
}

// RandevuService defines the read-only interface for appointment business logic operations
type RandevuService interface {
	GetByKodu(kodu string) (*models.Randevu, error)
	GetByHastaKodu(hastaKodu string, page, limit int) ([]models.Randevu, int64, error)
	GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Randevu, int64, error)
	GetByHekimKodu(hekimKodu string, page, limit int) ([]models.Randevu, int64, error)
	GetByTuru(randevuTuru string, page, limit int) ([]models.Randevu, int64, error)
	GetByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Randevu, int64, error)
}
