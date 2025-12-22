package models

import "time"

// Personel represents a staff member in the VEM 2.0 schema (replaces User)
type Personel struct {
	PersonelKodu             string     `gorm:"column:personel_kodu;primaryKey" json:"personel_kodu"`
	Ad                       string     `gorm:"column:ad;not null" json:"ad"`
	Soyadi                   string     `gorm:"column:soyadi;not null" json:"soyadi"`
	PersonelGorevKodu        string     `gorm:"column:personel_gorev_kodu;not null" json:"personel_gorev_kodu"`
	MedulaBransKodu          *string    `gorm:"column:medula_brans_kodu" json:"medula_brans_kodu,omitempty"`
	TescilNumarasi           *string    `gorm:"column:tescil_numarasi" json:"tescil_numarasi,omitempty"`
	TCKimlikNumarasi         *string    `gorm:"column:tc_kimlik_numarasi" json:"tc_kimlik_numarasi,omitempty"`
	AktiflikBilgisi          int        `gorm:"column:aktiflik_bilgisi;default:1" json:"aktiflik_bilgisi"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (Personel) TableName() string {
	return "personel"
}
