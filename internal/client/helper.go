package client

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func IsExisted(userId string) (bool, error) {
	var user models.User
	if mysql.IsMissing(mysql.DB.Where(models.User{UserId: userId}).First(&user)) {
		return false, nil
	}
	if user.UserId != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func GetPage(length, size int) int {
	if length%size == 0 {
		return length / size
	}
	return length/size + 1
}

func EnsureLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	statusInterface := session.Get("status")
	if statusInterface != nil {
		status := statusInterface.(string)
		if status == "Authorized" {
			return true
		}
	}
	return false
}
func MakeSaleData(product string, price float64) error {
	var saledata models.SaleData
	var user string
	tomorrow := time.Now().AddDate(0, 0, 1)
	zeroTimeTomorrow := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
		0, 0, 0, 0, tomorrow.Location())
	if err := mysql.DB.Model(&models.Product{}).Select("create_user").Where("product_id =?", product).Scan(&user).Error; err != nil {
		return err
	}

	isMiss := mysql.IsMissing(mysql.DB.Where("user_id = ?", user).Where("created_at =?", zeroTimeTomorrow).First(&saledata))
	if isMiss {
		if err := mysql.DB.Save(&models.SaleData{UserID: user, Amount: price, Model: gorm.Model{CreatedAt: zeroTimeTomorrow}}).Error; err != nil {
			return err
		}
	} else {
		saledata.Amount += price
		saledata.CreatedAt = zeroTimeTomorrow
		if err := mysql.DB.Save(&saledata).Error; err != nil {
			return err
		}
	}
	return nil
}
