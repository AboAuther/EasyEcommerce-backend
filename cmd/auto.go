package main

import (
	"EasyEcommerce-backend/internal/mysql"
)

func main() {
	adddress := new(mysql.MessageBoard)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)

}
