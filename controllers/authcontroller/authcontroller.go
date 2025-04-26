package authcontroller

import (
	"fmt"
	"math/rand"

	"strconv"
	"time"

	"github.com/MarcellinusAditya/go-atm-program/database"
	"github.com/MarcellinusAditya/go-atm-program/models"

	"gorm.io/gorm"
)

type CurrentAcc struct{
	Id int
	Name string
}

func Register(name string, pin string) {
	id, _ := strconv.Atoi(generateIdNumber())

	newUser := models.Account{
		Id : id,
		Name: name,
		Pin: pin,
		Balance: 0,
	}

	database.DB.Create(&newUser)
	fmt.Println("Registrasi berhasil untuk", name ," dengan id akun ", id)

}

func Login(id int, pin string) (bool,CurrentAcc){
	var user models.Account
	userInput := models.Account{
		Id: id,
		Pin: pin,
	}

	if err := database.DB.First(&user, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			fmt.Println("Id atau pin salah")
			return false, CurrentAcc{Id:0,Name:""}
		default:
			fmt.Println("error")
			return false, CurrentAcc{Id:0,Name:""}
		}
	}

	if userInput.Pin != user.Pin{
		fmt.Println("Id atau pin salah")
		return false, CurrentAcc{Id:0,Name:""}
	}
	loggedAcc := CurrentAcc{
		Id: user.Id,
		Name: user.Name,
	}

	fmt.Println("Login Berhasil")
	return true, loggedAcc
}

func generateIdNumber() string {
	rand.Seed(time.Now().UnixNano())                 
	return fmt.Sprintf("%08d", rand.Intn(100000000)) 
}