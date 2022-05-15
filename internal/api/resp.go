package api

import "EasyEcommerce-backend/internal/mysql/models"

type VerifyStruct struct {
	ID     string `json:"id"`
	Verify bool   `json:"verify"`
}

type SellerAggregation struct {
	models.Seller
	NickName       string  `json:"nickName" gorm:"column:nick_name"`
	Mobile         string  `json:"mobile" gorm:"column:mobile"`
	Region         string  `json:"region" gorm:"column:region"`
	Address        string  `json:"address" gorm:"column:address"`
	TotalPrice     float64 `json:"totalPrice"`
	YesterdayPrice float64 `json:"yesterDayPrice"`
}
