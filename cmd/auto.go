package main

import (
	"EasyEcommerce-backend/internal/mysql"
)

func main() {
	adddress := new(mysql.Notice)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)
	
}
