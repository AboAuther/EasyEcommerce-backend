package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"EasyEcommerce-backend/internal/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"path"
	"strings"
	"time"
)

func UploadImage(c *gin.Context) {
	entity := failedEntity
	file, err := c.FormFile("file")
	if err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		entity.Data = "上传失败!只允许png,jpg,gif,jpeg文件"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	filepath := fmt.Sprintf("%s%s", "./images/", file.Filename)
	_ = c.SaveUploadedFile(file, filepath)
	entity = successEntity
	url := "http://" + c.Request.Host + "/images/" + file.Filename
	entity.Data = url
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func RegisterSeller(c *gin.Context) {
	entity := failedEntity
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	seller.IsVerify = false
	seller.CreatedAt = time.Now()
	if err := mysql.DB.Create(&seller).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Register successfully,wait verify!"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func EditProduct(c *gin.Context) {
	entity := failedEntity
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&product).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "edit product successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func AddProduct(c *gin.Context) {
	entity := failedEntity
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	product.ProductId = utils.CreateRandomString()
	if err := mysql.DB.Save(&product).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "add product successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func GetProducts(c *gin.Context) {
	entity := failedEntity
	var products []models.Product
	user := c.Query("userID")
	if err := mysql.DB.Where("create_user = ?", user).Find(&products).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = products
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func GetOrderForSeller(c *gin.Context) {
	entity := failedEntity
	var orders []models.Order
	productIds := make([]string, 0)
	user := c.Query("userID")
	if err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Product{}).Where("create_user=?", user).Select("product_id").Scan(&productIds).Error; err != nil {
			return err
		}
		if err := tx.Where("product_id IN (?)", productIds).Find(&orders).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = orders
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func GetMessage(c *gin.Context) {
	entity := failedEntity
	var seller models.Seller
	var user models.User
	var totalPriceToday, totalPriceYesterday sql.NullFloat64
	productIds := make([]string, 0)
	nowTomorrow := time.Now().AddDate(0, 0, 1)
	nowYesterday := time.Now().AddDate(0, 0, -1)
	zeroTimeTomorrow := time.Date(nowTomorrow.Year(), nowTomorrow.Month(), nowTomorrow.Day(),
		0, 0, 0, 0, nowTomorrow.Location())
	zeroTimeNow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
		0, 0, 0, 0, time.Now().Location())
	zeroTimeYesterday := time.Date(nowYesterday.Year(), nowYesterday.Month(), nowYesterday.Day(),
		0, 0, 0, 0, nowTomorrow.Location())
	userID := c.Query("userID")
	if err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id=?", userID).Find(&seller).Error; err != nil {
			return err
		}
		if err := mysql.DB.Where("user_id=?", userID).Find(&user).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Product{}).Where("create_user=?", userID).Select("product_id").Scan(&productIds).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("sum(total_price) as total").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeNow, zeroTimeTomorrow).Scan(&totalPriceToday).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("sum(total_price) as total").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeYesterday, zeroTimeNow).Scan(&totalPriceYesterday).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	sellerAgg := SellerAggregation{
		Seller:         seller,
		NickName:       user.NickName,
		Mobile:         user.Mobile,
		Region:         user.Region,
		Address:        user.Address,
		TotalPrice:     totalPriceToday.Float64,
		YesterdayPrice: totalPriceYesterday.Float64,
	}

	entity = successEntity
	entity.Data = sellerAgg
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func GetSaleMessage(c *gin.Context) {
	entity := failedEntity
	var saleData, lastMonthSaleData []models.SaleData
	var totalPriceToday, totalPriceYesterday sql.NullFloat64
	var orderCount, yesterdayOrderCount, userCount, yesterdayUserCount, allOrder int64
	productIds := make([]string, 0)

	userID := c.Query("userID")
	tomorrow := time.Now().AddDate(0, 0, 1)
	sevenDaysAgo := time.Now().AddDate(0, 0, -5)
	lastMonth := tomorrow.AddDate(0, -1, 0)
	lastMonthSevenDaysAgo := sevenDaysAgo.AddDate(0, -1, 0)
	nowYesterday := time.Now().AddDate(0, 0, -1)

	zeroTimeNow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
		0, 0, 0, 0, time.Now().Location())
	zeroTimeYesterday := time.Date(nowYesterday.Year(), nowYesterday.Month(), nowYesterday.Day(),
		0, 0, 0, 0, nowYesterday.Location())
	zeroTimeTomorrow := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
		0, 0, 0, 0, tomorrow.Location())
	zeroTimeSevenDaysAgo := time.Date(sevenDaysAgo.Year(), sevenDaysAgo.Month(), sevenDaysAgo.Day(),
		0, 0, 0, 0, sevenDaysAgo.Location())
	zeroTimeLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), lastMonth.Day(),
		0, 0, 0, 0, lastMonth.Location())
	zeroTimeLastMonthSevenDaysAgo := time.Date(lastMonthSevenDaysAgo.Year(), lastMonthSevenDaysAgo.Month(), lastMonthSevenDaysAgo.Day(),
		0, 0, 0, 0, lastMonthSevenDaysAgo.Location())

	if err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Where("created_at between ? and ?", zeroTimeSevenDaysAgo, zeroTimeTomorrow).Find(&saleData).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Where("created_at between ? and ?", zeroTimeLastMonthSevenDaysAgo, zeroTimeLastMonth).Find(&lastMonthSaleData).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if len(saleData) > 0 {
		saleData = utils.SortCreateTime(saleData)
		if len(saleData) < 7 {
			var saleDataTmp []models.SaleData
			diff := 7 - len(saleData)
			createTimeFormat := saleData[len(saleData)-1].CreatedAt
			createTime := time.Date(createTimeFormat.Year(), createTimeFormat.Month(), createTimeFormat.Day(),
				16, 0, 0, 0, createTimeFormat.Location())
			for i := 0; i < diff; i++ {
				createTime = createTime.AddDate(0, 0, -1)
				tmp := models.SaleData{UserID: userID, Amount: 0, Model: gorm.Model{CreatedAt: createTime}}
				saleDataTmp = append(saleDataTmp, tmp)
			}
			saleData = append(saleData, saleDataTmp...)
		}
	} else {
		var saleDataTmp []models.SaleData
		diff := 7 - len(saleData)
		createTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
			16, 0, 0, 0, time.Now().Location())
		for i := 0; i < diff; i++ {
			tmp := models.SaleData{UserID: userID, Amount: 0, Model: gorm.Model{CreatedAt: createTime}}
			saleDataTmp = append(saleDataTmp, tmp)
			createTime = createTime.AddDate(0, 0, -1)
		}
		saleData = append(saleData, saleDataTmp...)
	}

	if len(lastMonthSaleData) > 0 {
		lastMonthSaleData = utils.SortCreateTime(lastMonthSaleData)
		if len(lastMonthSaleData) < 7 {
			var saleDataTmp []models.SaleData
			diff := 7 - len(lastMonthSaleData)
			createTime := lastMonthSaleData[len(lastMonthSaleData)-1].CreatedAt
			for i := 0; i < diff; i++ {
				tmp := models.SaleData{UserID: userID, Amount: 0, Model: gorm.Model{CreatedAt: createTime}}
				saleDataTmp = append(saleDataTmp, tmp)
				createTime = createTime.AddDate(0, 0, -1)
			}
			lastMonthSaleData = append(lastMonthSaleData, saleDataTmp...)
		}
	} else {
		var saleDataTmp []models.SaleData
		diff := 7 - len(lastMonthSaleData)
		createTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
			16, 0, 0, 0, time.Now().Location())
		for i := 0; i < diff; i++ {
			tmp := models.SaleData{UserID: userID, Amount: 0, Model: gorm.Model{CreatedAt: createTime}}
			saleDataTmp = append(saleDataTmp, tmp)
			createTime = createTime.AddDate(0, 0, -1)
		}
		lastMonthSaleData = append(lastMonthSaleData, saleDataTmp...)
	}

	var saleDataResp []SaleData
	saleData = utils.SortCreateTimeASC(saleData)
	lastMonthSaleData = utils.SortCreateTimeASC(lastMonthSaleData)
	for _, data := range saleData {
		timeStr := data.CreatedAt.Format("2006-01-02 15:04:05")
		saleDataResp = append(saleDataResp, SaleData{SaleData: data, Time: timeStr, Date: "CurrentMonth"})
	}
	for _, data := range lastMonthSaleData {
		timeStr := data.CreatedAt.Format("2006-01-02 15:04:05")
		saleDataResp = append(saleDataResp, SaleData{SaleData: data, Time: timeStr, Date: "LastMonth"})
	}
	saleDataAgg := SaleDataAggregation{
		SaleDatas: saleDataResp,
	}
	if err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Product{}).Where("create_user=?", userID).Select("product_id").Scan(&productIds).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("distinct(order_id) as order").Where("product_id IN (?)", productIds).Count(&allOrder).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("distinct(order_id) as order").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeNow, zeroTimeTomorrow).Count(&orderCount).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("distinct(order_id) as order").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeYesterday, zeroTimeNow).Count(&yesterdayOrderCount).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("sum(total_price) as total").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeNow, zeroTimeTomorrow).Scan(&totalPriceToday).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("sum(total_price) as total").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeYesterday, zeroTimeNow).Scan(&totalPriceYesterday).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("distinct(user_id) as user").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeNow, zeroTimeTomorrow).Count(&userCount).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Order{}).Select("distinct(user_id) as user").Where("product_id IN (?)", productIds).Where("created_at between ? and ?", zeroTimeYesterday, zeroTimeNow).Count(&yesterdayUserCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	saleDataAgg.TotalPrice = totalPriceToday.Float64
	saleDataAgg.YesterdayPrice = totalPriceYesterday.Float64
	saleDataAgg.SaleNum = len(productIds)
	saleDataAgg.SaleOutNum = 0
	saleDataAgg.RejectOrders = 0
	saleDataAgg.AllOrders = orderCount
	saleDataAgg.TotalUsers = userCount
	saleDataAgg.YesterdayUsers = yesterdayUserCount
	saleDataAgg.TotalOrders = orderCount
	saleDataAgg.YesterdayOrders = yesterdayOrderCount

	entity = successEntity
	entity.Data = saleDataAgg
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func EditMessage(c *gin.Context) {
	entity := failedEntity
	var sellerData SellerMessage
	if err := c.ShouldBindJSON(&sellerData); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	if err := mysql.DB.Model(&models.Seller{}).Where("user_id= ?", sellerData.UserID).Updates(models.Seller{HygieneUrl: sellerData.HygieneUrl, LicenseUrl: sellerData.LicenseUrl}).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Update successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
