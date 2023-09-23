package route

import (
	"Laode_Saady/17_ORM-Code-Structure-MVC/praktikum/controller"

	"github.com/labstack/echo/v4"
)

// Route / to handler function
func New() *echo.Echo {
	e := echo.New()

	user := e.Group("/users")
	user.GET("", controller.GetUsersController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.DELETE("/:id", controller.DeleteUserController)
	user.PUT("/:id", controller.UpdateUserController)

	buku := e.Group("/buku")
	buku.GET("", controller.GetBukuController)
	buku.POST("", controller.CreateBukuController)
	buku.GET("/:id", controller.GetBukuIdController)
	buku.DELETE("/:id", controller.DeletedBukuController)
	buku.PUT("/:id", controller.UpdatedBukuController)

	blog := e.Group("/blogs")
	blog.GET("", controller.GetBlogsController)
	blog.POST("", controller.CreateBlogController)
	blog.GET("/:id", controller.GetByIdController)
	blog.PUT("/:id", controller.UpdateBlogController)
	blog.DELETE("/:id", controller.DeleteBlogController)
	return e
}
