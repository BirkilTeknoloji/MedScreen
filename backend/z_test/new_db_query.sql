-- MEDSCREEN_YATAK_BASI_TABLET_FINAL
-- PostgreSQL DDL + Seed Data (Backend test için tutarlı örnekler)

BEGIN;

SET TIME ZONE 'Europe/Istanbul';

-- ---------------------------------------------------------
-- DROP (FK bağımlılık sırasına göre)
-- ---------------------------------------------------------
DROP TABLE IF EXISTS basvuru_tani CASCADE;
DROP TABLE IF EXISTS hasta_uyari CASCADE;
DROP TABLE IF EXISTS risk_skorlama CASCADE;
DROP TABLE IF EXISTS basvuru_yemek CASCADE;

DROP TABLE IF EXISTS recete_ilac CASCADE;
DROP TABLE IF EXISTS recete CASCADE;

DROP TABLE IF EXISTS tetkik_sonuc CASCADE;
DROP TABLE IF EXISTS tibbi_order_detay CASCADE;
DROP TABLE IF EXISTS tibbi_order CASCADE;
DROP TABLE IF EXISTS klinik_seyir CASCADE;
DROP TABLE IF EXISTS hasta_vital_fiziki_bulgu CASCADE;

DROP TABLE IF EXISTS anlik_yatan_hasta CASCADE;
DROP TABLE IF EXISTS tablet_cihaz CASCADE;
DROP TABLE IF EXISTS yatak CASCADE;

DROP TABLE IF EXISTS hasta_tibbi_bilgi CASCADE;
DROP TABLE IF EXISTS hasta_basvuru CASCADE;
DROP TABLE IF EXISTS hasta CASCADE;

DROP TABLE IF EXISTS nfc_kart CASCADE;
DROP TABLE IF EXISTS personel CASCADE;

-- ---------------------------------------------------------
-- 1. PERSONEL VE GÜVENLİK
-- ---------------------------------------------------------
CREATE TABLE personel (
  personel_kodu              varchar PRIMARY KEY,
  ad                         varchar NOT NULL,
  soyadi                     varchar NOT NULL,
  personel_gorev_kodu        varchar NOT NULL,
  medula_brans_kodu          varchar,
  tescil_numarasi            varchar,
  tc_kimlik_numarasi         varchar,
  aktiflik_bilgisi           int NOT NULL DEFAULT 1,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE INDEX idx_personel_ad_soyadi ON personel (ad, soyadi);

CREATE TABLE nfc_kart (
  nfc_kart_kodu              varchar PRIMARY KEY,
  personel_kodu              varchar NOT NULL REFERENCES personel(personel_kodu),
  kart_uid                   varchar NOT NULL UNIQUE,
  son_kullanim_tarihi        timestamp,
  aktiflik_bilgisi           int NOT NULL DEFAULT 1,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

-- ---------------------------------------------------------
-- 2. HASTA VE KABUL İŞLEMLERİ
-- ---------------------------------------------------------
CREATE TABLE hasta (
  hasta_kodu                 varchar PRIMARY KEY,
  tc_kimlik_numarasi         varchar UNIQUE,
  ad                         varchar NOT NULL,
  soyadi                     varchar NOT NULL,
  anne_hasta_kodu            varchar REFERENCES hasta(hasta_kodu),
  baba_hasta_kodu            varchar REFERENCES hasta(hasta_kodu),
  dogum_tarihi               date NOT NULL,
  cinsiyet                   varchar,
  kan_grubu                  varchar,
  uyruk                      varchar,
  hasta_tipi                 varchar,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE INDEX idx_hasta_tc ON hasta (tc_kimlik_numarasi);
CREATE INDEX idx_hasta_ad_soyadi ON hasta (ad, soyadi);

CREATE TABLE hasta_basvuru (
  hasta_basvuru_kodu         varchar PRIMARY KEY,
  hasta_kodu                 varchar NOT NULL REFERENCES hasta(hasta_kodu),
  basvuru_protokol_numarasi  varchar NOT NULL UNIQUE,
  hasta_kabul_zamani         timestamp NOT NULL,
  cikis_zamani               timestamp,
  hekim_kodu                 varchar REFERENCES personel(personel_kodu),
  basvuru_durumu             varchar,
  hayati_tehlike_durumu      varchar,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE TABLE hasta_tibbi_bilgi (
  hasta_tibbi_bilgi_kodu     varchar PRIMARY KEY,
  hasta_kodu                 varchar NOT NULL REFERENCES hasta(hasta_kodu),
  tibbi_bilgi_turu_kodu      varchar NOT NULL,
  tibbi_bilgi_alt_turu_kodu  varchar,
  aciklama                   text,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

-- ---------------------------------------------------------
-- 3. KONUM VE TABLET DONANIMI
-- ---------------------------------------------------------
CREATE TABLE yatak (
  yatak_kodu                 varchar PRIMARY KEY,
  birim_kodu                 varchar NOT NULL,
  oda_kodu                   varchar NOT NULL,
  yatak_adi                  varchar,
  yatak_turu_kodu            varchar,
  yogun_bakim_yatak_seviyesi varchar,
  ventilator_cihaz_kodu      varchar,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE TABLE tablet_cihaz (
  tablet_cihaz_kodu          varchar PRIMARY KEY,
  yatak_kodu                 varchar REFERENCES yatak(yatak_kodu),
  ip_adresi                  varchar,
  seri_numarasi              varchar UNIQUE,
  son_gorulme_zamani         timestamp,
  aktiflik_bilgisi           boolean NOT NULL DEFAULT true,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE INDEX idx_tablet_yatak_kodu ON tablet_cihaz (yatak_kodu);

CREATE TABLE anlik_yatan_hasta (
  anlik_yatan_hasta_kodu     varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  hasta_kodu                 varchar NOT NULL REFERENCES hasta(hasta_kodu),
  yatak_kodu                 varchar NOT NULL REFERENCES yatak(yatak_kodu),
  birim_kodu                 varchar,
  yatis_zamani               timestamp NOT NULL,
  hekim_kodu                 varchar REFERENCES personel(personel_kodu),
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL
);

CREATE INDEX idx_anlik_yatan_hasta_yatak ON anlik_yatan_hasta (yatak_kodu);
CREATE INDEX idx_anlik_yatan_hasta_basvuru ON anlik_yatan_hasta (hasta_basvuru_kodu);

-- ---------------------------------------------------------
-- 4. KLİNİK VERİLER (VİTAL, SEYİR VE TETKİK)
-- ---------------------------------------------------------
CREATE TABLE hasta_vital_fiziki_bulgu (
  hasta_vital_fiziki_bulgu_kodu varchar PRIMARY KEY,
  hasta_basvuru_kodu            varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  islem_zamani                  timestamp NOT NULL,
  ates                          varchar,
  nabiz                         varchar,
  sistolik_kan_basinci_degeri   varchar,
  diastolik_kan_basinci_degeri  varchar,
  solunum                       varchar,
  saturasyon                    varchar,
  boy                           varchar,
  agirlik                       varchar,
  hemsire_kodu                  varchar REFERENCES personel(personel_kodu),
  kayit_zamani                  timestamp NOT NULL,
  ekleyen_kullanici_kodu        varchar NOT NULL,
  guncelleme_zamani             timestamp,
  guncelleyen_kullanici_kodu    varchar
);

CREATE INDEX idx_vital_islem_zamani ON hasta_vital_fiziki_bulgu (islem_zamani);

CREATE TABLE klinik_seyir (
  klinik_seyir_kodu          varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  seyir_tipi                 varchar NOT NULL,
  seyir_zamani               timestamp NOT NULL,
  seyir_bilgisi              text NOT NULL,
  septik_sok                 int NOT NULL DEFAULT 0,
  sepsis_durumu              int NOT NULL DEFAULT 0,
  hekim_kodu                 varchar REFERENCES personel(personel_kodu),
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE INDEX idx_klinik_seyir_zamani ON klinik_seyir (seyir_zamani);

CREATE TABLE tibbi_order (
  tibbi_order_kodu           varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  order_turu_kodu            varchar NOT NULL,
  aciklama                   text,
  order_zamani               timestamp NOT NULL,
  hekim_kodu                 varchar NOT NULL REFERENCES personel(personel_kodu),
  iptal_durumu               int NOT NULL DEFAULT 0,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE TABLE tibbi_order_detay (
  tibbi_order_detay_kodu     varchar PRIMARY KEY,
  tibbi_order_kodu           varchar NOT NULL REFERENCES tibbi_order(tibbi_order_kodu),
  planlanan_uygulama_zamani  timestamp NOT NULL,
  uygulama_zamani            timestamp,
  uygulanma_durumu           int NOT NULL DEFAULT 0,
  uygulayan_personel_kodu    varchar REFERENCES personel(personel_kodu),
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL
);

CREATE TABLE tetkik_sonuc (
  tetkik_sonuc_kodu          varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  tetkik_adi                 varchar NOT NULL,
  sonuc_degeri               varchar,
  kritik_deger_araligi       varchar,
  onay_zamani                timestamp,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL
);

-- ---------------------------------------------------------
-- 5. REÇETE VE İLAÇ YÖNETİMİ
-- ---------------------------------------------------------
CREATE TABLE recete (
  recete_kodu                varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  medula_e_recete_numarasi   varchar,
  recete_turu_kodu           varchar NOT NULL,
  hekim_kodu                 varchar NOT NULL REFERENCES personel(personel_kodu),
  recete_zamani              timestamp NOT NULL,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar,
  aktiflik_bilgisi           int NOT NULL DEFAULT 1
);

CREATE TABLE recete_ilac (
  recete_ilac_kodu           varchar PRIMARY KEY,
  recete_kodu                varchar NOT NULL REFERENCES recete(recete_kodu),
  barkod                     varchar NOT NULL,
  ilac_adi                   varchar,
  ilac_kullanim_dozu         varchar NOT NULL,
  doz_birim                  varchar NOT NULL,
  ilac_kullanim_periyodu     int,
  ilac_kullanim_periyodu_birimi varchar NOT NULL,
  ilac_kullanim_sekli        varchar,
  kutu_adeti                 int,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

-- ---------------------------------------------------------
-- 6. DİYET, RİSK VE UYARILAR
-- ---------------------------------------------------------
CREATE TABLE basvuru_yemek (
  basvuru_yemek_kodu         varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  yemek_zamani_turu          varchar NOT NULL,
  yemek_turu                 varchar NOT NULL,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE TABLE risk_skorlama (
  risk_skorlama_kodu         varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  risk_skorlama_turu         varchar NOT NULL,
  risk_skorlama_toplam_puani varchar NOT NULL,
  islem_zamani               timestamp NOT NULL,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE TABLE hasta_uyari (
  hasta_uyari_kodu           varchar PRIMARY KEY,
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  uyari_turu                 varchar NOT NULL,
  uyari_aciklama             text,
  aktiflik_bilgisi           int NOT NULL DEFAULT 1,
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

CREATE TABLE basvuru_tani (
  basvuru_tani_kodu          varchar PRIMARY KEY,
  hasta_kodu                 varchar NOT NULL REFERENCES hasta(hasta_kodu),
  hasta_basvuru_kodu         varchar NOT NULL REFERENCES hasta_basvuru(hasta_basvuru_kodu),
  tani_kodu                  varchar NOT NULL,
  tani_turu                  varchar,
  birincil_tani              int NOT NULL DEFAULT 0,
  tani_zamani                timestamp NOT NULL,
  hekim_kodu                 varchar REFERENCES personel(personel_kodu),
  kayit_zamani               timestamp NOT NULL,
  ekleyen_kullanici_kodu     varchar NOT NULL,
  guncelleme_zamani          timestamp,
  guncelleyen_kullanici_kodu varchar
);

-- ---------------------------------------------------------
-- SEED DATA (her tablo 3 kayıt)
-- ---------------------------------------------------------

-- PERSONEL (3-4 kişi ile test senaryosu genişler, burada 4 verdim)
INSERT INTO personel (
  personel_kodu, ad, soyadi, personel_gorev_kodu, medula_brans_kodu, tescil_numarasi, tc_kimlik_numarasi,
  aktiflik_bilgisi, kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('P0001','Ahmet','Yilmaz','HEKIM','1300','TSCL-1001','11111111111',1,'2025-12-20 09:00:00','SYS',NULL,NULL),
  ('P0002','Elif','Kaya','HEMSIRE','9999','TSCL-2001','22222222222',1,'2025-12-20 09:05:00','SYS',NULL,NULL),
  ('P0003','Mehmet','Demir','HEKIM','4500','TSCL-1002','33333333333',1,'2025-12-20 09:10:00','SYS',NULL,NULL),
  ('P0004','Ayse','Celik','ADMIN',NULL,NULL,'44444444444',1,'2025-12-20 09:15:00','SYS',NULL,NULL);

-- NFC_KART (3 kayıt)
INSERT INTO nfc_kart (
  nfc_kart_kodu, personel_kodu, kart_uid, son_kullanim_tarihi, aktiflik_bilgisi,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('NFK0001','P0001','04AABBCCDD01','2025-12-21 08:30:00',1,'2025-12-20 10:00:00','P0004',NULL,NULL),
  ('NFK0002','P0002','04AABBCCDD02','2025-12-21 08:45:00',1,'2025-12-20 10:05:00','P0004',NULL,NULL),
  ('NFK0003','P0003','04AABBCCDD03','2025-12-21 09:00:00',1,'2025-12-20 10:10:00','P0004',NULL,NULL);

-- HASTA (3 kayıt, biri anne-baba referanslı çocuk)
INSERT INTO hasta (
  hasta_kodu, tc_kimlik_numarasi, ad, soyadi, anne_hasta_kodu, baba_hasta_kodu,
  dogum_tarihi, cinsiyet, kan_grubu, uyruk, hasta_tipi,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('H0001','55555555555','Fatma','Acar',NULL,NULL,'1984-05-12','K','A+','TR','Yatan','2025-12-20 11:00:00','SYS',NULL,NULL),
  ('H0002','66666666666','Ali','Acar',NULL,NULL,'1980-02-20','E','0+','TR','Ayaktan','2025-12-20 11:05:00','SYS',NULL,NULL),
  ('H0003','77777777777','Zeynep','Acar','H0001','H0002','2012-09-01','K','A+','TR','Yatan','2025-12-20 11:10:00','SYS',NULL,NULL);

-- HASTA_BASVURU (3 kayıt)
INSERT INTO hasta_basvuru (
  hasta_basvuru_kodu, hasta_kodu, basvuru_protokol_numarasi, hasta_kabul_zamani, cikis_zamani,
  hekim_kodu, basvuru_durumu, hayati_tehlike_durumu,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('B0001','H0001','PRT-2025-0001','2025-12-21 07:40:00',NULL,'P0001','YATIS','HAYIR','2025-12-21 07:41:00','SYS',NULL,NULL),
  ('B0002','H0003','PRT-2025-0002','2025-12-21 08:10:00',NULL,'P0003','YOGUN_BAKIM','EVET','2025-12-21 08:11:00','SYS',NULL,NULL),
  ('B0003','H0002','PRT-2025-0003','2025-12-21 09:20:00','2025-12-21 12:00:00','P0001','TABURCU','HAYIR','2025-12-21 09:21:00','SYS',NULL,NULL);

-- HASTA_TIBBI_BILGI (3 kayıt)
INSERT INTO hasta_tibbi_bilgi (
  hasta_tibbi_bilgi_kodu, hasta_kodu, tibbi_bilgi_turu_kodu, tibbi_bilgi_alt_turu_kodu, aciklama,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('TB0001','H0001','ALERJI','PENISILIN','Penisilin alerjisi bildirildi.','2025-12-21 08:00:00','P0001',NULL,NULL),
  ('TB0002','H0003','KRONIK','ASTIM','Çocukluk çağı astım öyküsü.','2025-12-21 08:20:00','P0003',NULL,NULL),
  ('TB0003','H0002','GECMIS','HT','Hipertansiyon tanısı mevcut.','2025-12-21 09:30:00','P0001',NULL,NULL);

-- YATAK (3 kayıt)
INSERT INTO yatak (
  yatak_kodu, birim_kodu, oda_kodu, yatak_adi, yatak_turu_kodu, yogun_bakim_yatak_seviyesi, ventilator_cihaz_kodu,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('Y001','ICU','ICU-01','Yatak-1','YOGUN_BAKIM','3','VENT-01','2025-12-20 12:00:00','P0004',NULL,NULL),
  ('Y002','SERVIS','SRV-12','Yatak-2','SERVIS',NULL,NULL,'2025-12-20 12:05:00','P0004',NULL,NULL),
  ('Y003','ICU','ICU-02','Yatak-3','YOGUN_BAKIM','2','VENT-02','2025-12-20 12:10:00','P0004',NULL,NULL);

-- TABLET_CIHAZ (3 kayıt)
INSERT INTO tablet_cihaz (
  tablet_cihaz_kodu, yatak_kodu, ip_adresi, seri_numarasi, son_gorulme_zamani, aktiflik_bilgisi,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('TBL001','Y001','10.10.1.11','SN-ICU-0001','2025-12-22 10:45:00',true,'2025-12-20 12:30:00','P0004',NULL,NULL),
  ('TBL002','Y002','10.10.1.12','SN-SRV-0002','2025-12-22 10:40:00',true,'2025-12-20 12:35:00','P0004',NULL,NULL),
  ('TBL003','Y003','10.10.1.13','SN-ICU-0003','2025-12-22 10:42:00',true,'2025-12-20 12:40:00','P0004',NULL,NULL);

-- ANLIK_YATAN_HASTA (3 kayıt)
INSERT INTO anlik_yatan_hasta (
  anlik_yatan_hasta_kodu, hasta_basvuru_kodu, hasta_kodu, yatak_kodu, birim_kodu, yatis_zamani, hekim_kodu,
  kayit_zamani, ekleyen_kullanici_kodu
) VALUES
  ('AYH001','B0001','H0001','Y002','SERVIS','2025-12-21 07:50:00','P0001','2025-12-21 07:51:00','SYS'),
  ('AYH002','B0002','H0003','Y001','ICU','2025-12-21 08:20:00','P0003','2025-12-21 08:21:00','SYS'),
  ('AYH003','B0003','H0002','Y003','ICU','2025-12-21 09:30:00','P0001','2025-12-21 09:31:00','SYS');

-- HASTA_VITAL_FIZIKI_BULGU (3 kayıt)
INSERT INTO hasta_vital_fiziki_bulgu (
  hasta_vital_fiziki_bulgu_kodu, hasta_basvuru_kodu, islem_zamani,
  ates, nabiz, sistolik_kan_basinci_degeri, diastolik_kan_basinci_degeri, solunum, saturasyon, boy, agirlik,
  hemsire_kodu, kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('VTL001','B0002','2025-12-21 08:30:00','38.2','112','110','70','24','92','145','42','P0002','2025-12-21 08:31:00','P0002',NULL,NULL),
  ('VTL002','B0002','2025-12-21 10:00:00','37.6','98','105','68','20','95','145','42','P0002','2025-12-21 10:01:00','P0002',NULL,NULL),
  ('VTL003','B0001','2025-12-21 09:10:00','36.9','84','125','80','18','98','165','70','P0002','2025-12-21 09:11:00','P0002',NULL,NULL);

-- KLINIK_SEYIR (3 kayıt)
INSERT INTO klinik_seyir (
  klinik_seyir_kodu, hasta_basvuru_kodu, seyir_tipi, seyir_zamani, seyir_bilgisi,
  septik_sok, sepsis_durumu, hekim_kodu,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('KS001','B0002','HEKIM_NOTU','2025-12-21 08:40:00','Ateş ve taşikardi mevcut, sepsis açısından yakın izlem.','0','1','P0003','2025-12-21 08:41:00','P0003',NULL,NULL),
  ('KS002','B0001','HEMSIRE_NOTU','2025-12-21 09:20:00','Hasta genel durumu iyi, mobilize ediliyor.','0','0',NULL,'2025-12-21 09:21:00','P0002',NULL,NULL),
  ('KS003','B0003','HEKIM_NOTU','2025-12-21 09:40:00','Ayaktan başvuru, tansiyon takibi önerildi.','0','0','P0001','2025-12-21 09:41:00','P0001',NULL,NULL);

-- TIBBI_ORDER (3 kayıt)
INSERT INTO tibbi_order (
  tibbi_order_kodu, hasta_basvuru_kodu, order_turu_kodu, aciklama, order_zamani, hekim_kodu,
  iptal_durumu, kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('ORD001','B0002','LAB','Hemogram + CRP isteniyor.','2025-12-21 08:45:00','P0003',0,'2025-12-21 08:46:00','P0003',NULL,NULL),
  ('ORD002','B0001','ILAC','Parasetamol 500mg gerektiğinde.','2025-12-21 09:25:00','P0001',0,'2025-12-21 09:26:00','P0001',NULL,NULL),
  ('ORD003','B0003','KONTROL','24 saat içinde kontrol önerisi.','2025-12-21 09:50:00','P0001',0,'2025-12-21 09:51:00','P0001',NULL,NULL);

-- TIBBI_ORDER_DETAY (3 kayıt)
INSERT INTO tibbi_order_detay (
  tibbi_order_detay_kodu, tibbi_order_kodu, planlanan_uygulama_zamani, uygulama_zamani,
  uygulanma_durumu, uygulayan_personel_kodu, kayit_zamani, ekleyen_kullanici_kodu
) VALUES
  ('ODT001','ORD001','2025-12-21 09:00:00','2025-12-21 09:05:00',1,'P0002','2025-12-21 09:06:00','P0002'),
  ('ODT002','ORD002','2025-12-21 10:00:00',NULL,0,NULL,'2025-12-21 09:27:00','P0001'),
  ('ODT003','ORD003','2025-12-22 09:00:00',NULL,0,NULL,'2025-12-21 09:52:00','P0001');

-- TETKIK_SONUC (3 kayıt)
INSERT INTO tetkik_sonuc (
  tetkik_sonuc_kodu, hasta_basvuru_kodu, tetkik_adi, sonuc_degeri, kritik_deger_araligi, onay_zamani,
  kayit_zamani, ekleyen_kullanici_kodu
) VALUES
  ('TS001','B0002','CRP','78 mg/L','0-5 mg/L','2025-12-21 11:10:00','2025-12-21 11:11:00','SYS'),
  ('TS002','B0002','WBC','17.2 x10^9/L','4-11 x10^9/L','2025-12-21 11:12:00','2025-12-21 11:13:00','SYS'),
  ('TS003','B0001','Glukoz','102 mg/dL','70-110 mg/dL','2025-12-21 10:20:00','2025-12-21 10:21:00','SYS');

-- RECETE (3 kayıt)
INSERT INTO recete (
  recete_kodu, hasta_basvuru_kodu, medula_e_recete_numarasi, recete_turu_kodu, hekim_kodu,
  recete_zamani, kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu, aktiflik_bilgisi
) VALUES
  ('RCP001','B0001','E-REC-10001','NORMAL','P0001','2025-12-21 09:30:00','2025-12-21 09:31:00','P0001',NULL,NULL,1),
  ('RCP002','B0002','E-REC-10002','NORMAL','P0003','2025-12-21 10:10:00','2025-12-21 10:11:00','P0003',NULL,NULL,1),
  ('RCP003','B0003','E-REC-10003','NORMAL','P0001','2025-12-21 11:30:00','2025-12-21 11:31:00','P0001',NULL,NULL,1);

-- RECETE_ILAC (3 kayıt)
INSERT INTO recete_ilac (
  recete_ilac_kodu, recete_kodu, barkod, ilac_adi, ilac_kullanim_dozu, doz_birim,
  ilac_kullanim_periyodu, ilac_kullanim_periyodu_birimi, ilac_kullanim_sekli, kutu_adeti,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('RCI001','RCP001','869000000001','Paracetamol','500','mg',8,'SAAT','PO',1,'2025-12-21 09:32:00','P0001',NULL,NULL),
  ('RCI002','RCP002','869000000002','Amoxicillin','250','mg',12,'SAAT','PO',1,'2025-12-21 10:12:00','P0003',NULL,NULL),
  ('RCI003','RCP003','869000000003','Amlodipin','5','mg',24,'SAAT','PO',1,'2025-12-21 11:32:00','P0001',NULL,NULL);

-- BASVURU_YEMEK (3 kayıt)
INSERT INTO basvuru_yemek (
  basvuru_yemek_kodu, hasta_basvuru_kodu, yemek_zamani_turu, yemek_turu,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('BY001','B0001','KAHVALTI','DIYET-STD','2025-12-21 08:05:00','SYS',NULL,NULL),
  ('BY002','B0002','OGLE','DIYET-PROTEIN','2025-12-21 08:06:00','SYS',NULL,NULL),
  ('BY003','B0003','AKSAM','DIYET-TUZSUZ','2025-12-21 08:07:00','SYS',NULL,NULL);

-- RISK_SKORLAMA (3 kayıt)
INSERT INTO risk_skorlama (
  risk_skorlama_kodu, hasta_basvuru_kodu, risk_skorlama_turu, risk_skorlama_toplam_puani,
  islem_zamani, kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('RSK001','B0002','NEWS2','7','2025-12-21 08:35:00','2025-12-21 08:36:00','P0002',NULL,NULL),
  ('RSK002','B0001','FALL_RISK','2','2025-12-21 09:15:00','2025-12-21 09:16:00','P0002',NULL,NULL),
  ('RSK003','B0003','FALL_RISK','1','2025-12-21 09:55:00','2025-12-21 09:56:00','P0002',NULL,NULL);

-- HASTA_UYARI (3 kayıt)
INSERT INTO hasta_uyari (
  hasta_uyari_kodu, hasta_basvuru_kodu, uyari_turu, uyari_aciklama, aktiflik_bilgisi,
  kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('UY001','B0001','ALERJI','Penisilin alerjisi var. Alternatif antibiyotik seçiniz.',1,'2025-12-21 08:02:00','P0001',NULL,NULL),
  ('UY002','B0002','KRITIK','Sepsis şüphesi. Vital bulgular sık izlenecek.',1,'2025-12-21 08:42:00','P0003',NULL,NULL),
  ('UY003','B0003','DIYET','Tuzsuz diyet önerisi.',1,'2025-12-21 09:57:00','P0001',NULL,NULL);

-- BASVURU_TANI (3 kayıt)
INSERT INTO basvuru_tani (
  basvuru_tani_kodu, hasta_kodu, hasta_basvuru_kodu, tani_kodu, tani_turu, birincil_tani,
  tani_zamani, hekim_kodu, kayit_zamani, ekleyen_kullanici_kodu, guncelleme_zamani, guncelleyen_kullanici_kodu
) VALUES
  ('TN001','H0002','B0003','I10','KESIN',1,'2025-12-21 09:45:00','P0001','2025-12-21 09:46:00','P0001',NULL,NULL),
  ('TN002','H0003','B0002','A41.9','ON_TANI',1,'2025-12-21 08:50:00','P0003','2025-12-21 08:51:00','P0003',NULL,NULL),
  ('TN003','H0001','B0001','R50.9','ON_TANI',1,'2025-12-21 09:05:00','P0001','2025-12-21 09:06:00','P0001',NULL,NULL);

COMMIT;
