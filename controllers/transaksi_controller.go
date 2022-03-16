package controllers

import (
	"account-service-app/config"
	"account-service-app/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB
func init() {
	db = config.MysqlConnect()
}


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

func findUserByPhone(nomorHP string) (entity.User, bool) {
	user := entity.User{}
	err := db.Where("nomor_hp = ?", nomorHP).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

func TopUp() {
	listUser()
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
	fmt.Println("Transfer")
}

func HistoryTopUp() {
	fmt.Println("HistoryTopUp")
}

func HistoryTransfer() {
	fmt.Println("History Transfer")
}