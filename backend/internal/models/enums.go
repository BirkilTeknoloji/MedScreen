package models

// VEM 2.0 uses varchar codes for most type fields instead of enums.
// This file contains minimal enums that may still be useful for the read-only system.

// PersonelGorevKodu represents common personnel role codes in VEM 2.0
type PersonelGorevKodu string

const (
	GorevHekim   PersonelGorevKodu = "HEKIM"
	GorevHemsire PersonelGorevKodu = "HEMSIRE"
	GorevDiger   PersonelGorevKodu = "DIGER"
)

// TibbiBilgiTuruKodu represents medical information type codes in VEM 2.0
type TibbiBilgiTuruKodu string

const (
	TibbiBilgiAlerji      TibbiBilgiTuruKodu = "ALERJI"
	TibbiBilgiGecmisTibbi TibbiBilgiTuruKodu = "GECMIS_TIBBI"
	TibbiBilgiAmeliyat    TibbiBilgiTuruKodu = "AMELIYAT"
)

// AktiflikDurumu represents active status values in VEM 2.0
type AktiflikDurumu int

const (
	Pasif AktiflikDurumu = 0
	Aktif AktiflikDurumu = 1
)
