package main

import (
	"EasyEcommerce-backend/internal/api"
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/utils"
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

	product := r.Group("/api/product")
	{
		product.GET("/list", api.ProductionList)
		product.GET("/banner", api.Banner)
		product.GET("/listByCategory", api.ProductionListByCategory)
		//product.GET("/info/:id", ProductHandler.ProductInfoHandler)
		//product.POST("/add", ProductHandler.AddProductHandler)
		//product.POST("/edit", ProductHandler.EditProductHandler)
		//product.POST("/delete/:id", ProductHandler.DeleteProductHandler)
	}
	user := r.Group("/api/user")
	{
		user.POST("/login", api.UserLogin)
		user.POST("register", api.UserRegister)
		//user.GET("/list", UserHandler.UserListHandler)
		//user.GET("/info/:id", UserHandler.UserInfoHandler)
		//user.POST("/add", UserHandler.AddUserHandler)
		//user.POST("/edit", UserHandler.EditUserHandler)
		//user.POST("/delete/:id", UserHandler.DeleteUserHandler)
	}
	//banner := r.Group("/api/banner")
	//{
	//	banner.GET("/list", BannerHandler.BannerListHandler)
	//	banner.GET("/info/:id", BannerHandler.BannerInfoHandler)
	//	banner.POST("/add", BannerHandler.AddBannerHandler)
	//	banner.POST("/edit", BannerHandler.EditBannerHandler)
	//	banner.POST("/delete/:id", BannerHandler.DeleteBannerHandler)
	//}
	//
	//category := r.Group("/api/category")
	//{
	//	category.GET("/list", CategoryHandler.CategoryListHandler)
	//	category.GET("/list4backend", CategoryHandler.CategoryList4BackendHandler)
	//	category.GET("/info/:id", CategoryHandler.CategoryInfoHandler)
	//	category.POST("/add", CategoryHandler.AddCategoryHandler)
	//	category.POST("/edit", CategoryHandler.EditCategoryHandler)
	//	category.POST("/delete/:id", CategoryHandler.DeleteCategoryHandler)
	//}

	//order := r.Group("/api/order")
	//{
	//	order.GET("/list", OrderHandler.OrderListHandler)
	//	order.GET("/info/:id", OrderHandler.OrderInfoHandler)
	//	order.POST("/add", OrderHandler.AddOrderHandler)
	//	order.POST("/edit", OrderHandler.EditOrderHandler)
	//	order.POST("/delete/:id", OrderHandler.DeleteOrderHandler)
	//}

	port := utils.GetStringEnv("PORT", ":8080")
	if err := r.Run(port); err != nil {
		panic("run failed!")
	}
}
