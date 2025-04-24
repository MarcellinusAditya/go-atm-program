package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcellinusAditya/go-atm-program/controllers/authcontroller"
	"github.com/MarcellinusAditya/go-atm-program/database"
	"github.com/urfave/cli/v2"
)

var isLoggedIn = false
var currentUser = ""

func main() {
	database.ConnectDatabase()
	app := &cli.App{
		Name:  "ATM CLI",
		Usage: "Simulasi program ATM berbasis terminal",
		Commands: []*cli.Command{
			{
				Name:  "register",
				Usage: "Daftarkan akun baru",
				Action: func(c *cli.Context) error {
					fmt.Println("== Register ==")
					fmt.Print("Masukkan username: ")
					var name string
					fmt.Scanln(&name)
					fmt.Print("Masukkan pin: ")
					var pin string
					fmt.Scanln(&pin)
					
					authcontroller.Register(name, pin)
					return nil
				},
			},
			{
				Name:  "login",
				Usage: "Login ke akun ATM",
				Action: func(c *cli.Context) error {
					if isLoggedIn {
						fmt.Println("Sudah login sebagai", currentUser)
						return nil
					}
					fmt.Println("== Login ==")
					fmt.Print("Masukkan username: ")
					var username string
					fmt.Scanln(&username)
					// abaikan autentikasi
					isLoggedIn = true
					currentUser = username
					fmt.Println("Login berhasil. Selamat datang,", username)
					return nil
				},
			},
			{
				Name:  "check-balance",
				Usage: "Cek saldo saat ini",
				Action: func(c *cli.Context) error {
					if !authCheck() {
						return nil
					}
					fmt.Println("Saldo saat ini: Rp 1.000.000")
					return nil
				},
			},
			{
				Name:  "deposit",
				Usage: "Setor uang ke akun",
				Action: func(c *cli.Context) error {
					if !authCheck() {
						return nil
					}
					fmt.Print("Masukkan jumlah yang akan disetor: ")
					var amount int
					fmt.Scanln(&amount)
					fmt.Println("Deposit berhasil sejumlah", amount)
					return nil
				},
			},
			{
				Name:  "withdraw",
				Usage: "Tarik uang dari akun",
				Action: func(c *cli.Context) error {
					if !authCheck() {
						return nil
					}
					fmt.Print("Masukkan jumlah yang akan ditarik: ")
					var amount int
					fmt.Scanln(&amount)
					fmt.Println("Penarikan berhasil sejumlah", amount)
					return nil
				},
			},
			{
				Name:  "transfer",
				Usage: "Transfer uang ke akun lain",
				Action: func(c *cli.Context) error {
					if !authCheck() {
						return nil
					}
					fmt.Print("Masukkan username tujuan: ")
					var toUser string
					fmt.Scanln(&toUser)

					fmt.Print("Masukkan jumlah transfer: ")
					var amount int
					fmt.Scanln(&amount)

					fmt.Println("Transfer berhasil sebesar", amount, "ke", toUser)
					return nil
				},
			},
			{
				Name:  "logout",
				Usage: "Keluar dari akun",
				Action: func(c *cli.Context) error {
					if !isLoggedIn {
						fmt.Println("Belum login")
						return nil
					}
					fmt.Println("Logout dari", currentUser)
					currentUser = ""
					isLoggedIn = false
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func authCheck() bool {
	if !isLoggedIn {
		fmt.Println("Anda belum login. Gunakan perintah `login` terlebih dahulu.")
		return false
	}
	return true
}
