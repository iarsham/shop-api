package db

import (
	"bufio"
	"fmt"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/models"
	"github.com/iarsham/shop-api/internal/validators"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func CreateAdminUser(logs *common.Logger) {
	var phone, passWord string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your phone: ")
		scanner.Scan()
		phone = scanner.Text()
		if len(phone) != 0 && validators.IrPhoneValidate(phone) {
			break
		}
	}
	for {
		fmt.Print("Enter your password: ")
		scanner.Scan()
		passWord = scanner.Text()
		if len(passWord) != 0 {
			break
		}
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.DefaultCost)
	adminUser := models.Users{
		Phone:    phone,
		Password: string(hashPass),
		IsAdmin:  true,
		IsActive: true,
	}
	db := GetDB()
	if err := db.Create(&adminUser).Error; err != nil {
		logs.Fatal(err.Error())
	}
	logs.Info("Admin Created Successfully.")
	os.Exit(1)
}
