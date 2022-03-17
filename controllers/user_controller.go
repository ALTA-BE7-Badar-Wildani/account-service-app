package controllers

import (
	"account-service-app/entity"
	"fmt"
	"time"
)

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

	// Input data tanggal lahir
	var tahun, bulan, tanggal int
	fmt.Println("Masukkan Tanggal Lahir (tanggal): ")
	fmt.Scanln(&tanggal)
	fmt.Println("Masukkan Tanggal Lahir (bulan): ")
	fmt.Scanln(&bulan)
	fmt.Println("Masukkan Tanggal Lahir (tahun): ")
	fmt.Scanln(&tahun)
	addUser.TanggalLahir = time.Date(tahun, time.Month(bulan), tanggal, 0, 0, 0, 0, time.Local)

	tx := db.Save(&addUser)
	if tx.Error != nil {
		panic(tx.Error)
	}

	fmt.Println("----------------------------")
	fmt.Println("Add User Successfully")
}

func ListUser() {
	var listUser []entity.User
	tx := db.Find(&listUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	for _, value := range listUser {
		fmt.Println("ID:", value.ID, "Nama:", value.Nama, "Jenis Kelamin:", value.JenisKelamin, "Nomor HP:", value.NomorHP, "Saldo:", value.Saldo, "Email:", value.Email, "Alamat:", value.Alamat, "Tanggal Lahir:", value.TanggalLahir)
	}
}

func UpdateUser() {
	// Input ID User untuk melakukan Update
	var IDUser int = 0
	fmt.Println("Ketikkan nomor ID User untuk melakukan Update")
	fmt.Scanln(&IDUser)

	// Mengambil data user dari database
	user := entity.User{}
	db.Find(&user, IDUser)
	fmt.Println(user)

	var nama string

	fmt.Println("Menu Update")
	fmt.Println("----------------------------")
	fmt.Print("1) Update Nama (", user.Nama, ") : ")
	fmt.Scanln(&nama)
	if nama != "" {
		user.Nama = nama
	}

	fmt.Println(user)

	fmt.Println("2) Update Jenis Kelamin")
	fmt.Println("3) Update Nomor HP")
	fmt.Println("4) Update Saldo")
	fmt.Println("5) Update Email")
	fmt.Println("6) Update Alamat")
	fmt.Println("7) Update Tanggal Lahir")

	menuUpdate := ""
	fmt.Println("Masukkan pilihan update anda:")
	fmt.Scanln(&menuUpdate)

	switch menuUpdate {
	case "1":
		var NamaBaru string
		fmt.Println("Masukkan nama baru anda:")
		fmt.Scanln(&NamaBaru)
		db.Model(&IDUser).Update("Nama", &NamaBaru)
	}
}

func DeleteUser() {
	var id uint
	fmt.Println("Masukkan ID User yang akan anda hapus:")
	fmt.Scanln(id)

	// Megnambil dataa dari db
	user := entity.User{}
	db.Find(&user, id)

	db.Delete(&user)
}
