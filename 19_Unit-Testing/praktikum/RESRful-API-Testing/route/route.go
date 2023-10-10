package route

import (
	"Laode_Saady/19_Unit-Testing/praktikum/RESRful-API-Testing/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	UserProtect := e.Group("/users")

	UserProtect.GET("", controllers.GetUsersController)
	UserProtect.GET("/:id", controllers.GetUserController)
	UserProtect.DELETE("/:id", controllers.DeleteUserController)
	UserProtect.PUT("/:id", controllers.UpdateUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	BooksProtect := e.Group("/books")

	BooksProtect.GET("", controllers.GetAllBooksController)
	BooksProtect.GET("/:id", controllers.GetBookController)
	BooksProtect.PUT("/:id", controllers.UpdateBookController)
	BooksProtect.DELETE("/:id", controllers.DeleteBookController)
	BooksProtect.POST("", controllers.CreateBookController)
}
