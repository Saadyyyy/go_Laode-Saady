package router

import (
	"saady/app/middleware"
	"saady/feature/users/controller"
	"saady/feature/users/repository"
	uc "saady/feature/users/useCase"

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
