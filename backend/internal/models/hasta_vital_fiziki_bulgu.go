package models

import "time"

// HastaVitalFizikiBulgu represents patient vital signs in the VEM 2.0 schema (replaces VitalSign)
type HastaVitalFizikiBulgu struct {
	HastaVitalFizikiBulguKodu string        `gorm:"column:hasta_vital_fiziki_bulgu_kodu;primaryKey" json:"hasta_vital_fiziki_bulgu_kodu"`
	HastaBasvuruKodu          string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru              *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	IslemZamani               time.Time     `gorm:"column:islem_zamani;not null;index" json:"islem_zamani"`
	Ates                      *string       `gorm:"column:ates" json:"ates,omitempty"`
	Nabiz                     *string       `gorm:"column:nabiz" json:"nabiz,omitempty"`
	SistolikKanBasinciDegeri  *string       `gorm:"column:sistolik_kan_basinci_degeri" json:"sistolik_kan_basinci_degeri,omitempty"`
	DiastolikKanBasinciDegeri *string       `gorm:"column:diastolik_kan_basinci_degeri" json:"diastolik_kan_basinci_degeri,omitempty"`
	Solunum                   *string       `gorm:"column:solunum" json:"solunum,omitempty"`
	Saturasyon                *string       `gorm:"column:saturasyon" json:"saturasyon,omitempty"`
	Boy                       *string       `gorm:"column:boy" json:"boy,omitempty"`
	Agirlik                   *string       `gorm:"column:agirlik" json:"agirlik,omitempty"`
	HemsireKodu               *string       `gorm:"column:hemsire_kodu" json:"hemsire_kodu,omitempty"`
	Hemsire                   *Personel     `gorm:"foreignKey:HemsireKodu;references:PersonelKodu" json:"hemsire,omitempty"`
	KayitZamani               time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu      string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani          *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu  *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (HastaVitalFizikiBulgu) TableName() string {
	return "hasta_vital_fiziki_bulgu"
}
