package repository

import (
	"medscreen/internal/models"
	"time"
)

// PersonelRepository defines the read-only interface for personnel data access
type PersonelRepository interface {
	FindByKodu(kodu string) (*models.Personel, error)
	FindAll(page, limit int) ([]models.Personel, int64, error)
	FindByGorevKodu(gorevKodu string, page, limit int) ([]models.Personel, int64, error)
}

// NFCKartRepository defines the read-only interface for NFC card data access
type NFCKartRepository interface {
	FindByKodu(kodu string) (*models.NFCKart, error)
	FindByKartUID(kartUID string) (*models.NFCKart, error)
	FindByPersonelKodu(personelKodu string, page, limit int) ([]models.NFCKart, int64, error)
	FindAll(page, limit int) ([]models.NFCKart, int64, error)
}

// HastaRepository defines the read-only interface for patient data access
type HastaRepository interface {
	FindByKodu(kodu string) (*models.Hasta, error)
	FindByTCKimlik(tcKimlik string) (*models.Hasta, error)
	FindAll(page, limit int) ([]models.Hasta, int64, error)
	SearchByName(name string, page, limit int) ([]models.Hasta, int64, error)
}

// HastaBasvuruRepository defines the read-only interface for patient visit/admission data access
type HastaBasvuruRepository interface {
	FindByKodu(kodu string) (*models.HastaBasvuru, error)
	FindByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaBasvuru, int64, error)
	FindByHekimKodu(hekimKodu string, page, limit int) ([]models.HastaBasvuru, int64, error)
	FindByDurum(durum string, page, limit int) ([]models.HastaBasvuru, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.HastaBasvuru, int64, error)
}

// YatakRepository defines the read-only interface for bed data access
type YatakRepository interface {
	FindByKodu(kodu string) (*models.Yatak, error)
	FindByBirimAndOda(birimKodu, odaKodu string, page, limit int) ([]models.Yatak, int64, error)
	FindAll(page, limit int) ([]models.Yatak, int64, error)
}

// TabletCihazRepository defines the read-only interface for tablet device data access
type TabletCihazRepository interface {
	FindByKodu(kodu string) (*models.TabletCihaz, error)
	FindByYatakKodu(yatakKodu string, page, limit int) ([]models.TabletCihaz, int64, error)
	FindAll(page, limit int) ([]models.TabletCihaz, int64, error)
}

// AnlikYatanHastaRepository defines the read-only interface for current inpatient data access
type AnlikYatanHastaRepository interface {
	FindByKodu(kodu string) (*models.AnlikYatanHasta, error)
	FindByYatakKodu(yatakKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error)
	FindByHastaKodu(hastaKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error)
	FindByBirimKodu(birimKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error)
}

// HastaVitalFizikiBulguRepository defines the read-only interface for vital signs data access
type HastaVitalFizikiBulguRepository interface {
	FindByKodu(kodu string) (*models.HastaVitalFizikiBulgu, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error)
}

// KlinikSeyirRepository defines the read-only interface for clinical progress notes data access
type KlinikSeyirRepository interface {
	FindByKodu(kodu string) (*models.KlinikSeyir, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.KlinikSeyir, int64, error)
	FindBySeyirTipi(seyirTipi string, page, limit int) ([]models.KlinikSeyir, int64, error)
	FindByDateRange(startDate, endDate time.Time, page, limit int) ([]models.KlinikSeyir, int64, error)
}

// TibbiOrderRepository defines the read-only interface for medical orders data access
type TibbiOrderRepository interface {
	FindByKodu(kodu string) (*models.TibbiOrder, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TibbiOrder, int64, error)
	FindDetayByOrderKodu(orderKodu string, page, limit int) ([]models.TibbiOrderDetay, int64, error)
}

// TetkikSonucRepository defines the read-only interface for test results data access
type TetkikSonucRepository interface {
	FindByKodu(kodu string) (*models.TetkikSonuc, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TetkikSonuc, int64, error)
}

// ReceteRepository defines the read-only interface for prescription data access
type ReceteRepository interface {
	FindByKodu(kodu string) (*models.Recete, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Recete, int64, error)
	FindByHekimKodu(hekimKodu string, page, limit int) ([]models.Recete, int64, error)
	FindIlacByReceteKodu(receteKodu string, page, limit int) ([]models.ReceteIlac, int64, error)
}

// BasvuruTaniRepository defines the read-only interface for diagnosis data access
type BasvuruTaniRepository interface {
	FindByKodu(kodu string) (*models.BasvuruTani, error)
	FindByHastaKodu(hastaKodu string, page, limit int) ([]models.BasvuruTani, int64, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruTani, int64, error)
	FindByTaniKodu(taniKodu string, page, limit int) ([]models.BasvuruTani, int64, error)
}

// HastaTibbiBilgiRepository defines the read-only interface for patient medical information data access
type HastaTibbiBilgiRepository interface {
	FindByKodu(kodu string) (*models.HastaTibbiBilgi, error)
	FindByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error)
	FindByTuru(turuKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error)
}

// HastaUyariRepository defines the read-only interface for patient warnings data access
type HastaUyariRepository interface {
	FindByKodu(kodu string) (*models.HastaUyari, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaUyari, int64, error)
	FindByTuru(uyariTuru string, page, limit int) ([]models.HastaUyari, int64, error)
	FindByAktiflik(aktiflik int, page, limit int) ([]models.HastaUyari, int64, error)
}

// RiskSkorlamaRepository defines the read-only interface for risk scoring data access
type RiskSkorlamaRepository interface {
	FindByKodu(kodu string) (*models.RiskSkorlama, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.RiskSkorlama, int64, error)
	FindByTuru(turu string, page, limit int) ([]models.RiskSkorlama, int64, error)
}

// BasvuruYemekRepository defines the read-only interface for diet/meal data access
type BasvuruYemekRepository interface {
	FindByKodu(kodu string) (*models.BasvuruYemek, error)
	FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruYemek, int64, error)
	FindByTuru(yemekTuru string, page, limit int) ([]models.BasvuruYemek, int64, error)
}
