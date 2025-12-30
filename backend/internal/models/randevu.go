package models

import "time"

// Randevu represents appointment information in the VEM 2.0 schema
type Randevu struct {
	RandevuKodu              string        `gorm:"column:randevu_kodu;primaryKey" json:"randevu_kodu"`
	HastaKodu                string        `gorm:"column:hasta_kodu;not null" json:"hasta_kodu"`
	Hasta                    *Hasta        `gorm:"foreignKey:HastaKodu;references:HastaKodu" json:"hasta,omitempty"`
	HastaBasvuruKodu         *string       `gorm:"column:hasta_basvuru_kodu" json:"hasta_basvuru_kodu,omitempty"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	HekimKodu                *string       `gorm:"column:hekim_kodu" json:"hekim_kodu,omitempty"`
	Hekim                    *Personel     `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	BirimKodu                *string       `gorm:"column:birim_kodu" json:"birim_kodu,omitempty"`
	RandevuTuru              string        `gorm:"column:randevu_turu;not null" json:"randevu_turu"`
	RandevuZamani            time.Time     `gorm:"column:randevu_zamani;not null" json:"randevu_zamani"`
	RandevuGelmeDurumu       *string       `gorm:"column:randevu_gelme_durumu" json:"randevu_gelme_durumu,omitempty"`
	Aciklama                 *string       `gorm:"column:aciklama;type:text" json:"aciklama,omitempty"`
	IptalDurumu              int           `gorm:"column:iptal_durumu;default:0" json:"iptal_durumu"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (Randevu) TableName() string {
	return "randevu"
}
