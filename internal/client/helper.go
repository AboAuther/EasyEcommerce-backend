package client

import "EasyEcommerce-backend/internal/mysql"

func IsExisted(name string) (bool, error) {
	var user mysql.User
	if err := mysql.DB.Where(mysql.User{UserId: name}).First(&user).Error; err != nil {
		return false, err
	}
	if user.UserId != "" {
		return true, nil
	} else {
		return false, nil
	}
}
