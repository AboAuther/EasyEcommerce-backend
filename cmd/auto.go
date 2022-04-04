package main

import "EasyEcommerce-backend/internal/mysql"

func main() {
	adddress := new(mysql.ShoppingAddress)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)
}
