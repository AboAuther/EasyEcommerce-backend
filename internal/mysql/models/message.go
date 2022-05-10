package models

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Title   string
	Content string
}

type MessageBoard struct {
	gorm.Model
	UserID   string
	Nickname string
	Topic    string
	Content  string
	IsVerify bool
}
