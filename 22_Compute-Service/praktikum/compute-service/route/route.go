package routes

import (
	"praktikum_22/controllers"
	"praktikum_22/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	// proses pembuatan object dan inject
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} \n",
	}))

	//endpoint
	e.GET("/users", userController.GetUserController)
}
