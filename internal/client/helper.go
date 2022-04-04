package client

import "EasyEcommerce-backend/internal/mysql"

func IsExisted(obj interface{}) (bool, error) {
	var user mysql.User
	if mysql.IsMissing(mysql.DB.Where(obj).First(&user)) {
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
