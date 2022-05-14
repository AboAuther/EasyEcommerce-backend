package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductId            string    `json:"productId" gorm:"column:product_id"`
	ProductName          string    `json:"productName" gorm:"column:product_name"`
	ProductIntro         string    `json:"productIntro" gorm:"column:product_intro"`
	CategoryId           int       `json:"categoryId" gorm:"column:category_id"`
	ProductCoverImg      string    `json:"productCoverImg" gorm:"column:product_cover_img"`
	ProductBanner        string    `json:"productBanner" gorm:"column:product_banner"`
	OriginalPrice        float64   `json:"originalPrice" gorm:"column:original_price"`
	SellingPrice         float64   `json:"sellingPrice" gorm:"column:selling_price"`
	StockNum             int       `json:"stockNum" gorm:"column:stock_num"`
	ClickNum             int       `json:"click_num" gorm:"column:click_num"`
	CreateUser           string    `json:"createUser" gorm:"column:create_user"`
	UpdateUser           string    `json:"updateUser" gorm:"column:update_user"`
	ProductDetailContent string    `json:"productDetailContent" gorm:"column:product_detail_content"`
	IsDeleted            time.Time `json:"isDeleted" gorm:"column:is_deleted"`
	CreatedAt            time.Time `json:"createAt" gorm:"column:created_at"`
	UpdatedAt            time.Time `json:"updateAt" gorm:"column:updated_at"`
}

type ProductEvaluation struct { //商品评价
	gorm.Model
	CreateUser string `json:"createUser" gorm:"column:create_user"`
	ProductId  string `json:"productId" gorm:"column:product_id"`
	OrderId    string `json:"orderId" gorm:"column:order_id"`
	Evaluation string `json:"evaluation" gorm:"column:evaluation"`
	Star       int    `json:"star" gorm:"column:star"`
}

type CombineProductAndEvaluation struct {
	Product     Product             `json:"product"`
	Evaluations []ProductEvaluation `json:"evaluation"`
}
