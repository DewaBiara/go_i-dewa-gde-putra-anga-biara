package routes

import (
	"RESTFUL_API/config"
	"RESTFUL_API/controllers"
	"RESTFUL_API/lib/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	Echo        *echo.Echo
	BookService database.BookService
	UserService database.UserService
}

func NewRouter(Echo *echo.Echo, bookService database.BookService, userService database.UserService) *Router {
	return &Router{
		Echo:        Echo,
		BookService: bookService,
		UserService: userService,
	}
}

func (r *Router) Init() {
	r.Echo.Pre(middleware.RemoveTrailingSlash())

	r.Echo.Use(middleware.Logger())
	r.Echo.Use(middleware.Recover())

	r.InitUserRouter()
	r.InitBookRouter()
}

func (r *Router) InitUserRouter() {
	userController := controllers.NewUserController(r.UserService)

	secureGroup := r.Echo.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	r.Echo.POST("/users", userController.CreateUserController)
	r.Echo.POST("/login", userController.LoginController)

	// Auth
	secureGroup.GET("/users", userController.GetUsersController)
	secureGroup.GET("/users/:id", userController.GetUserController)
	secureGroup.DELETE("/users/:id", userController.DeleteUserController)
	secureGroup.PUT("/users/:id", userController.UpdateUserController)
}

func (r *Router) InitBookRouter() {
	bookController := controllers.NewBookController(r.BookService)

	secureGroup := r.Echo.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	r.Echo.GET("/books", bookController.GetBooksController)
	r.Echo.GET("/books/:id", bookController.GetBookController)

	// Auth
	secureGroup.POST("/books", bookController.CreateBookController)
	secureGroup.DELETE("/books/:id", bookController.DeleteBookController)
	secureGroup.PUT("/books/:id", bookController.UpdateBookController)
}
