package main

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
)

func main() {
	adddress := new(models.ShoppingAddress)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)

}
