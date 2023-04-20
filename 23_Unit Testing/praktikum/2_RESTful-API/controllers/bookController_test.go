package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/suryaadi44/go-training-restful/lib/database"
	"github.com/suryaadi44/go-training-restful/models"
	"gorm.io/gorm"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) GetBooks() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func (m *MockBookService) GetBookByID(id int) (interface{}, error) {
	args := m.Called(id)
	return args.Get(0), args.Error(1)
}

func (m *MockBookService) CreateBook(book models.Book) (interface{}, error) {
	args := m.Called(book)
	return args.Get(0), args.Error(1)
}

func (m *MockBookService) DeleteBook(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookService) UpdateBook(id int, book models.Book) error {
	args := m.Called(id, book)
	return args.Error(0)
}

func TestNewBookController(t *testing.T) {
	mockBookService := new(MockBookService)

	bookController := NewBookController(mockBookService)

	assert.NotNil(t, bookController)
	assert.NotNil(t, bookController.BookService)
}

func TestBookController_GetBooksController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedCode int
		ExpectedBody interface{}
		ExpectedErr  error
	}{
		{
			Name: "Get books success",
			Args: map[string]interface{}{
				"method": "GET",
				"books": []models.Book{
					{
						Title:  "The Hobbit",
						Author: "J.R.R. Tolkien",
						Genres: "Fantasy",
						Year:   1937,
					},
					{
						Title:  "The Lord of the Rings",
						Author: "J.R.R. Tolkien",
						Genres: "Fantasy",
						Year:   1954,
					},
				},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success getting books",
				"data": []interface{}{
					map[string]interface{}{
						"CreatedAt": "0001-01-01T00:00:00Z",
						"UpdatedAt": "0001-01-01T00:00:00Z",
						"DeletedAt": nil,
						"ID":        float64(0),
						"title":     "The Hobbit",
						"author":    "J.R.R. Tolkien",
						"genre":     "Fantasy",
						"year":      float64(1937),
					},
					map[string]interface{}{
						"CreatedAt": "0001-01-01T00:00:00Z",
						"UpdatedAt": "0001-01-01T00:00:00Z",
						"DeletedAt": nil,
						"ID":        float64(0),
						"title":     "The Lord of the Rings",
						"author":    "J.R.R. Tolkien",
						"genre":     "Fantasy",
						"year":      float64(1954),
					},
				},
			},
		},
		{
			Name: "Get books with empty data",
			Args: map[string]interface{}{
				"method": "GET",
				"books":  []models.Book{},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success getting books",
				"data":    []interface{}{},
			},
		},
		{
			Name: "Get books with error",
			Args: map[string]interface{}{
				"method": "GET",
				"books":  nil,
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error getting books"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(tc.Args["method"].(string), "/", nil)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)

			mockBookService := new(MockBookService)
			mockBookService.On("GetBooks").Return(tc.Args["books"], tc.ExpectedErr)

			bookController := BookController{
				BookService: mockBookService,
			}

			// Execution
			err := bookController.GetBooksController(c)

			// Validation
			if tc.ExpectedErr != nil {
				assert.Error(t, echo.NewHTTPError(tc.ExpectedCode, tc.ExpectedErr.Error()))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, w.Result().StatusCode)

				var response map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tc.ExpectedBody, response)
			}
		})
	}
}

func TestBookController_GetBookController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedCode int
		ExpectedBody interface{}
		ExpectedErr  error
	}{
		{
			Name: "Get book with id 1 success",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "1",
				"book": models.Book{
					Title:  "The Hobbit",
					Author: "J.R.R. Tolkien",
					Genres: "Fantasy",
					Year:   1937,
				},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success getting book",
				"data": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": nil,
					"ID":        float64(0),
					"title":     "The Hobbit",
					"author":    "J.R.R. Tolkien",
					"genre":     "Fantasy",
					"year":      float64(1937),
				},
			},
		},
		{
			Name: "Get book with id 1 not found",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "1",
				"book":   models.Book{},
			},
			ExpectedCode: 404,
			ExpectedErr:  gorm.ErrRecordNotFound,
		},
		{
			Name: "Get books with error",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "1",
				"book":   models.Book{},
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error getting books"),
		},
		{
			Name: "Get books with invalid id",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "string",
				"book":   models.Book{},
			},
			ExpectedCode: 400,
			ExpectedErr:  ErrInvalidID,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(tc.Args["method"].(string), "/", nil)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.Args["id"].(string))

			args, _ := strconv.Atoi(tc.Args["id"].(string))
			mockBookService := new(MockBookService)
			mockBookService.On("GetBookByID", args).Return(tc.Args["book"], tc.ExpectedErr)

			bookController := BookController{
				BookService: mockBookService,
			}

			// Execution
			err := bookController.GetBookController(c)

			// Validation
			if tc.ExpectedErr != nil {
				assert.Error(t, echo.NewHTTPError(tc.ExpectedCode, tc.ExpectedErr.Error()))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, w.Result().StatusCode)

				var response map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tc.ExpectedBody, response)
			}
		})
	}
}

func TestBookController_CreateBookController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedCode int
		ExpectedBody interface{}
		ExpectedErr  error
	}{
		{
			Name: "Create book success",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.Book{
					Title:  "The Hobbit",
					Author: "J.R.R. Tolkien",
					Genres: "Fantasy",
					Year:   1937,
				},
				"mime": "application/json",
			},
			ExpectedCode: 201,
			ExpectedBody: map[string]interface{}{
				"message": "success creating book",
				"data": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": nil,
					"ID":        float64(0),
					"title":     "The Hobbit",
					"author":    "J.R.R. Tolkien",
					"genre":     "Fantasy",
					"year":      float64(1937),
				},
			},
		},
		{
			Name: "Create book with error",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.Book{
					Title:  "The Hobbit",
					Author: "J.R.R. Tolkien",
					Genres: "Fantasy",
					Year:   1937,
				},
				"mime": "application/json",
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error creating book"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.Args["body"])
			r := httptest.NewRequest(tc.Args["method"].(string), "/", bytes.NewBuffer(jsonBody))
			r.Header.Set("Content-Type", tc.Args["mime"].(string))
			w := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(r, w)

			mockBookService := new(MockBookService)
			mockBookService.On("CreateBook", tc.Args["body"].(models.Book)).Return(tc.Args["body"].(models.Book), tc.ExpectedErr)

			bookController := BookController{
				BookService: mockBookService,
			}

			// Execution
			err := bookController.CreateBookController(c)

			// Validation
			if tc.ExpectedErr != nil {
				assert.Error(t, echo.NewHTTPError(tc.ExpectedCode, tc.ExpectedErr.Error()))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, w.Result().StatusCode)

				var response map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tc.ExpectedBody, response)
			}
		})
	}
}

func TestBookController_DeleteBookController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedCode int
		ExpectedBody interface{}
		ExpectedErr  error
	}{
		{
			Name: "Delete book with id 1 success",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "1",
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success deleting book",
			},
		},
		{
			Name: "Get book with id 1 not found",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "1",
			},
			ExpectedCode: 404,
			ExpectedErr:  database.ErrIDNotFound,
		},
		{
			Name: "Delete books with generic error",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "1",
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error deleting books"),
		},
		{
			Name: "Get books with invalid id",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "string",
			},
			ExpectedCode: 400,
			ExpectedErr:  ErrInvalidID,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(tc.Args["method"].(string), "/", nil)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.Args["id"].(string))

			args, _ := strconv.Atoi(tc.Args["id"].(string))
			mockBookService := new(MockBookService)
			mockBookService.On("DeleteBook", args).Return(tc.ExpectedErr)

			bookController := BookController{
				BookService: mockBookService,
			}

			// Execution
			err := bookController.DeleteBookController(c)

			// Validation
			if tc.ExpectedErr != nil {
				assert.Error(t, echo.NewHTTPError(tc.ExpectedCode, tc.ExpectedErr.Error()))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, w.Result().StatusCode)

				var response map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tc.ExpectedBody, response)
			}
		})
	}
}

func TestBookController_UpdateBookController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedCode int
		ExpectedBody interface{}
		ExpectedErr  error
	}{
		{
			Name: "Update book with id 1 success",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.Book{
					Title: "The Hobbit",
				},
				"mime": "application/json",
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success updating book",
			},
		},
		{
			Name: "Update book with id 1 not found",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.Book{
					Title: "The Hobbit",
				},
				"mime": "application/json",
			},
			ExpectedCode: 404,
			ExpectedErr:  database.ErrIDNotFound,
		},
		{
			Name: "Update book with generic error",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.Book{
					Title: "The Hobbit",
				},
				"mime": "application/json",
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error updating book"),
		},
		{
			Name: "Update book with invalid id",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "string",
				"body": models.Book{
					Title: "The Hobbit",
				},
				"mime": "application/json",
			},
			ExpectedCode: 400,
			ExpectedErr:  ErrInvalidID,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.Args["body"])
			r := httptest.NewRequest(tc.Args["method"].(string), "/", bytes.NewBuffer(jsonBody))
			r.Header.Set("Content-Type", tc.Args["mime"].(string))
			w := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(r, w)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.Args["id"].(string))

			args, _ := strconv.Atoi(tc.Args["id"].(string))
			mockBookService := new(MockBookService)
			mockBookService.On("UpdateBook", args, tc.Args["body"].(models.Book)).Return(tc.ExpectedErr)

			bookController := BookController{
				BookService: mockBookService,
			}

			// Execution
			err := bookController.UpdateBookController(c)

			// Validation
			if tc.ExpectedErr != nil {
				assert.Error(t, echo.NewHTTPError(tc.ExpectedCode, tc.ExpectedErr.Error()))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, w.Result().StatusCode)

				var response map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tc.ExpectedBody, response)
			}
		})
	}
}
