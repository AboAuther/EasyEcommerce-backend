package api

import (
	"EasyEcommerce-backend/internal/client"
	"EasyEcommerce-backend/internal/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {

	entity := Entity{
		Code: int(OperateFail),
		Msg:  OperateFail.String(),
		Data: "Wrong username or password",
	}
	var user, user1 mysql.User
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
	if err := mysql.DB.Where(mysql.User{
		UserId:   user.UserId,
		Password: user.Password,
	}).First(&user1); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if user1.UserId != "" {
		entity = Entity{
			Code: http.StatusOK,
			Msg:  OperateOk.String(),
			Data: "Login successfully",
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
	var user mysql.User
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
		entity.Msg = OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
