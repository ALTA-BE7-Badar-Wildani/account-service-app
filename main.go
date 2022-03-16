package main

import (
	"account-service-app/controllers"
	"fmt"
)

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
	fmt.Println("0) Exit")
	fmt.Println("----------------------------")
	fmt.Print("Masukkan pilihan menu anda: ")
	fmt.Scanln(&menu)
	fmt.Println("----------------------------")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	menuController(menu)
}

func menuController(menu string) {
	switch menu {
	case "1":
		controllers.AddUser()
	case "2":
		controllers.ListUser()
	case "3":
		controllers.UpdateUser()
	case "4":
		controllers.DeleteUser()
	case "5":
		controllers.TopUp()
	case "6":
		controllers.Transfer()
	case "7":
		controllers.HistoryTopUp()
	case "8":
		controllers.HistoryTransfer()
	}
	fmt.Println("----------------------------")
	fmt.Println("  Terima kasih telah bertransaksi")
}
