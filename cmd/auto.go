package main

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
)

func main() {
	adddress := new(models.ShoppingCart)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)

}
