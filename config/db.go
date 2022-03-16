package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Tidak ada file konfigurasi dotenv tersedia")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err != nil {
		panic("Tidak dapat terhubung dengan database")
	}
	return db
}