package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
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

func AddProduct(c *gin.Context) {
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
	product_ids := make([]string, 0)
	user := c.Query("userID")
	if err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Product{}).Where("create_user=?", user).Select("product_id").Scan(&product_ids).Error; err != nil {
			return err
		}
		if err := tx.Where("product_id IN (?)", product_ids).Find(&orders).Error; err != nil {
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
