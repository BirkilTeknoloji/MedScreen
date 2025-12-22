package models

import "time"

// KlinikSeyir represents clinical progress notes in the VEM 2.0 schema (new entity)
type KlinikSeyir struct {
	KlinikSeyirKodu          string        `gorm:"column:klinik_seyir_kodu;primaryKey" json:"klinik_seyir_kodu"`
	HastaBasvuruKodu         string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	SeyirTipi                string        `gorm:"column:seyir_tipi;not null" json:"seyir_tipi"`
	SeyirZamani              time.Time     `gorm:"column:seyir_zamani;not null;index" json:"seyir_zamani"`
	SeyirBilgisi             string        `gorm:"column:seyir_bilgisi;type:text;not null" json:"seyir_bilgisi"`
	SeptikSok                int           `gorm:"column:septik_sok;default:0" json:"septik_sok"`
	SepsisDurumu             int           `gorm:"column:sepsis_durumu;default:0" json:"sepsis_durumu"`
	HekimKodu                *string       `gorm:"column:hekim_kodu" json:"hekim_kodu,omitempty"`
	Hekim                    *Personel     `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (KlinikSeyir) TableName() string {
	return "klinik_seyir"
}
