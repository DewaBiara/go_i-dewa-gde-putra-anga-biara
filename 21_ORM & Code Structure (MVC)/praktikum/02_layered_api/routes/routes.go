package routes

import (
	"layered_api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	userGroup := e.Group("/users")
	NewUserControllers(userGroup)

	bookGroup := e.Group("/books")
	NewBookControllers(bookGroup)

	return e
}

func NewUserControllers(group *echo.Group) {
	group.GET("", controllers.GetUsersController)
	group.GET("/:id", controllers.GetUserController)
	group.POST("", controllers.CreateUserController)
	group.DELETE("/:id", controllers.DeleteUserController)
	group.PUT("/:id", controllers.UpdateUserController)
}

func NewBookControllers(group *echo.Group) {
	group.GET("", controllers.GetBooksController)
	group.GET("/:id", controllers.GetBookController)
	group.POST("", controllers.CreateBookController)
	group.DELETE("/:id", controllers.DeleteBookController)
	group.PUT("/:id", controllers.UpdateBookController)
}
