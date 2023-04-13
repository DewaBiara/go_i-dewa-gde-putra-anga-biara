package controllers

import (
	"net/http"
	"strconv"

	"1_JWT_Auth/lib/database"
	"1_JWT_Auth/lib/utils"
	"1_JWT_Auth/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting users",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	user, err := database.GetUserByID(id)
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

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	newUser, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating user",
		"data":    newUser,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = database.DeleteUser(id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting user",
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	user := models.User{}
	c.Bind(&user)

	updatedUser, err := database.UpdateUser(id, user)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating user",
		"data":    updatedUser,
	})
}

func LoginController(c echo.Context) error {
	userRequest := models.User{}
	c.Bind(&userRequest)

	user, err := database.LoginUser(userRequest)
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
