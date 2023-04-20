package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/suryaadi44/go-training-restful/lib/database"
	"github.com/suryaadi44/go-training-restful/models"
	"gorm.io/gorm"
)

type BookController struct {
	BookService database.BookService
}

func NewBookController(BookService database.BookService) *BookController {
	return &BookController{
		BookService: BookService,
	}
}

func (b *BookController) GetBooksController(c echo.Context) error {
	books, err := b.BookService.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting books",
		"data":    books,
	})
}

func (b *BookController) GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	book, err := b.BookService.GetBookByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting book",
		"data":    book,
	})
}

func (b *BookController) CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	newBook, err := b.BookService.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success creating book",
		"data":    newBook,
	})
}

func (b *BookController) DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = b.BookService.DeleteBook(id)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting book",
	})
}

func (b *BookController) UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	book := models.Book{}
	c.Bind(&book)

	err = b.BookService.UpdateBook(id, book)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating book",
	})
}
