package models

import "time"

// BasvuruTani represents a diagnosis in the VEM 2.0 schema (replaces Diagnosis)
type BasvuruTani struct {
	BasvuruTaniKodu          string        `gorm:"column:basvuru_tani_kodu;primaryKey" json:"basvuru_tani_kodu"`
	HastaKodu                string        `gorm:"column:hasta_kodu;not null" json:"hasta_kodu"`
	Hasta                    *Hasta        `gorm:"foreignKey:HastaKodu;references:HastaKodu" json:"hasta,omitempty"`
	HastaBasvuruKodu         string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	TaniKodu                 string        `gorm:"column:tani_kodu;not null" json:"tani_kodu"`
	TaniTuru                 *string       `gorm:"column:tani_turu" json:"tani_turu,omitempty"`
	BirincilTani             int           `gorm:"column:birincil_tani;default:0" json:"birincil_tani"`
	TaniZamani               time.Time     `gorm:"column:tani_zamani;not null" json:"tani_zamani"`
	HekimKodu                *string       `gorm:"column:hekim_kodu" json:"hekim_kodu,omitempty"`
	Hekim                    *Personel     `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (BasvuruTani) TableName() string {
	return "basvuru_tani"
}
