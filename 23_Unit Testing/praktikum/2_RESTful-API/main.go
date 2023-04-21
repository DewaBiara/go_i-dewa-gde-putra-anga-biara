package main

import (
	"RESTFUL_API/config"
	"RESTFUL_API/lib/database"
	"RESTFUL_API/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.NewDB()

	bookService := database.NewBookService(db)
	userService := database.NewUserService(db)

	e := echo.New()
	router := routes.NewRouter(e, bookService, userService)
	router.Init()

	router.Echo.Logger.Fatal(e.Start(config.PORT))
}
