package models

import "time"

// RiskSkorlama represents risk scoring in the VEM 2.0 schema (new entity)
type RiskSkorlama struct {
	RiskSkorlamaKodu         string        `gorm:"column:risk_skorlama_kodu;primaryKey" json:"risk_skorlama_kodu"`
	HastaBasvuruKodu         string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	RiskSkorlamaTuru         string        `gorm:"column:risk_skorlama_turu;not null" json:"risk_skorlama_turu"`
	RiskSkorlamaToplamPuani  string        `gorm:"column:risk_skorlama_toplam_puani;not null" json:"risk_skorlama_toplam_puani"`
	IslemZamani              time.Time     `gorm:"column:islem_zamani;not null" json:"islem_zamani"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (RiskSkorlama) TableName() string {
	return "risk_skorlama"
}
