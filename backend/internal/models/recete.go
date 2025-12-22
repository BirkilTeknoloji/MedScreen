package models

import "time"

// Recete represents a prescription in the VEM 2.0 schema (replaces Prescription)
type Recete struct {
	ReceteKodu               string        `gorm:"column:recete_kodu;primaryKey" json:"recete_kodu"`
	HastaBasvuruKodu         string        `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	MedulaEReceteNumarasi    *string       `gorm:"column:medula_e_recete_numarasi" json:"medula_e_recete_numarasi,omitempty"`
	ReceteTuruKodu           string        `gorm:"column:recete_turu_kodu;not null" json:"recete_turu_kodu"`
	HekimKodu                string        `gorm:"column:hekim_kodu;not null" json:"hekim_kodu"`
	Hekim                    *Personel     `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	ReceteZamani             time.Time     `gorm:"column:recete_zamani;not null" json:"recete_zamani"`
	KayitZamani              time.Time     `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string        `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time    `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string       `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
	AktiflikBilgisi          int           `gorm:"column:aktiflik_bilgisi;default:1" json:"aktiflik_bilgisi"`
	Ilaclar                  []ReceteIlac  `gorm:"foreignKey:ReceteKodu;references:ReceteKodu" json:"ilaclar,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (Recete) TableName() string {
	return "recete"
}

// ReceteIlac represents prescription medications in the VEM 2.0 schema
type ReceteIlac struct {
	ReceteIlacKodu             string     `gorm:"column:recete_ilac_kodu;primaryKey" json:"recete_ilac_kodu"`
	ReceteKodu                 string     `gorm:"column:recete_kodu;not null" json:"recete_kodu"`
	Recete                     *Recete    `gorm:"foreignKey:ReceteKodu;references:ReceteKodu" json:"recete,omitempty"`
	Barkod                     string     `gorm:"column:barkod;not null" json:"barkod"`
	IlacAdi                    *string    `gorm:"column:ilac_adi" json:"ilac_adi,omitempty"`
	IlacKullanimDozu           string     `gorm:"column:ilac_kullanim_dozu;not null" json:"ilac_kullanim_dozu"`
	DozBirim                   string     `gorm:"column:doz_birim;not null" json:"doz_birim"`
	IlacKullanimPeriyodu       *int       `gorm:"column:ilac_kullanim_periyodu" json:"ilac_kullanim_periyodu,omitempty"`
	IlacKullanimPeriyoduBirimi string     `gorm:"column:ilac_kullanim_periyodu_birimi;not null" json:"ilac_kullanim_periyodu_birimi"`
	IlacKullanimSekli          *string    `gorm:"column:ilac_kullanim_sekli" json:"ilac_kullanim_sekli,omitempty"`
	KutuAdeti                  *int       `gorm:"column:kutu_adeti" json:"kutu_adeti,omitempty"`
	KayitZamani                time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu       string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani           *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu   *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (ReceteIlac) TableName() string {
	return "recete_ilac"
}
