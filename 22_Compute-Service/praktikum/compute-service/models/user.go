package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UsersResponse struct {
	UserID uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func (u *User) TableName() string {
	return "user"
}
