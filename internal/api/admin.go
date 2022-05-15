package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSellerForVerify(c *gin.Context) {
	entity := failedEntity
	var sellers []models.Seller
	if err := mysql.DB.Where("is_verify= false").Find(&sellers).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = sellers
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func VerifySeller(c *gin.Context) {
	entity := failedEntity
	var model VerifyStruct
	if err := c.ShouldBindJSON(&model); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if model.Verify {
		if err := mysql.DB.Model(&models.Seller{}).Where("user_id= ?", model.ID).Update("is_verify", true).Error; err != nil {
			entity.Data = err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
		entity = successEntity
		entity.Data = "Verify successfully"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

}

func VerifyMessage(c *gin.Context) {
	entity := failedEntity
	var model VerifyStruct
	if err := c.ShouldBindJSON(&model); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	id, _ := strconv.Atoi(model.ID)
	if model.Verify {
		if err := mysql.DB.Model(&models.MessageBoard{}).Where("id = ?", id).Update("is_verify", true).Error; err != nil {
			entity.Data = err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
		entity = successEntity
		entity.Data = "Verify successfully"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	} else {
		if err := mysql.DB.Where("id= ?", model.ID).Delete(&models.MessageBoard{}).Error; err != nil {
			entity.Data = err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
		entity = successEntity
		entity.Data = "delete"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

}
