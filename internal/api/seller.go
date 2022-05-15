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
	c.SaveUploadedFile(file, filepath)
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
