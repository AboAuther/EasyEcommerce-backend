package mysql

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
	Tag                  string    `json:"tag" gorm:"column:tag"`
	SellStatus           int       `json:"sellStatus" gorm:"column:sell_status"`
	CreateUser           string    `json:"createUser" gorm:"column:create_user"`
	UpdateUser           string    `json:"updateUser" gorm:"column:update_user"`
	ProductDetailContent string    `json:"productDetailContent" gorm:"column:product_detail_content"`
	IsDeleted            bool      `json:"isDeleted" gorm:"column:is_deleted"`
	CreatedAt            time.Time `json:"createAt" gorm:"column:created_at"`
	UpdatedAt            time.Time `json:"updateAt" gorm:"column:updated_at"`
}

type Stock struct {
	ProductId    string `json:"productId"`
	ProductCount int    `json:"productCount"`
}

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

type Order struct {
	gorm.Model
	OrderId      string  `json:"orderId" gorm:"column:order_id"`
	UserId       string  `json:"userId" gorm:"column:user_id"`
	ProductId    string  `json:"productId" gorm:"column:product_id"`
	Name         string  `json:"name" gorm:"column:name"`
	Description  string  `json:"description" gorm:"column:description"`
	Mobile       string  `json:"mobile" gorm:"column:mobile"`
	TotalPrice   float64 `json:"totalPrice" gorm:"column:total_price"`
	ProductPrice float64 `json:"productPrice" gorm:"column:product_price"`
	ProductNum   int     `json:"productNum" gorm:"column:product_num"`
	PayStatus    string  `json:"payStatus" gorm:"column:pay_status"`
	OrderStatus  string  `json:"orderStatus" gorm:"column:order_status"`
	UserAddress  string  `json:"userAddress" gorm:"column:user_address"`
}

type User struct {
	UserId    string    `json:"userID" gorm:"column:user_id"`
	NickName  string    `json:"nickName" gorm:"column:nick_name"`
	Mobile    string    `json:"mobile" gorm:"column:mobile"`
	Password  string    `json:"password" gorm:"column:password"`
	Region    string    `json:"region" gorm:"column:region"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted bool      `json:"isDeleted" gorm:"column:is_deleted"`
	IsLocked  bool      `json:"isLocked" gorm:"column:is_locked"`
	CreatedAt time.Time `json:"createAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updateAt" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

type ShoppingAddress struct {
	gorm.Model
	CreateUser string `json:"createUser" gorm:"column:create_user"`
	Name       string `json:"name" gorm:"column:name"`
	Region     string `json:"region" gorm:"column:region"`
	Address    string `json:"address" gorm:"column:address"`
	Mobile     string `json:"mobile" gorm:"column:mobile"`
	Default    bool   `json:"default" gorm:"column:default"`
}
type ProductEvaluation struct {
	gorm.Model
	CreateUser   string `json:"createUser" gorm:"column:create_user"`
	ProductId    string `json:"productId" gorm:"column:product_id"`
	ProductName  string `json:"productName" gorm:"column:product_name"`
	ProductNum   int    `json:"productNum" gorm:"column:product_num"`
	ProductIntro string `json:"productIntro" gorm:"column:product_intro"`
	Evaluation   string `json:"evaluation" gorm:"column:evaluation"`
	Star         int    `json:"star" gorm:"column:star"`
}

type CombineProductAndEvaluation struct {
	Product     Product             `json:"product"`
	Evaluations []ProductEvaluation `json:"evaluation"`
}

type CombineProductsAndAddress struct {
	Products []Product `json:"products"`
	Data     Order     `json:"data"`
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
