package controllers

import (
	"net/http"
	"strconv"

	"RESTFUL_API/lib/database"
	"RESTFUL_API/lib/utils"
	"RESTFUL_API/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	UserService database.UserService
}

func NewUserController(userService database.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (u *UserController) GetUsersController(c echo.Context) error {
	users, err := u.UserService.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting users",
		"data":    users,
	})
}

func (u *UserController) GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	user, err := u.UserService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting user",
		"data":    user,
	})
}

func (u *UserController) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	newUser, err := u.UserService.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success creating user",
		"data":    newUser,
	})
}

func (u *UserController) DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = u.UserService.DeleteUser(id)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting user",
	})
}

func (u *UserController) UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	user := models.User{}
	c.Bind(&user)

	if user.Password != "" {
		hash, err := utils.HashPassword(user.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		user.Password = hash
	}

	err = u.UserService.UpdateUser(id, user)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating user",
	})
}

func (u *UserController) LoginController(c echo.Context) error {
	userRequest := models.User{}
	c.Bind(&userRequest)

	user, err := u.UserService.LoginUser(userRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidCredentials.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !utils.ComparePassword(userRequest.Password, user.Password) {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidCredentials.Error())
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}
