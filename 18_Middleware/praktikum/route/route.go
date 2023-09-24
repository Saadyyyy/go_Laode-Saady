package route

import (
	"Laode_Saady/18_Middleware/praktikum/controller"
	"Laode_Saady/18_Middleware/praktikum/middleware"

	"github.com/labstack/echo/v4"
)

// Route / to handler function
func New() *echo.Echo {
	e := echo.New()
	middleware.LogMiddleware(e)

	user := e.Group("/users")
	user.GET("", controller.GetUsersController, middleware.JwtMiddleware())
	user.GET("/:id", controller.GetUserController, middleware.JwtMiddleware())
	user.POST("", controller.CreateUserController)
	user.DELETE("/:id", controller.DeleteUserController, middleware.JwtMiddleware())
	user.PUT("/:id", controller.UpdateUserController, middleware.JwtMiddleware())
	user.POST("/login", controller.LoginUserController)

	buku := e.Group("/buku")
	buku.GET("", controller.GetBukuController, middleware.JwtMiddleware())
	buku.POST("", controller.CreateBukuController, middleware.JwtMiddleware())
	buku.GET("/:id", controller.GetBukuIdController, middleware.JwtMiddleware())
	buku.DELETE("/:id", controller.DeletedBukuController, middleware.JwtMiddleware())
	buku.PUT("/:id", controller.UpdatedBukuController, middleware.JwtMiddleware())

	blog := e.Group("/blogs")
	blog.GET("", controller.GetBlogsController)
	blog.POST("", controller.CreateBlogController)
	blog.GET("/:id", controller.GetByIdController)
	blog.PUT("/:id", controller.UpdateBlogController)
	blog.DELETE("/:id", controller.DeleteBlogController)
	return e
}
