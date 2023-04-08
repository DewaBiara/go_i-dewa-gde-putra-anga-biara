package routes

import (
	"1_JWT_Auth/config"
	"1_JWT_Auth/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	NewUserControllers(e)
	NewBookControllers(e)

	return e
}

func NewUserControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	// Auth
	secureGroup.GET("/users", controllers.GetUsersController)
	secureGroup.GET("/users/:id", controllers.GetUserController)
	secureGroup.DELETE("/users/:id", controllers.DeleteUserController)
	secureGroup.PUT("/users/:id", controllers.UpdateUserController)
}

func NewBookControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)

	// Auth
	secureGroup.POST("/books", controllers.CreateBookController)
	secureGroup.DELETE("/books/:id", controllers.DeleteBookController)
	secureGroup.PUT("/books/:id", controllers.UpdateBookController)
}
