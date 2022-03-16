package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/account_service_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err != nil {
		panic("Tidak dapat terhubung dengan database")
	}
	return db
}