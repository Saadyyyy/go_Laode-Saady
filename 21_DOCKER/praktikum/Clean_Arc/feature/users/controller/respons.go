package controller

import "gorm.io/gorm"

type UserRespons struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
