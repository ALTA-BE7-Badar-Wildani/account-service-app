package controllers

import (
	"account-service-app/entity"
	"fmt"
	"time"
)

func AddUser() {
	addUser := entity.User{}

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&addUser.Nama)
	fmt.Print("Masukkan Jenis Kelamin: ")
	fmt.Scanln(&addUser.JenisKelamin)
	fmt.Print("Masukkan Alamat: ")
	fmt.Scanln(&addUser.Alamat)
	fmt.Print("Masukkan Nomor HP: ")
	fmt.Scanln(&addUser.NomorHP)
	fmt.Print("Masukkan Saldo: ")
	fmt.Scanln(&addUser.Saldo)
	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&addUser.Email)

	// Input data tanggal lahir
	var tahun, bulan, tanggal int
	fmt.Print("Masukkan Tanggal Lahir (tanggal): ")
	fmt.Scanln(&tanggal)
	fmt.Print("Masukkan Tanggal Lahir (bulan): ")
	fmt.Scanln(&bulan)
	fmt.Print("Masukkan Tanggal Lahir (tahun): ")
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
	var IDUser uint
	fmt.Println("Ketikkan nomor ID User untuk melakukan Update")
	fmt.Scanln(&IDUser)

	// Mengambil data user dari database
	user := entity.User{}
	db.Find(&user, IDUser)
	fmt.Println(user)

	var nama, jeniskelamin, nomorhp, email, alamat string
	var saldo uint

	fmt.Println("----------------------------")
	fmt.Print("1) Update Nama (", user.Nama, ") : ")
	fmt.Scanln(&nama)
	if nama != "" {
		user.Nama = nama
	}
	fmt.Print("2) Update Jenis Kelamin (", user.JenisKelamin, ") : ")
	fmt.Scanln(&jeniskelamin)
	if jeniskelamin != "" {
		user.JenisKelamin = jeniskelamin
	}
	fmt.Print("3) Update Nomor HP (", user.NomorHP, ") : ")
	fmt.Scanln(&nomorhp)
	if nomorhp != "" {
		user.NomorHP = nomorhp
	}
	fmt.Print("4) Update Saldo (", user.Saldo, ") : ")
	fmt.Scanln(&saldo)
	if saldo != user.Saldo {
		user.Saldo = saldo
	}
	fmt.Print("5) Update Email (", user.Email, ") : ")
	fmt.Scanln(&email)
	if email != "" {
		user.Email = email
	}
	fmt.Print("6) Update Alamat (", user.Alamat, ") : ")
	fmt.Scanln(&alamat)
	if alamat != "" {
		user.Alamat = alamat
	}
	// Update data tanggal lahir
	var tahun, bulan, tanggal int
	fmt.Println("7) Tanggal Lahir (", user.TanggalLahir, ") : ")
	fmt.Print("--  Update Tanggal Lahir: ")
	fmt.Scanln(&tanggal)
	fmt.Print("--  Update Bulan Lahir: ")
	fmt.Scanln(&bulan)
	fmt.Print("--  Update Tahun Lahir: ")
	fmt.Scanln(&tahun)
	if tanggal != 0 || bulan != 0 || tahun != 0 {
		user.TanggalLahir = time.Date(tahun, time.Month(bulan), tanggal, 0, 0, 0, 0, time.Local)
	}

	db.Save(user) //menyimpan seluruh update yang dilakukan
}

func DeleteUser() {
	var id uint
	fmt.Print("Masukkan ID User yang akan anda hapus:")
	fmt.Scanln(&id)

	//Delete User sesuai dengan ID User
	user := entity.User{}
	db.Delete(&user, id)
}
