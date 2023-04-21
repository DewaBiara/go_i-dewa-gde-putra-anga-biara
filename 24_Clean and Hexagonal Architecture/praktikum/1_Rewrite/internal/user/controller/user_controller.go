package controller

import (
	"net/http"
	"rewrite/internal/user/dto"
	"rewrite/internal/user/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) InitRoutes(e *echo.Echo) {
	e.GET("/users", u.GetAllUser)
	e.POST("/users", u.CreateUser)
}

func (u *UserController) GetAllUser(c echo.Context) error {
	users, err := u.userService.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(users) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No user found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "All users",
		"data":    users,
	})
}

func (u *UserController) CreateUser(c echo.Context) error {
	var user dto.User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = u.userService.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User created",
	})
}
