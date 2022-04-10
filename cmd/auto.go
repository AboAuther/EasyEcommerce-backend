package main

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/utils"
	"fmt"
)

func main() {
	adddress := new(mysql.Order)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)
	fmt.Println(utils.CreateRandomNumber())

}
