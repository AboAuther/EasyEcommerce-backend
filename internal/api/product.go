package api

import (
	"EasyEcommerce-backend/internal/client"
	"EasyEcommerce-backend/internal/mysql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func ProductionList(c *gin.Context) {
	var productionLists []mysql.Product
	entity := Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	if err := mysql.DB.Order("click_num desc").Limit(15).Find(&productionLists).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if len(productionLists) == 15 {
		entity = Entity{
			Code:      http.StatusOK,
			Msg:       OperateOk.String(),
			Success:   true,
			Total:     15,
			TotalPage: 1,
			Data:      productionLists,
		}
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	} else {
		entity.Data = "repository do not have enough products"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
}

func ProductionListByCategory(c *gin.Context) {
	var productionLists []mysql.Product
	entity := Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	model := new(mysql.Product)
	categoryStr := c.Query("category")
	category, _ := strconv.Atoi(categoryStr)
	model.CategoryId = category
	switch category {
	case AllCategory:
		model = new(mysql.Product)
	}
	priceStr := c.Query("price")
	price, _ := strconv.Atoi(priceStr)
	priceRange := PriceMap[price]
	err := mysql.DB.Where(model).Where("selling_price between ? and ?", priceRange[0], priceRange[1]).Order("click_num desc").Find(&productionLists).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	if len(productionLists) > 0 {
		entity = Entity{
			Code:      http.StatusOK,
			Msg:       OperateOk.String(),
			Success:   true,
			Total:     len(productionLists),
			TotalPage: client.GetPage(len(productionLists), 12),
			Data:      productionLists,
		}
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	} else {
		entity.Data = "repository do not have enough products"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
}
func Banner(c *gin.Context) {
	entity := Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	var banner []*mysql.Banner
	if err := mysql.DB.Order("updated_at desc").Find(&banner).Limit(4).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if len(banner) == 4 {
		entity = Entity{
			Code:      http.StatusOK,
			Msg:       OperateOk.String(),
			Success:   true,
			Total:     4,
			TotalPage: 1,
			Data:      banner,
		}
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	} else {
		entity.Data = "repository do not have enough products"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
}

func ProductByID(c *gin.Context) {
	entity := Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	var product mysql.Product
	var evaluations, tmp []mysql.ProductEvaluation

	id := c.Param("id")
	if err := mysql.DB.Where(mysql.Product{ProductId: id}).Find(&product).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Where(mysql.ProductEvaluation{ProductId: id}).Find(&evaluations).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	var combine mysql.CombineProductAndEvaluation
	if product.ProductId != "" {
		combine.Product = product
	}
	for _, evaluation := range evaluations {
		if evaluation.ProductId != "" && evaluation.Evaluation != "" {
			tmp = append(tmp, evaluation)
		}
	}
	combine.Evaluations = tmp
	entity = Entity{
		Code:      http.StatusOK,
		Msg:       OperateOk.String(),
		Success:   true,
		Total:     1,
		TotalPage: 1,
		Data:      combine,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
func ProductByName(c *gin.Context) {
	entity := Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	var product []mysql.Product
	name := c.Param("name")
	name = fmt.Sprintf("%s%s%s", "%", name, "%")
	if err := mysql.DB.Where("product_name LIKE ?", name).Find(&product).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = Entity{
		Code:      http.StatusOK,
		Msg:       OperateOk.String(),
		Success:   true,
		Total:     len(product),
		TotalPage: 1,
		Data:      product,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return

}
