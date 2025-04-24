package authcontroller

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/MarcellinusAditya/go-atm-program/database"
	"github.com/MarcellinusAditya/go-atm-program/models"
)

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

func generateIdNumber() string {
	rand.Seed(time.Now().UnixNano())                 
	return fmt.Sprintf("%08d", rand.Intn(100000000)) 
}