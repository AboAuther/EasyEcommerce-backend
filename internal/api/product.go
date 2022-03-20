package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := mysql.DB.Order("click_num desc").Limit(15).Find(&productionLists); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if len(productionLists) == 15 {
		entity = Entity{
			Code:      http.StatusOK,
			Msg:       OperateOk.String(),
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
