package models

import "time"

// TibbiOrder represents a medical order in the VEM 2.0 schema (new entity)
type TibbiOrder struct {
	TibbiOrderKodu           string            `gorm:"column:tibbi_order_kodu;primaryKey" json:"tibbi_order_kodu"`
	HastaBasvuruKodu         string            `gorm:"column:hasta_basvuru_kodu;not null" json:"hasta_basvuru_kodu"`
	HastaBasvuru             *HastaBasvuru     `gorm:"foreignKey:HastaBasvuruKodu;references:HastaBasvuruKodu" json:"hasta_basvuru,omitempty"`
	OrderTuruKodu            string            `gorm:"column:order_turu_kodu;not null" json:"order_turu_kodu"`
	Aciklama                 *string           `gorm:"column:aciklama;type:text" json:"aciklama,omitempty"`
	OrderZamani              time.Time         `gorm:"column:order_zamani;not null" json:"order_zamani"`
	HekimKodu                string            `gorm:"column:hekim_kodu;not null" json:"hekim_kodu"`
	Hekim                    *Personel         `gorm:"foreignKey:HekimKodu;references:PersonelKodu" json:"hekim,omitempty"`
	IptalDurumu              int               `gorm:"column:iptal_durumu;default:0" json:"iptal_durumu"`
	KayitZamani              time.Time         `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string            `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time        `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string           `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
	Detaylar                 []TibbiOrderDetay `gorm:"foreignKey:TibbiOrderKodu;references:TibbiOrderKodu" json:"detaylar,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (TibbiOrder) TableName() string {
	return "tibbi_order"
}

// TibbiOrderDetay represents medical order execution details in the VEM 2.0 schema (new entity)
type TibbiOrderDetay struct {
	TibbiOrderDetayKodu     string      `gorm:"column:tibbi_order_detay_kodu;primaryKey" json:"tibbi_order_detay_kodu"`
	TibbiOrderKodu          string      `gorm:"column:tibbi_order_kodu;not null" json:"tibbi_order_kodu"`
	TibbiOrder              *TibbiOrder `gorm:"foreignKey:TibbiOrderKodu;references:TibbiOrderKodu" json:"tibbi_order,omitempty"`
	PlanlananUygulamaZamani time.Time   `gorm:"column:planlanan_uygulama_zamani;not null" json:"planlanan_uygulama_zamani"`
	UygulamaZamani          *time.Time  `gorm:"column:uygulama_zamani" json:"uygulama_zamani,omitempty"`
	UygulanmaDurumu         int         `gorm:"column:uygulanma_durumu;default:0" json:"uygulanma_durumu"`
	UygulayanPersonelKodu   *string     `gorm:"column:uygulayan_personel_kodu" json:"uygulayan_personel_kodu,omitempty"`
	UygulayanPersonel       *Personel   `gorm:"foreignKey:UygulayanPersonelKodu;references:PersonelKodu" json:"uygulayan_personel,omitempty"`
	KayitZamani             time.Time   `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu    string      `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
}

// TableName returns the VEM 2.0 table name
func (TibbiOrderDetay) TableName() string {
	return "tibbi_order_detay"
}
