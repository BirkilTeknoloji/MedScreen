package models

import "time"

// Yatak represents a bed in the VEM 2.0 schema (new entity)
type Yatak struct {
	YatakKodu                string     `gorm:"column:yatak_kodu;primaryKey" json:"yatak_kodu"`
	BirimKodu                string     `gorm:"column:birim_kodu;not null" json:"birim_kodu"`
	OdaKodu                  string     `gorm:"column:oda_kodu;not null" json:"oda_kodu"`
	YatakAdi                 *string    `gorm:"column:yatak_adi" json:"yatak_adi,omitempty"`
	YatakTuruKodu            *string    `gorm:"column:yatak_turu_kodu" json:"yatak_turu_kodu,omitempty"`
	YogunBakimYatakSeviyesi  *string    `gorm:"column:yogun_bakim_yatak_seviyesi" json:"yogun_bakim_yatak_seviyesi,omitempty"`
	VentilatorCihazKodu      *string    `gorm:"column:ventilator_cihaz_kodu" json:"ventilator_cihaz_kodu,omitempty"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (Yatak) TableName() string {
	return "yatak"
}
