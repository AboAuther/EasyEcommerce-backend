package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId      string         `json:"userID" gorm:"primaryKey"`
	Password    string         `json:"password" gorm:"column:password"`
	NickName    string         `json:"nickName" gorm:"column:nickname"`
	Mobile      string         `json:"mobile" gorm:"column:mobile"`
	Region      string         `json:"region" gorm:"column:region"`
	Address     string         `json:"address" gorm:"column:address"`
	Information string         `json:"information" gorm:"column:information"`
	CreatedAt   time.Time      `json:"createAt" gorm:"column:created_at"`
	UpdatedAt   time.Time      `json:"updateAt" gorm:"column:updated_at"`
	IsDeleted   gorm.DeletedAt `json:"isDeleted" gorm:"column:is_deleted"`
}

func (User) TableName() string {
	return "users"
}

type ShoppingAddress struct {
	gorm.Model
	CreateUser string `json:"createUser" gorm:"column:create_user"`
	Name       string `json:"name" gorm:"column:name"`
	Region     string `json:"region" gorm:"column:region"`
	Detail     string `json:"detail" gorm:"column:detail"`
	Mobile     string `json:"mobile" gorm:"column:mobile"`
	Default    bool   `json:"default" gorm:"column:default"`
}
