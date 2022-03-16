package controllers

import (
	"account-service-app/config"
	"account-service-app/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ------
// Inisialisasi koneksi database
// ------
var db *gorm.DB
func init() {
	db = config.MysqlConnect()
}


// ------ 
// Menampilkan daftar data user
// ------ 
func listUser() {
	users := []entity.User{}
	db.Find(&users)
	fmt.Println("Daftar User")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("# \tNama \t\tGender \tNomor HP \tAlamat")
	fmt.Println("----------------------------------------------------------------")
	for _, user := range users {
		fmt.Print(user.ID, "\t")
		fmt.Print(user.Nama, "\t\t")
		fmt.Print(user.JenisKelamin, "\t")
		fmt.Print(user.NomorHP, "\t")
		fmt.Print(user.Alamat, "\t")
		fmt.Println()
	}
	fmt.Println("----------------------------------------------------------------")
}

// ------ 
// Mengambil data user berdasarkan nomor hp
// ------ 
func findUserByPhone(nomorHP string) (entity.User, bool) {
	user := entity.User{}
	err := db.Where("nomor_hp = ?", nomorHP).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

// ------ 
// Mengambil data user berdasarkan nomor hp dengan membawa relasi tabel Topup
// ------ 
func findUserByPhoneWithTopup(nomorHP string) (entity.User, bool) {
	user := entity.User{}
	err := db.Preload("TopUp").Where("nomor_hp = ?", nomorHP).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

// ------ 
// Mengambil data user berdasarkan nomor hp dengan membawa relasi tabel Transfer
// ------ 
func findUserByPhoneWithTransferKe(nomorHP string) (entity.User, bool) {
	user := entity.User{}
	err := db.Preload("TransferKe").Where("nomor_hp = ?", nomorHP).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

func TopUp() {
	// List user
	listUser()

	// Input nomor HP
	nomorHP := ""
	fmt.Print("Ketikkan nomor hp untuk top-up (q untuk keluar): ")
	fmt.Scanln(&nomorHP)

	//  Skip jika input = q
	if nomorHP != "q" {
		user, userExist := findUserByPhone(nomorHP)
		if !userExist {
			fmt.Println("User tidak ditemukan")
			TopUp()
		}
		
		// Input nominal top-up
		var nominalTopUp uint = 0
		fmt.Print("Masukkan jumlah Top-up: ")
		fmt.Scanln(&nominalTopUp)

		// Update saldo di tabel user
		user.Saldo = user.Saldo + nominalTopUp
		db.Save(&user)

		// Insert ke tabel top-up
		topUp := entity.TopUp{
			UserId: user.ID,
			Nominal: nominalTopUp,
		}
		db.Create(&topUp)

		// Menampilkan hasil
		fmt.Println("----------------------")
		fmt.Println("Top-up Berhasil!")
		fmt.Println("----------------------")
		fmt.Println("Nama:", user.Nama)
		fmt.Println("Jumlah top-up:", topUp.Nominal)
		fmt.Println("Saldo Sekarang:", user.Saldo)
		fmt.Println("Waktu: ", topUp.CreatedAt)
	} 
}

func Transfer() {
	// Display list user
	listUser()

	// Input nomor HP pengirim
	nomorHPPengirim := ""
	fmt.Print("Ketikkan nomor hp anda: ")
	fmt.Scanln(&nomorHPPengirim)

	// Skip jika input = q
	if nomorHPPengirim != "q" {
		// Cari user pengirim berdasarkan nomor HP nya
		userPengirim, userPengirimExist := findUserByPhone(nomorHPPengirim)
		if !userPengirimExist {
			fmt.Println("User tidak ditemukan")
			Transfer()  // Kembali ke inputan HP jika tidak ditemukan
		}

		// Input nomor HP penerima
		nomorHPPenerima := ""
		fmt.Print("Ketikkan nomor hp yang akan di transfer (q untuk keluar): ")
		fmt.Scanln(&nomorHPPenerima)

		// Skip jika input = q
		if nomorHPPenerima != "q" {

			// Cari user dengan nomor HP yang sama
			userPenerima, userPenerimaExist := findUserByPhone(nomorHPPenerima)
			if !userPenerimaExist {
				fmt.Println("User tidak ditemukan")
				Transfer()  // Kembali ke inputan HP jika tidak ditemukan
			}

			// Masukan nominal transfer, 
			// Jika saldo tidak cukup akan diminta input lagi
			var nominalTransfer uint = 0
			for {
				fmt.Print("Masukkan jumlah Top-up: ")
				fmt.Scanln(&nominalTransfer)
				if userPengirim.Saldo >= nominalTransfer {
					break
				} 
				fmt.Println("!! : Saldo anda tidak cukup!")
			}

			// Proses pengurangan dan penambahan saldo
			userPengirim.Saldo = userPengirim.Saldo - nominalTransfer
			userPenerima.Saldo = userPenerima.Saldo + nominalTransfer
			db.Save(&userPengirim)
			db.Save(&userPenerima)

			// Insert data ke tabel transfer
			transfer := entity.Transfer{
				UserId: userPengirim.ID,
				UserPenerimaId: userPenerima.ID,
				Nominal: nominalTransfer,
			}
			db.Create(&transfer)

			// Menampilkan hasil
			fmt.Println("----------------------")
			fmt.Println("Transfer Berhasil!")
			fmt.Println("----------------------")
			fmt.Println("Nama Pengirim \t:", userPengirim.Nama)
			fmt.Println("Nama Penerima \t:", userPenerima.Nama)
			fmt.Println("Jumlah transfer \t:", transfer.Nominal)
			fmt.Println("Saldo Anda Sekarang\t:", userPengirim.Saldo)
			fmt.Println("Waktu: ", transfer.CreatedAt)
		}
	}
}

func HistoryTopUp() {
	// Display list user
	listUser()

	// Input nomor HP pengirim
	nomorHP := ""
	fmt.Print("Ketikkan nomor hp anda: ")
	fmt.Scanln(&nomorHP)

	// Skip jika input = q
	if nomorHP != "q" {
		// Cari user pengirim berdasarkan nomor HP nya
		user, userExist := findUserByPhoneWithTopup(nomorHP)
		if !userExist {
			fmt.Println("User tidak ditemukan")
			HistoryTopUp()  // Kembali ke inputan HP jika tidak ditemukan
		}

		// Menampilkan hasil
		fmt.Println("------------------------------------------------------")
		fmt.Println("Data Transaksi Topup - ", user.Nama, "| Saldo sekarang:", user.Saldo)
		fmt.Println("------------------------------------------------------")
		for _, topUp := range user.TopUp {
			fmt.Println(topUp.Nominal, "\t", topUp.CreatedAt)
		}

		// Menampilkan hasil kosong
		if len(user.TransferKe) <= 0 {
			fmt.Println("Tidak ada riwayat transaksi")
		}
	}
}

func HistoryTransfer() {
	// Display list user
	listUser()

	// Input nomor HP
	nomorHP := ""
	fmt.Print("Ketikkan nomor hp anda: ")
	fmt.Scanln(&nomorHP)

	// Skip jika input = q
	if nomorHP != "q" {
		// Cari user berdasarkan nomor HP nya
		user, userExist := findUserByPhoneWithTransferKe(nomorHP)
		if !userExist {
			fmt.Println("User tidak ditemukan")
			HistoryTransfer()  // Kembali ke inputan HP jika tidak ditemukan
		}

		// Menampilkan hasil
		fmt.Println("------------------------------------------------------")
		fmt.Println("Data Transaksi Kirim Transfer - ", user.Nama, "| Saldo sekarang:", user.Saldo)
		fmt.Println("------------------------------------------------------")
		for _, topUp := range user.TransferKe {
			user := entity.User{}
			db.Find(&user, topUp.UserPenerimaId)
			fmt.Println("Transfer", topUp.Nominal, "\t", "ke", user.Nama, "pada waktu: ", topUp.CreatedAt)
		}

		// Menampilkan hasil kosong
		if len(user.TransferKe) <= 0 {
			fmt.Println("Tidak ada riwayat transaksi")
		}
	}
}