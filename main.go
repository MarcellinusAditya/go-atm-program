package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MarcellinusAditya/go-atm-program/controllers/authcontroller"
	"github.com/MarcellinusAditya/go-atm-program/controllers/transactioncontroller"
	"github.com/MarcellinusAditya/go-atm-program/database"
)

var isLoggedIn = false
var currentUser authcontroller.CurrentAcc

func main() {
	database.ConnectDatabase()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !isLoggedIn {
			fmt.Println("\n=== ATM CLI ===")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("0. Keluar")
			fmt.Print("Pilih menu: ")
		} else {
			fmt.Printf("\n=== ATM - Selamat Datang %s ===\n", currentUser.Name)
			fmt.Println("1. Cek Saldo")
			fmt.Println("2. Deposit")
			fmt.Println("3. Tarik Tunai")
			fmt.Println("4. Transfer")
			fmt.Println("5. Logout")
			fmt.Print("Pilih menu: ")
		}

		scanner.Scan()
		menu := strings.TrimSpace(scanner.Text())

		if !isLoggedIn {
			switch menu {
			case "1":
				fmt.Println("== Register ==")
				fmt.Print("Masukkan username: ")
				var name string
				fmt.Scanln(&name)
				fmt.Print("Masukkan pin: ")
				var pin string
				fmt.Scanln(&pin)
					
				if name != "" && pin != ""{
					authcontroller.Register(name, pin)
				}else{
					fmt.Println("nama dan pin perlu diisi")
				}
			case "2":
				if isLoggedIn {
					fmt.Println("Sudah login sebagai", currentUser)
						return 
					}
					fmt.Println("== Login ==")
					fmt.Print("Masukkan id akun: ")
					var id int
					fmt.Scanln(&id)
					fmt.Print("Masukkan pin akun: ")
					var pin string
					fmt.Scanln(&pin)
										
					isLoggedIn, currentUser = authcontroller.Login(id,pin)
				
			case "0":
				fmt.Println("Sampai jumpa!")
				return
			default:
				fmt.Println("Menu tidak valid.")
			}
		} else {
			switch menu {
			case "1":
				fmt.Println(" ")
				fmt.Println("== Cek Saldo ==")
				transactioncontroller.CekSaldo(currentUser.Id)
				fmt.Println("")
			case "2":
				fmt.Println(" ")
				fmt.Println("== Deposit ==")
				fmt.Print("Masukkan nominal uang yang ingin dideposit (minimum Rp. 50.000): ")
				var amount float64
				fmt.Scanln(&amount)

				if amount <50000{
					fmt.Print("Nominal uang minimal Rp. 50.000")
					return
				}
				transactioncontroller.Deposit(amount, currentUser.Id)

			case "3":
				fmt.Println(" ")
				fmt.Println("== Withdraw ==")
				fmt.Print("Masukkan nominal uang yang ingin ditarik (minimum Rp. 50.000): ")
				var amount float64
				fmt.Scanln(&amount)

				if amount <50000{
					fmt.Print("Nominal uang minimal Rp. 50.000")
					return
				}
				transactioncontroller.Withdraw(amount, currentUser.Id)
			case "4":
				fmt.Println(" ")
				fmt.Println("== Transfer ==")
				fmt.Print("Masukkan id akun yang ingin ditransfer: ")
				var targetId int
				fmt.Scanln(&targetId)
				transactioncontroller.Transfer(targetId, currentUser.Id)

			case "5":
				isLoggedIn = false
				currentUser.Id = 0
				currentUser.Name =""
			default:
				fmt.Println("Menu tidak valid.")
			}
		}
	}
}
