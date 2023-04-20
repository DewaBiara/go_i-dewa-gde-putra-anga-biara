package main

import (
	"github.com/labstack/echo/v4"
	"github.com/suryaadi44/go-training-restful/config"
	"github.com/suryaadi44/go-training-restful/lib/database"
	"github.com/suryaadi44/go-training-restful/routes"
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
