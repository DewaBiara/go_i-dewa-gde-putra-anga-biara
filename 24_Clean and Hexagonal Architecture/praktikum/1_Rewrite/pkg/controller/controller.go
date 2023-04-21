package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userControllerPkg "rewrite/internal/user/controller"
	userRepositoryPkg "rewrite/internal/user/repository"
	userServicePkg "rewrite/internal/user/service"
)

func InitControllers(e *echo.Echo, db *gorm.DB) {
	userRepository := userRepositoryPkg.NewUserRepositoryImpl(db)
	userService := userServicePkg.NewUserServiceImpl(userRepository)
	userController := userControllerPkg.NewUserController(userService)
	userController.InitRoutes(e)
}
