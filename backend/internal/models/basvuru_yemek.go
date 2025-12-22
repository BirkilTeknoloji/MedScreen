package models

import "time"

// BasvuruYemek represents patient diet/meal information in the VEM 2.0 schema (new entity)
type BasvuruYemek struct {
	BasvuruYemekKodu         string        `gorm:"column:basvuru_yemek_kodu;primaryKey" json:"basvuru_yemek_kodu"`
	HastaBasvuruKodu         string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	YemekZamaniTuru          string        `gorm:"column:yemek_zamani_turu;not null" json:"yemek_zamani_turu"`
	YemekTuru                string        `gorm:"column:yemek_turu;not null" json:"yemek_turu"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (BasvuruYemek) TableName() string {
	return "basvuru_yemek"
}
