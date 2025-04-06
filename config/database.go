package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbconn *gorm.DB
var once sync.Once // Agar ConnectDB hanya dipanggil sekali

func ConnectDB() {
	once.Do(func() { // Pastikan hanya dipanggil sekali
		errenv := godotenv.Load()
		if errenv != nil {
			log.Fatal("Gagal membaca .env: ", errenv)
		}

		DB_HOST := os.Getenv("DB_HOST")
		DB_NAME := os.Getenv("DB_NAME")
		DB_USER := os.Getenv("DB_USER")
		DB_PASS := os.Getenv("DB_PASS")
		DB_PORT := os.Getenv("DB_PORT")

		// Cek apakah variabel environment sudah terisi
		if DB_HOST == "" || DB_NAME == "" || DB_USER == "" || DB_PASS == "" || DB_PORT == "" {
			log.Fatal("Pastikan semua variabel database di .env sudah diisi!")
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
		if err != nil {
			log.Fatal("Koneksi Gagal: ", err)
		}

		// Cek apakah koneksi berhasil
		sqlDB, err := database.DB()
		if err != nil {
			log.Fatal("Gagal mendapatkan koneksi database: ", err)
		}

		// Ping database untuk memastikan koneksi hidup
		err = sqlDB.Ping()
		if err != nil {
			log.Fatal("Gagal ping database: ", err)
		}

		dbconn = database
		fmt.Println("Database berhasil terkoneksi!")

		
	})
}

func GetDB() *gorm.DB {
	ConnectDB()
	if dbconn == nil {
		log.Fatal("Database is not connected. Pastikan ConnectDB() sudah dipanggil!")
	}
	return dbconn
}