package client

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
