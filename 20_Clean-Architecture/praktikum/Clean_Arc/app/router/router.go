package router

import (
	"Laode_Saady/20_Clean-Architecture/praktikum/belajar-go-echo/app/middleware"
	"Laode_Saady/20_Clean-Architecture/praktikum/belajar-go-echo/feature/users/controller"
	"Laode_Saady/20_Clean-Architecture/praktikum/belajar-go-echo/feature/users/repository"
	uc "Laode_Saady/20_Clean-Architecture/praktikum/belajar-go-echo/feature/users/useCase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := repository.New(db)
	userUseCase := uc.New(userRepository)
	userController := controller.New(userUseCase)
	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetAllUsers, middleware.JwtMiddleware())
	e.POST("/users/login", userController.Login)
}
