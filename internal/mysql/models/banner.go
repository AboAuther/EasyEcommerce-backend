package models

import "time"

type Banner struct {
	BannerID      int       `json:"bannerID" gorm:"column:banner_id"`
	Url           string    `json:"url" gorm:"column:url"`
	RedirectUrl   string    `json:"redirectUrl" gorm:"column:redirect_url"`
	OriginalPrice float64   `json:"originalPrice" gorm:"column:original_price"`
	SellingPrice  float64   `json:"sellingPrice" gorm:"column:selling_price"`
	ProductIntro  string    `json:"productIntro" gorm:"column:product_intro"`
	CreateUser    string    `json:"createUser" gorm:"column:create_user"`
	UpdateUser    string    `json:"updateUser" gorm:"column:update_user"`
	IsDeleted     bool      `json:"isDeleted" gorm:"column:is_deleted"`
	CreatedAt     time.Time `json:"createAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updateAt" gorm:"column:updated_at"`
}
