package repository

import (
	"saady/app/middleware"
	users "saady/feature/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataInterface {
	return &UserRepository{
		db: db,
	}
}

// GetAll implements model.DataInterface.
func (UserRepo *UserRepository) GetAll() ([]users.Users, error) {
	user := []users.Users{}
	result := UserRepo.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (UserRepo *UserRepository) Insert(data users.Users) error {
	var input = User{
		Email:    data.Email,
		Password: data.Password,
	}
	tx := UserRepo.db.Create(&input)
	if tx != nil {
		return tx.Error
	}
	return nil
}

// Login implements model.DataInterface.
func (UserRepo *UserRepository) Login(data users.Users) (users.Users, string, error) {
	result := UserRepo.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&data)

	if result.Error != nil {
		return data, "", result.Error
	}
	var token string
	if result.RowsAffected > 0 {
		var errToken error
		token, errToken = middleware.CreateToken(int(data.ID))
		if errToken != nil {
			return data, "", errToken
		}
	}
	return data, token, nil
}
