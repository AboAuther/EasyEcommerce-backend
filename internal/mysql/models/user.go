package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId    string    `json:"userID" gorm:"column:user_id"`
	NickName  string    `json:"nickName" gorm:"column:nick_name"`
	Mobile    string    `json:"mobile" gorm:"column:mobile"`
	Password  string    `json:"password" gorm:"column:password"`
	Region    string    `json:"region" gorm:"column:region"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted bool      `json:"isDeleted" gorm:"column:is_deleted"`
	IsLocked  bool      `json:"isLocked" gorm:"column:is_locked"`
	CreatedAt time.Time `json:"createAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updateAt" gorm:"column:updated_at"`
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
