package api

import (
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotice(c *gin.Context) {
	entity := failedEntity
	var notices []models.Notice
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
func GetMessageBoard(c *gin.Context) {
	entity := failedEntity
	var messages []models.MessageBoard
	isVerify := c.Query("is_verify")
	if err := mysql.DB.Where("is_verify = ?", isVerify).Order("id desc").Find(&messages).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = messages
	entity.Total = len(messages)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func AddMessage(c *gin.Context) {
	entity := failedEntity
	var message models.MessageBoard
	if err := c.ShouldBindJSON(&message); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	message.IsVerify = false
	if err := mysql.DB.Save(&message).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Message successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
func AddNotice(c *gin.Context) {
	entity := failedEntity
	var notice models.Notice
	if err := c.ShouldBindJSON(&notice); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Save(&notice).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Add notice successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func VerifyMessage(c *gin.Context) {
	entity := failedEntity
	var message []*models.MessageBoard
	if err := c.ShouldBindJSON(&message); err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	for _, board := range message {
		board.IsVerify = true
	}
	if err := mysql.DB.Save(message).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Verify successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
