package models

import "time"

// HastaTibbiBilgi represents patient medical information in the VEM 2.0 schema
// (replaces Allergy, MedicalHistory, SurgeryHistory)
type HastaTibbiBilgi struct {
	HastaTibbiBilgiKodu      string     `gorm:"column:hasta_tibbi_bilgi_kodu;primaryKey" json:"hasta_tibbi_bilgi_kodu"`
	HastaKodu                string     `gorm:"column:hasta_kodu;not null" json:"hasta_kodu"`
	Hasta                    *Hasta     `gorm:"foreignKey:HastaKodu;references:HastaKodu" json:"hasta,omitempty"`
	TibbiBilgiTuruKodu       string     `gorm:"column:tibbi_bilgi_turu_kodu;not null" json:"tibbi_bilgi_turu_kodu"`
	TibbiBilgiAltTuruKodu    *string    `gorm:"column:tibbi_bilgi_alt_turu_kodu" json:"tibbi_bilgi_alt_turu_kodu,omitempty"`
	Aciklama                 *string    `gorm:"column:aciklama;type:text" json:"aciklama,omitempty"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (HastaTibbiBilgi) TableName() string {
	return "hasta_tibbi_bilgi"
}
