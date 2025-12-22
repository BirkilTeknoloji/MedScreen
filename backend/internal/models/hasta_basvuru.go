package models

import "time"

// HastaBasvuru represents a patient visit/admission in the VEM 2.0 schema (replaces Appointment)
type HastaBasvuru struct {
	HastaBasvuruKodu         string     `gorm:"column:hasta_basvuru_kodu;primaryKey" json:"hasta_basvuru_kodu"`
	HastaKodu                string     `gorm:"column:hasta_kodu;not null" json:"hasta_kodu"`
	Hasta                    *Hasta     `gorm:"foreignKey:HastaKodu;references:HastaKodu" json:"hasta,omitempty"`
	BasvuruProtokolNumarasi  string     `gorm:"column:basvuru_protokol_numarasi;uniqueIndex;not null" json:"basvuru_protokol_numarasi"`
	HastaKabulZamani         time.Time  `gorm:"column:hasta_kabul_zamani;not null" json:"hasta_kabul_zamani"`
	CikisZamani              *time.Time `gorm:"column:cikis_zamani" json:"cikis_zamani,omitempty"`
	HekimKodu                *string    `gorm:"column:hekim_kodu" json:"hekim_kodu,omitempty"`
	Hekim                    *Personel  `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	BasvuruDurumu            *string    `gorm:"column:basvuru_durumu" json:"basvuru_durumu,omitempty"`
	HayatiTehlikeDurumu      *string    `gorm:"column:hayati_tehlike_durumu" json:"hayati_tehlike_durumu,omitempty"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (HastaBasvuru) TableName() string {
	return "hasta_basvuru"
}
