package models

import "time"

// Hasta represents a patient in the VEM 2.0 schema (replaces Patient)
type Hasta struct {
	HastaKodu                string     `gorm:"column:hasta_kodu;primaryKey" json:"hasta_kodu"`
	TCKimlikNumarasi         *string    `gorm:"column:tc_kimlik_numarasi;uniqueIndex" json:"tc_kimlik_numarasi,omitempty"`
	Ad                       string     `gorm:"column:ad;not null" json:"ad"`
	Soyadi                   string     `gorm:"column:soyadi;not null" json:"soyadi"`
	AnneHastaKodu            *string    `gorm:"column:anne_hasta_kodu" json:"anne_hasta_kodu,omitempty"`
	Anne                     *Hasta     `gorm:"foreignKey:AnneHastaKodu;references:HastaKodu" json:"anne,omitempty"`
	BabaHastaKodu            *string    `gorm:"column:baba_hasta_kodu" json:"baba_hasta_kodu,omitempty"`
	Baba                     *Hasta     `gorm:"foreignKey:BabaHastaKodu;references:HastaKodu" json:"baba,omitempty"`
	DogumTarihi              time.Time  `gorm:"column:dogum_tarihi;not null" json:"dogum_tarihi"`
	Cinsiyet                 *string    `gorm:"column:cinsiyet" json:"cinsiyet,omitempty"`
	KanGrubu                 *string    `gorm:"column:kan_grubu" json:"kan_grubu,omitempty"`
	Uyruk                    *string    `gorm:"column:uyruk" json:"uyruk,omitempty"`
	HastaTipi                *string    `gorm:"column:hasta_tipi" json:"hasta_tipi,omitempty"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (Hasta) TableName() string {
	return "hasta"
}
