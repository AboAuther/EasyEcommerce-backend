package main

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
)

func main() {
	adddress := new(models.SaleData)
	mysql.InitDB()
	mysql.DB.AutoMigrate(adddress)
	//os.Mkdir("./images", 0777)

}
