package repositorie

import (
	"Laode_Saady/18_Middleware/praktikum/config"
	"Laode_Saady/18_Middleware/praktikum/middleware"
	"Laode_Saady/18_Middleware/praktikum/model"
)

func CheckLogin(email string, password string) (model.User, string, error) {
	var data model.User
	// select * from users where email = `email` and password = `password`
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return model.User{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middleware.CreateToken(int(data.ID))
		if errToken != nil {
			return model.User{}, "", errToken
		}
	}
	return data, token, nil
}
