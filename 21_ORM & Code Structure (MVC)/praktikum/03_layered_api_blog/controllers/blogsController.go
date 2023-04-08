package controllers

import (
	"net/http"
	"strconv"

	"03_layered_api_blog/lib/database"
	"03_layered_api_blog/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBlogsController(c echo.Context) error {
	blogs, err := database.GetBlogs()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting blogs",
		"data":    blogs,
	})
}

func GetBlogsIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	blogs, err := database.GetBlogsByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting blogs",
		"data":    blogs,
	})
}

func CreateBlogsController(c echo.Context) error {
	blogs := models.Blogs{}
	c.Bind(&blogs)

	newBlogs, err := database.CreateBlogs(blogs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating blogs",
		"data":    newBlogs,
	})
}

func DeleteBlogsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteBlogs(id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting blogs",
	})
}

func UpdateBlogsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	blogs := models.Blogs{}
	c.Bind(&blogs)

	updatedBlogs, err := database.UpdateBlogs(id, blogs)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating blogs",
		"data":    updatedBlogs,
	})
}
