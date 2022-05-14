package models

import "time"

type Seller struct {
	UserId          string    `json:"userID" gorm:"column:user_id" gorm:"primaryKey"`
	Identity        string    `json:"identity" gorm:"identity"`
	RegisterAddress string    `json:"registerAddress" gorm:"register_address"`
	LicenseUrl      string    `json:"licenseUrl" gorm:"license_url"`
	HygieneUrl      string    `json:"hygieneUrl" gorm:"hygiene_url"` //卫生许可
	ShopName        string    `json:"shopName" gorm:"shop_name"`
	CreatedAt       time.Time `json:"createdAt" gorm:"created_at"`
	IsVerify        bool
}
