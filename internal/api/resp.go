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

type UserVerify struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
	IsSeller bool   `json:"is_seller"`
}

type SaleData struct {
	models.SaleData
	Time string `json:"time"`
	Date string `json:"date"`
}
type SaleDataAggregation struct {
	TotalPrice      float64    `json:"totalPrice"`
	YesterdayPrice  float64    `json:"yesterdayPrice"`
	TotalOrders     int64      `json:"totalOrders"`
	YesterdayOrders int64      `json:"yesterdayOrders"`
	TotalUsers      int64      `json:"totalUsers"`
	YesterdayUsers  int64      `json:"yesterdayUsers"`
	SaleNum         int        `json:"saleNum"`
	SaleOutNum      int        `json:"saleOutNum"`
	AllOrders       int64      `json:"allOrders"`
	RejectOrders    int64      `json:"rejectOrders"`
	SaleDatas       []SaleData `json:"saleDatas"`
}
