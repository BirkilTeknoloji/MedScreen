package models

import "time"

// NFCKart represents an NFC card in the VEM 2.0 schema (replaces NFCCard)
type NFCKart struct {
	NFCKartKodu              string     `gorm:"column:nfc_kart_kodu;primaryKey" json:"nfc_kart_kodu"`
	PersonelKodu             string     `gorm:"column:personel_kodu;not null" json:"personel_kodu"`
	Personel                 *Personel  `gorm:"foreignKey:PersonelKodu;references:PersonelKodu" json:"personel,omitempty"`
	KartUID                  string     `gorm:"column:kart_uid;uniqueIndex;not null" json:"kart_uid"`
	SonKullanimTarihi        *time.Time `gorm:"column:son_kullanim_tarihi" json:"son_kullanim_tarihi,omitempty"`
	AktiflikBilgisi          int        `gorm:"column:aktiflik_bilgisi;default:1" json:"aktiflik_bilgisi"`
	KayitZamani              time.Time  `gorm:"column:kayit_zamani;not null" json:"kayit_zamani"`
	EkleyenKullaniciKodu     string     `gorm:"column:ekleyen_kullanici_kodu;not null" json:"ekleyen_kullanici_kodu"`
	GuncellemeZamani         *time.Time `gorm:"column:guncelleme_zamani" json:"guncelleme_zamani,omitempty"`
	GuncelleyenKullaniciKodu *string    `gorm:"column:guncelleyen_kullanici_kodu" json:"guncelleyen_kullanici_kodu,omitempty"`
}

// TableName returns the VEM 2.0 table name
func (NFCKart) TableName() string {
	return "nfc_kart"
}
