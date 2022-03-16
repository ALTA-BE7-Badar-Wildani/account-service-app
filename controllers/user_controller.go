package controllers

import (
	"account-service-app/config"
	"account-service-app/entity"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.MysqlConnect()
}

func AddUser() {
	addUser := entity.User{}

	fmt.Println("Masukkan Nama:")
	fmt.Scanln(&addUser.Nama)
	fmt.Println("Masukkan Jenis Kelamin:")
	fmt.Scanln(&addUser.JenisKelamin)
	fmt.Println("Masukkan Alamat:")
	fmt.Scanln(&addUser.Alamat)
	fmt.Println("Masukkan Nomor HP:")
	fmt.Scanln(&addUser.NomorHP)
	fmt.Println("Masukkan Saldo:")
	fmt.Scanln(&addUser.Saldo)
	fmt.Println("Masukkan Email:")
	fmt.Scanln(&addUser.Email)
	fmt.Println("Masukkan Tanggal Lahir")
	fmt.Scanln(&addUser.TanggalLahir)

	tx := db.Save(&addUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	if tx.RowsAffected == 0 {
		fmt.Println("add User failed")
	}
	fmt.Println("Add User Successfully")
}

func ListUser() {
	var listUser []entity.User
	tx := db.Find(&listUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	for _, value := range listUser {
		fmt.Println("ID:", value.ID, "Nama:", value.Nama, "Jenis Kelamin:", value.JenisKelamin, "Alamat:", value.Alamat, "Nomor HP:", value.NomorHP, "Saldo:", value.Saldo, "Email:", value.Email, "Tanggal Lahir:", value.TanggalLahir)
	}
}

func UpdateUser() {
	fmt.Println("Update User")
}

func DeleteUser() {
	fmt.Println("Delete User")
}
