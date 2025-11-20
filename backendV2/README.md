# MedScreen Backend - Kurulum ve Çalıştırma Rehberi

Bu döküman, MedScreen Backend projesini yerel ortamınızda nasıl çalıştıracağınızı (sunucu ve veritabanı kurulumu) adım adım açıklar.

## Gereksinimler

*   **Go (Golang)**: Sürüm 1.24
*   **PostgreSQL**: Veritabanı sunucusu.

## 1. Veritabanı Kurulumu (PostgreSQL)

Projenin çalışması için bir PostgreSQL veritabanına ihtiyacı vardır.

1.  **PostgreSQL'i Kurun**: PostgreSQL 18.1 kurulumu yapın.
2.  **Veritabanı Oluşturun**: pgAdmin veya psql terminal aracı ile PostgreSQL sunucunuza bağlanın ve yeni bir veritabanı oluşturun.
    ```sql
    CREATE DATABASE medscreen_test;
    ```
3.  **Tabloları Oluşturun**: Proje ana dizininde bulunan `database_query.sql` dosyasının içeriğini kopyalayın ve oluşturduğunuz veritabanında çalıştırın. Gerekli tablolar ve başlangıç verileri oluşturulur.

## 2. Konfigürasyon (.env)

Projenin veritabanına bağlanabilmesi için ortam değişkenlerine ihtiyacı vardır. Proje ana dizininde `.env` adında bir dosya oluşturun ve aşağıdaki içeriği kendi ayarlarınıza göre düzenleyerek yapıştırın:

```env
# Test Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1234 #veya kendi postgres şifreniz
DB_NAME=medscreen_test
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8081
SERVER_HOST=0.0.0.0
GIN_MODE=debug

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8081
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=Content-Type,Authorization

# Logging
LOG_LEVEL=debug
LOG_FORMAT=json
```

## 3. Projeyi Çalıştırma

Terminali açın ve proje dizinine gidin.

### Adım 1: Bağımlılıkları Yükleme

Gerekli Go kütüphanelerini indirmek için şu komutu çalıştırın:

```bash
go mod download
```

### Adım 2: Sunucuyu Başlatma

Projeyi derlemeden doğrudan çalıştırmak için:

```bash
go run cmd/server/main.go
```

Eğer her şey doğru yapılandırıldıysa, terminalde sunucunun başladığına dair logları göreceksiniz (Örn: `Listening and serving HTTP on 0.0.0.0:8080`).


## Sorun Giderme

*   **Veritabanı Bağlantı Hatası**: `.env` dosyasındaki `DB_USER`, `DB_PASSWORD` ve `DB_NAME` bilgilerinin doğruluğundan emin olun. PostgreSQL servisinin çalıştığını kontrol edin.
*   **Port Hatası**: Eğer 8080 portu doluysa, `.env` dosyasından `SERVER_PORT` değerini değiştirebilirsiniz (Örn: 8081).

## Yapılacaklar

1. JWT tabanlı authentication sistemi
2. RBAC (Role Based Access Control) sistemi