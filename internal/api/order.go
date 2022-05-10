package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"EasyEcommerce-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetOrder(c *gin.Context) {
	entity := failedEntity
	id := c.Query("userID")
	var orders []models.Order
	if err := mysql.DB.Where(models.Order{UserId: id}).Find(&orders).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Total = len(orders)
	entity.Data = orders
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func MakeOrders(c *gin.Context) {
	entity := failedEntity
	var data models.CombineProductsAndOrder
	if err := c.ShouldBindJSON(&data); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	var orders []models.Order
	orderId := "Ebuy" + utils.CreateRandomNumber()
	for i := 0; i < len(data.Products); i++ {
		order := models.Order{
			OrderId:          orderId,
			UserId:           data.Extra.UserId,
			ProductId:        data.Products[i].ProductId,
			ProductImg:       data.Products[i].ProductImg,
			TotalPrice:       data.Products[i].SellingPrice * float64(data.Products[i].BuyNum),
			ProductPrice:     data.Products[i].SellingPrice,
			ProductNum:       data.Products[i].BuyNum,
			Description:      data.Products[i].ProductIntro,
			Mobile:           data.Extra.Mobile,
			UserAddress:      data.Extra.UserAddress,
			PayStatus:        "Paid",
			OrderStatus:      "Bought",
			EvaluationStatus: false,
		}
		orders = append(orders, order)
	}
	if err := mysql.DB.Save(&orders).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Bought successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func DeleteOrder(c *gin.Context) {
	entity := failedEntity
	orderID := c.Query("order_id")
	if err := mysql.DB.Where("order_id=?", orderID).Delete(&models.Order{}).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Delete order successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func AddCart(c *gin.Context) {
	entity := failedEntity
	var shoppingCart models.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		entity.Data = err.Error
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Save(&shoppingCart).Error; err != nil {
		entity.Data = err.Error
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Add successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func EditCart(c *gin.Context) {
	entity := failedEntity
	var shoppingCart models.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		entity.Data = err.Error
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if shoppingCart.ProductNum > 0 {
		if err := mysql.DB.UpdateColumn("product_num", shoppingCart.ProductNum).Error; err != nil {
			entity.Data = err.Error
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
	} else {
		if err := mysql.DB.Where(models.ShoppingCart{UserId: shoppingCart.UserId, ProductId: shoppingCart.ProductId}).Delete(models.ShoppingCart{}).Error; err != nil {
			entity.Data = err.Error
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
	}

	entity = successEntity
	entity.Data = "Edit successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func DeleteCart(c *gin.Context) {
	entity := failedEntity
	var shoppingCart []models.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		entity.Data = err.Error
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	for _, cart := range shoppingCart {
		if err := mysql.DB.Where(models.ShoppingCart{UserId: cart.UserId, ProductId: cart.ProductId}).Delete(models.ShoppingCart{}).Error; err != nil {
			entity.Data = err.Error
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
	}
	entity = successEntity
	entity.Data = "Deleted successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func AddEvaluation(c *gin.Context) {
	entity := failedEntity
	var evaluation models.ProductEvaluation
	if err := c.ShouldBindJSON(&evaluation); err != nil {
		entity.Data = err.Error
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&evaluation).Error; err != nil {
			return err
		}
		if err := tx.Model(models.Order{}).Where("order_id=? and product_id =?", evaluation.OrderId, evaluation.ProductId).Update("evaluation_status", true).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Evaluated successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
