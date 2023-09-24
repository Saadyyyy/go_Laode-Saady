package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	User_id uint   `json:"user_id"`
	User    User
}
