package models

import "gorm.io/gorm"

type CombineProductsAndOrder struct {
	Products []OrderOfProduct `json:"products"`
	Extra    OrderOfExtra     `json:"extra"`
}

type OrderOfProduct struct {
	ProductId    string  `json:"productId"`
	ProductImg   string  `json:"productImg"`
	ProductIntro string  `json:"productIntro"`
	SellingPrice float64 `json:"sellingPrice" gorm:"column:selling_price"`
	BuyNum       int     `json:"buyNum"`
}
type OrderOfExtra struct {
	UserId      string `json:"userID"`
	Mobile      string `json:"mobile"`
	UserAddress string `json:"userAddress"`
}

type Order struct {
	gorm.Model
	OrderId          string  `json:"orderId" gorm:"column:order_id"`
	UserId           string  `json:"userId" gorm:"column:user_id"`
	ProductId        string  `json:"productId" gorm:"column:product_id"`
	ProductImg       string  `json:"productImg"`
	Description      string  `json:"description" gorm:"column:description"`
	Mobile           string  `json:"mobile" gorm:"column:mobile"`
	TotalPrice       float64 `json:"totalPrice" gorm:"column:total_price"`
	ProductPrice     float64 `json:"productPrice" gorm:"column:product_price"`
	ProductNum       int     `json:"productNum" gorm:"column:product_num"`
	PayStatus        string  `json:"payStatus" gorm:"column:pay_status"`
	OrderStatus      string  `json:"orderStatus" gorm:"column:order_status"`
	EvaluationStatus bool    `json:"evaluationStatus" gorm:"column:evaluation_status"`
	UserAddress      string  `json:"userAddress" gorm:"column:user_address"`
}

type ShoppingCart struct {
	gorm.Model
	UserId          string  `json:"userId" gorm:"column:user_id"`
	ProductId       string  `json:"productId" gorm:"column:product_id"`
	ProductCoverImg string  `json:"productCoverImg" gorm:"column:product_cover_img"`
	Description     string  `json:"description" gorm:"column:description"`
	ProductPrice    float64 `json:"productPrice" gorm:"column:product_price"`
	ProductNum      int     `json:"productNum" gorm:"column:product_num"`
}
