package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotice(c *gin.Context) {
	entity := failedEntity
	var notices []mysql.Notice
	if err := mysql.DB.Order("id desc").Limit(3).Find(&notices).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = notices
	entity.Total = len(notices)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
