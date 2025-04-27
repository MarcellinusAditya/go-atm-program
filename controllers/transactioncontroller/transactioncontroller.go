package transactioncontroller

import (
	"fmt"

	"github.com/MarcellinusAditya/go-atm-program/database"
	"github.com/MarcellinusAditya/go-atm-program/models"
	"gorm.io/gorm"
)

func CekSaldo(id int) {
	var user models.Account

	if err := database.DB.First(&user, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			fmt.Println("Id atau pin salah")
			return 
		default:
			fmt.Println("error")
			return 
		}
	}

	fmt.Printf("Saldo akun atas nama %s sebesar Rp.%.2f", user.Name, user.Balance)
}
func Deposit(amount float64, id int) {
	var user models.Account

	if err := database.DB.First(&user, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			fmt.Println("Id atau pin salah")
			return 
		default:
			fmt.Println("error")
			return 
		}
	}

	finalAmount := user.Balance + amount
	user.Balance = finalAmount

	if database.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		fmt.Println("tidak dapat update balance")
		return
	}

	transaction := models.Transaction{
		AccountID: id,
		Type: "Deposit",
		Amount: amount,
	}

	database.DB.Create(&transaction)

	fmt.Printf("Saldo telah ditambahkan sebesar Rp.%.2f \n", amount)
	fmt.Printf("Total saldo sekarang sebesar Rp.%.2f", finalAmount)

}
func Withdraw(amount float64, id int) {
	var user models.Account

	if err := database.DB.First(&user, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			fmt.Println("Id atau pin salah")
			return 
		default:
			fmt.Println("error")
			return 
		}
	}

	if user.Balance<amount{
		fmt.Println("Saldo tidak mencukupi")
		return
	}

	finalAmount := user.Balance - amount
	user.Balance = finalAmount

	if database.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		fmt.Println("tidak dapat update balance")
		return
	}

	transaction := models.Transaction{
		AccountID: id,
		Type: "Withdraw",
		Amount: amount,
	}

	database.DB.Create(&transaction)

	fmt.Printf("Saldo telah ditarik sebesar Rp.%.2f \n", amount)
	fmt.Printf("Total saldo sekarang sebesar Rp.%.2f", finalAmount)

}
func Transfer(targetId int, userId int) {
	var targetUser, user models.Account

	if err := database.DB.First(&targetUser, targetId).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			fmt.Println("User tidak ditemukan")
			return 
		default:
			fmt.Println("error")
			return 
		}
	}
	if err := database.DB.First(&user, userId).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			fmt.Println("User tidak ditemukan")
			return 
		default:
			fmt.Println("error")
			return 
		}
	}

	fmt.Printf("User yang ingin ditransfer atas nama %s \n", targetUser.Name)
	fmt.Println("Masukan nominal pengiriman: (minimal Rp. 50.000)")
	var nominal float64
	fmt.Scanln(&nominal)

	if nominal < 50000{
		fmt.Println("nominal minimal Rp. 50.000")
		return
	}
	if nominal > user.Balance{
		fmt.Println("Saldo anda tidak mencukupi")
		return
	}

	finalTargetAmount := targetUser.Balance+nominal
	targetUser.Balance = finalTargetAmount

	finalUserAmount := user.Balance-nominal
	user.Balance=finalUserAmount

	if database.DB.Model(&targetUser).Where("id = ?", targetId).Updates(&targetUser).RowsAffected == 0 {
		fmt.Println("Gagal transfer")
		return
	}
	if database.DB.Model(&user).Where("id = ?", userId).Updates(&user).RowsAffected == 0 {
		fmt.Println("Gagal transfer")
		return
	}

	targetTransaction := models.Transaction{
		AccountID: targetId,
		Type: "transfer_in",
		Amount: nominal,
	}

	database.DB.Create(&targetTransaction)

	userTransaction := models.Transaction{
		AccountID: userId,
		Type: "transfer_out",
		Amount: nominal,
	}

	database.DB.Create(&userTransaction)

	fmt.Println("Transfer berhasil")

}
