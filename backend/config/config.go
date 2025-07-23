package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("HATA: .env dosyası yüklenemedi")
	}

	// .env'den verileri oku
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// PostgreSQL bağlantı cümlesi (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// GORM ile bağlan
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}

	DB = db
	fmt.Println("✅ Veritabanına başarıyla bağlanıldı.")
}
