package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	failedEntity = Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	successEntity = Entity{
		Code:      int(OperateOk),
		Msg:       OperateOk.String(),
		Success:   true,
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
)

func GetOrder(c *gin.Context) {
	entity := failedEntity
	id := c.Query("username")
	var orders []mysql.Order
	if err := mysql.DB.Where(mysql.Order{UserId: id}).Find(&orders).Error; err != nil {
		entity.Data = err
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
	var data mysql.CombineProductsAndAddress
	if err := c.ShouldBindJSON(&data); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	var orders []mysql.Order
	orderId := "Ebuy" + utils.CreateRandomNumber()
	for i := 0; i < len(data.Products); i++ {
		order := mysql.Order{
			OrderId:      orderId,
			UserId:       data.Data.UserId,
			ProductId:    data.Products[i].ProductId,
			TotalPrice:   data.Data.TotalPrice,
			Name:         data.Data.Name,
			ProductPrice: data.Products[i].SellingPrice,
			ProductNum:   data.Products[i].StockNum,
			Description:  data.Products[i].ProductIntro,
			Mobile:       data.Data.Mobile,
			UserAddress:  data.Data.UserAddress,
			PayStatus:    "Paid",
			OrderStatus:  "Bought",
		}
		orders = append(orders, order)
	}
	if err := mysql.DB.Save(&orders).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Bought successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
func AddCart(c *gin.Context) {
	entity := failedEntity
	var shoppingCart mysql.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Save(&shoppingCart).Error; err != nil {
		entity.Data = err
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
	var shoppingCart mysql.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if shoppingCart.ProductNum > 0 {
		if err := mysql.DB.UpdateColumn("product_num", shoppingCart.ProductNum).Error; err != nil {
			entity.Data = err
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
	} else {
		if err := mysql.DB.Where(mysql.ShoppingCart{UserId: shoppingCart.UserId, ProductId: shoppingCart.ProductId}).Delete(mysql.ShoppingCart{}).Error; err != nil {
			entity.Data = err
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
	var shoppingCart []mysql.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	for _, cart := range shoppingCart {
		if err := mysql.DB.Where(mysql.ShoppingCart{UserId: cart.UserId, ProductId: cart.ProductId}).Delete(mysql.ShoppingCart{}).Error; err != nil {
			entity.Data = err
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
	}
	entity = successEntity
	entity.Data = "Deleted successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
