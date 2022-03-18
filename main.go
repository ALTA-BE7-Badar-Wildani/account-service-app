package main

import (
	"account-service-app/config"
	"account-service-app/controllers"
	"account-service-app/entity"
	"fmt"
	"os"
	"os/exec"
)

func init() {
	db := config.MysqlConnect()
	db.AutoMigrate(&entity.User{}, &entity.Transfer{}, &entity.TopUp{})
}

func main() {
	menu := ""
	fmt.Println("----------------------------")
	fmt.Println("Account Service App")
	fmt.Println("----------------------------")
	fmt.Println("1) Add User")
	fmt.Println("2) List User")
	fmt.Println("3) Update User")
	fmt.Println("4) Delete User")
	fmt.Println("5) Top-up")
	fmt.Println("6) Transfer")
	fmt.Println("7) History Top-up")
	fmt.Println("8) History Transfer")
	fmt.Println("----------------------------")
	fmt.Println("0) Keluar")
	fmt.Println("----------------------------")
	fmt.Print("Masukkan pilihan menu anda: ")
	fmt.Scanln(&menu)
	fmt.Println("----------------------------")

	menuController(menu)
}

func menuController(menu string) {
	clearScreen()
	switch menu {
	case "1":
		fmt.Println("Add User")
		fmt.Println("-------------------------")
		controllers.AddUser()
	case "2":
		fmt.Println("List User")
		fmt.Println("-------------------------")
		controllers.ListUser()
	case "3":
		fmt.Println("Update User")
		fmt.Println("-------------------------")
		controllers.UpdateUser()
	case "4":
		fmt.Println("Delete User")
		fmt.Println("-------------------------")
		controllers.DeleteUser()
	case "5":
		fmt.Println("Topup")
		fmt.Println("-------------------------")
		controllers.TopUp()
	case "6":
		fmt.Println("Transfer")
		fmt.Println("-------------------------")
		controllers.Transfer()
	case "7":
		fmt.Println("History Top Up")
		fmt.Println("-------------------------")
		controllers.HistoryTopUp()
	case "8":
		fmt.Println("History Transfer")
		fmt.Println("-------------------------")
		controllers.HistoryTransfer()
	case "0":
		fmt.Println("----------------------------")
		fmt.Println("  Terima kasih telah bertransaksi")
		fmt.Println("----------------------------")
		os.Exit(3)
	}

	fmt.Println("--------------------------- ")
	fmt.Println("Tekan enter untuk lanjut... ")
	fmt.Scanln()
	clearScreen()
	main()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
