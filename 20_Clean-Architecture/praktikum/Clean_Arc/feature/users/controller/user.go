package controller

import (
	users "Laode_Saady/20_Clean-Architecture/praktikum/Clean_Arc/feature/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userCase users.UserCaseInterface
}

func New(uc users.UserCaseInterface) *UserController {
	return &UserController{
		userCase: uc,
	}
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	userList, err := uc.userCase.GetAll()
	if err != nil {
		// Handle the error here
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting all users",
		})
	}

	// Handle the successful case here
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success getting all users",
		"users":   userList,
	})
}

func (uc *UserController) CreateUser(c echo.Context) error {
	input := new(UserRespons)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "Eror Bind user",
		})
	}

	data := users.Users{
		Email:    input.Email,
		Password: input.Password,
	}
	err := uc.userCase.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Err create user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Succes create user",
	})
}

func (uc *UserController) Login(c echo.Context) error {
	input := new(UserRespons)

	errBind := c.Bind(input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "Error Parshing login request",
		})
	}
	// memanggil usecase login
	user, token, err := uc.userCase.Login(users.Users{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "Internal server eror",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Succes Login",
		"user":    user,
		"token":   token,
	})

}
