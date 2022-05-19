package main

import (
	"EasyEcommerce-backend/internal/api"
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func main() {
	if level, err := logger.ParseLevel(utils.GetStringEnv("LOG_LEVEL", "debug")); err != nil {
		logger.Panic(err)
	} else {
		logger.SetLevel(level)
		logger.Info("Log Level Set")
	}
	// loading other components
	if err := mysql.InitDB(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(Cors())
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.StaticFS("/images", http.Dir("./images"))

	user := r.Group("/api/user")
	{
		user.POST("/login", api.UserLogin)
		user.POST("/register", api.UserRegister)
		user.POST("/edit", api.EditUser)
		user.GET("/getAddress", api.GetAddress)
		user.POST("/addAddress", api.AddAddress)
		user.POST("/deleteAddress/:id", api.DeleteAddress)
		user.GET("/getMessage", api.GetUser)
	}
	seller := r.Group("/api/seller")
	{
		seller.POST("/upload", api.UploadImage)
		seller.POST("/register", api.RegisterSeller)
		seller.POST("/addProduct", api.AddProduct)
		seller.POST("/editProduct", api.EditProduct)
		seller.GET("/getOrders", api.GetOrderForSeller)
		seller.GET("/getProduct", api.GetProducts)
		seller.GET("/getMessage", api.GetMessage)
		seller.GET("/getSaleMessage", api.GetSaleMessage)
	}
	product := r.Group("/api/product")
	{
		product.GET("/list", api.ProductionList)
		product.GET("/banner", api.Banner)
		product.GET("/listByCategory", api.ProductionListByCategory)
		product.GET("id/:id", api.ProductByID)
		product.GET("name/:name", api.ProductByName)
	}

	order := r.Group("/api/order")
	{
		order.GET("/getOrder", api.GetOrder)
		order.POST("/makeOrder", api.MakeOrders)
		order.GET("/getCart", api.GetCart)
		order.POST("/addCart", api.AddCart)
		order.POST("/editCart", api.EditCart)
		order.POST("/deleteCart", api.DeleteCart)
		order.POST("/addEvaluation", api.AddEvaluation)
		order.POST("/deleteOrder", api.DeleteOrder)
	}
	notice := r.Group("/api/message")
	{
		notice.GET("/notice", api.GetNotice)
		notice.GET("/board", api.GetMessageBoard)
		notice.POST("/addMessage", api.AddMessage)
		notice.POST("/addNotice", api.AddNotice)
	}
	admin := r.Group("/api/admin")
	{
		admin.POST("/verifyMessage", api.VerifyMessage)
		admin.POST("/verifySeller", api.VerifySeller)
		admin.GET("/getSeller", api.GetSellerForVerify)
	}
	port := utils.GetStringEnv("PORT", ":8080")
	if err := r.Run(port); err != nil {
		panic("run failed!")
	}
}
