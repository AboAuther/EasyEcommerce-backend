package api

import (
	"EasyEcommerce-backend/internal/client"
	"EasyEcommerce-backend/internal/mysql"
	"EasyEcommerce-backend/internal/mysql/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	entity := Entity{
		Code: int(OperateFail),
		Msg:  OperateFail.String(),
		Data: "Wrong username or password",
	}
	var user, user1 models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	isExisted, err := client.IsExisted(user.UserId)
	if err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if !isExisted {
		entity.Msg = OperateFail.String()
		entity.Data = "The user does not exist"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Where(models.User{
		UserId:   user.UserId,
		Password: user.Password,
	}).First(&user1).Error; err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if user1.UserId != "" {
		entity = Entity{
			Code:    http.StatusOK,
			Success: true,
			Msg:     OperateOk.String(),
			Data:    "Login successfully",
		}
		session.Set("status", true)
		err := session.Save()
		if err != nil {
			entity.Msg = OperateFail.String()
			entity.Data = err
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func UserRegister(c *gin.Context) {
	entity := Entity{
		Code:  int(OperateFail),
		Msg:   OperateFail.String(),
		Total: 0,
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	isExisted, err := client.IsExisted(user.UserId)
	if err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if isExisted {
		entity.Msg = OperateFail.String()
		entity.Data = "The user is already existed"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if mysql.DB.Create(&user).Error != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err.Error() + "can not add user"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	} else {
		entity.Code = int(OperateOk)
		entity.Msg = OperateOk.String()
		entity.Success = true
		entity.Data = "Register successful"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func UserEdit(c *gin.Context) {
	entity := Entity{
		Code:  int(OperateFail),
		Msg:   OperateFail.String(),
		Total: 0,
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	isExisted, err := client.IsExisted(user.UserId)
	if err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if isExisted {
		entity.Msg = OperateFail.String()
		entity.Data = "The user is already existed"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if mysql.DB.Save(&user).Error != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err.Error() + "can not save user data"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	} else {
		entity.Code = int(OperateOk)
		entity.Msg = OperateOk.String()
		entity.Success = true
		entity.Data = "Edit successfully"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func GetAddress(c *gin.Context) {
	entity := failedEntity
	id := c.Query("userID")
	var addresses []models.ShoppingAddress
	if err := mysql.DB.Where("create_user = ?", id).Order("id desc").Find(&addresses).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = addresses
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func AddAddress(c *gin.Context) {
	entity := failedEntity
	var address models.ShoppingAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if address.CreateUser != "" {
		if address.Default {
			if err := mysql.DB.Model(models.ShoppingAddress{}).Where("create_user=?", address.CreateUser).Update("default", false).Error; err != nil {
				entity.Data = err.Error()
				c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
				return
			}
		}
		if err := mysql.DB.Save(&address).Error; err != nil {
			entity.Data = err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
			return
		}
		entity = successEntity
		entity.Data = "Add successfully"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Data = "Invalid data"
	c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
	return
}

func DeleteAddress(c *gin.Context) {
	entity := failedEntity
	id := c.Param("id")
	if err := mysql.DB.Where("id=?", id).Delete(&models.ShoppingAddress{}).Error; err != nil {
		entity.Data = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = "Deleted successfully"
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}
