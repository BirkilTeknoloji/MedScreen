package models

import "time"

// HastaUyari represents patient warnings/alerts in the VEM 2.0 schema (new entity)
type HastaUyari struct {
	HastaUyariKodu           string        `gorm:"column:hasta_uyari_kodu;primaryKey" json:"hasta_uyari_kodu"`
	HastaBasvuruKodu         string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	UyariTuru                string        `gorm:"column:uyari_turu;not null" json:"uyari_turu"`
	UyariAciklama            *string       `gorm:"column:uyari_aciklama;type:text" json:"uyari_aciklama,omitempty"`
	AktiflikBilgisi          int           `gorm:"column:aktiflik_bilgisi;default:1" json:"aktiflik_bilgisi"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (HastaUyari) TableName() string {
	return "hasta_uyari"
}
