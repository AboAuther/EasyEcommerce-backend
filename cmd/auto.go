package main

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
)

func main() {
	adddress := new(models.ProductEvaluation)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)

}
