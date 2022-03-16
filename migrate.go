package main

import (
	"account-service-app/config"
	"account-service-app/entity"
)

func main() {
	db := config.MysqlConnect()	
	db.AutoMigrate(&entity.User{}, &entity.Transfer{}, &entity.TopUp{})
}