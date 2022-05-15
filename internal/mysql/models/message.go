package models

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

type MessageBoard struct {
	gorm.Model
	UserID   string
	Nickname string
	Topic    string
	Content  string
	IsVerify bool
}
