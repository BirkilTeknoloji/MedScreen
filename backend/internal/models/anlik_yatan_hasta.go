package models

import "time"

// AnlikYatanHasta represents a current inpatient in the VEM 2.0 schema (new entity)
type AnlikYatanHasta struct {
	AnlikYatanHastaKodu  string        `gorm:"column:anlik_yatan_hasta_kodu;primaryKey" json:"anlik_yatan_hasta_kodu"`
	HastaBasvuruKodu     string        `gorm:"column:hasta_basvuru_kodu;not null;index" json:"hasta_basvuru_kodu"`
	HastaBasvuru         *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	HastaKodu            string        `gorm:"column:hasta_kodu;not null" json:"hasta_kodu"`
	Hasta                *Hasta        `gorm:"foreignKey:HastaKodu;references:HastaKodu" json:"hasta,omitempty"`
	YatakKodu            string        `gorm:"column:yatak_kodu;not null;index" json:"yatak_kodu"`
	Yatak                *Yatak        `gorm:"foreignKey:YatakKodu;references:YatakKodu" json:"yatak,omitempty"`
	BirimKodu            *string       `gorm:"column:birim_kodu" json:"birim_kodu,omitempty"`
	YatisZamani          time.Time     `gorm:"column:yatis_zamani;not null" json:"yatis_zamani"`
	HekimKodu            *string       `gorm:"column:hekim_kodu" json:"hekim_kodu,omitempty"`
	Hekim                *Personel     `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	KayitZamani          time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
}

// TableName returns the VEM 2.0 table name
func (AnlikYatanHasta) TableName() string {
	return "anlik_yatan_hasta"
}
