package model

import "time"

type Users struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdateAt  time.Time `json:"UpdateAt"`
}

type DataInterface interface {
	Insert(Users) error
	GetAll() ([]Users, error)
	Login(Users) (Users, string, error)
}

type UserCaseInterface interface {
	Create(Users) error
	GetAll() ([]Users, error)
	Login(Users) (Users, string, error)
}
