package models

import "time"

// TetkikSonuc represents test results in the VEM 2.0 schema (replaces MedicalTest)
type TetkikSonuc struct {
	TetkikSonucKodu      string        `gorm:"column:tetkik_sonuc_kodu;primaryKey" json:"tetkik_sonuc_kodu"`
	HastaBasvuruKodu     string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru         *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	TetkikAdi            string        `gorm:"column:tetkik_adi;not null" json:"tetkik_adi"`
	SonucDegeri          *string       `gorm:"column:sonuc_degeri" json:"sonuc_degeri,omitempty"`
	KritikDegerAraligi   *string       `gorm:"column:kritik_deger_araligi" json:"kritik_deger_araligi,omitempty"`
	OnayZamani           *time.Time    `gorm:"column:onay_zamani" json:"onay_zamani,omitempty"`
	KayitZamani          time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
}

// TableName returns the VEM 2.0 table name
func (TetkikSonuc) TableName() string {
	return "tetkik_sonuc"
}
