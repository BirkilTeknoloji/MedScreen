package models

import "time"

// TabletCihaz represents a tablet device in the VEM 2.0 schema (replaces Device)
type TabletCihaz struct {
	TabletCihazKodu          string     `gorm:"column:tablet_cihaz_kodu;primaryKey" json:"tablet_cihaz_kodu"`
	YatakKodu                *string    `gorm:"column:yatak_kodu;index" json:"yatak_kodu,omitempty"`
	Yatak                    *Yatak     `gorm:"foreignKey:YatakKodu;references:YatakKodu" json:"yatak,omitempty"`
	IPAdresi                 *string    `gorm:"column:ip_adresi" json:"ip_adresi,omitempty"`
	SeriNumarasi             *string    `gorm:"column:seri_numarasi;uniqueIndex" json:"seri_numarasi,omitempty"`
	SonGorulmeZamani         *time.Time `gorm:"column:son_gorulme_zamani" json:"son_gorulme_zamani,omitempty"`
	AktiflikBilgisi          bool       `gorm:"column:aktiflik_bilgisi;default:true" json:"aktiflik_bilgisi"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (TabletCihaz) TableName() string {
	return "tablet_cihaz"
}
