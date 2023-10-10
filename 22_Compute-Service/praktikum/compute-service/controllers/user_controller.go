package controllers

import (
	"net/http"
	"praktikum_22/helpers"
	"praktikum_22/repositories"
	"praktikum_22/responses"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(ur repositories.UserRepository) *UserController {
	return &UserController{
		repo: ur,
	}
}

func (uc *UserController) GetUserController(c echo.Context) error {
	responseData, err := uc.repo.Select()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}

	var usersResponse = []responses.UserResponse{}
	for _, value := range responseData {
		usersResponse = append(usersResponse, responses.UserResponse{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success recieve user data", usersResponse))
}
