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

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUsers() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func (m *MockUserService) GetUserByID(id int) (interface{}, error) {
	args := m.Called(id)
	return args.Get(0), args.Error(1)
}

func (m *MockUserService) CreateUser(user models.User) (interface{}, error) {
	args := m.Called(user)
	return args.Get(0), args.Error(1)
}

func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) UpdateUser(id int, user models.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func (m *MockUserService) LoginUser(requestUser models.User) (models.User, error) {
	args := m.Called(requestUser)
	return args.Get(0).(models.User), args.Error(1)
}

func TestNewUserController(t *testing.T) {
	mockUserService := new(MockUserService)

	userController := NewUserController(mockUserService)

	assert.NotNil(t, userController)
	assert.NotNil(t, userController.UserService)
}

func TestUserController_GetUsersController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedBody interface{}
		ExpectedCode int
		ExpectedErr  error
	}{
		{
			Name: "Get users success",
			Args: map[string]interface{}{
				"method": "GET",
				"users": []models.User{
					{
						Name:     "Surya",
						Email:    "123@123.com",
						Password: "123",
					},
					{
						Name:     "Adi",
						Email:    "456@456.com",
						Password: "456",
					},
				},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success getting users",
				"data": []interface{}{
					map[string]interface{}{
						"CreatedAt": "0001-01-01T00:00:00Z",
						"UpdatedAt": "0001-01-01T00:00:00Z",
						"DeletedAt": nil,
						"ID":        float64(0),
						"name":      "Surya",
						"email":     "123@123.com",
						"password":  "123",
					},
					map[string]interface{}{
						"CreatedAt": "0001-01-01T00:00:00Z",
						"UpdatedAt": "0001-01-01T00:00:00Z",
						"DeletedAt": nil,
						"ID":        float64(0),
						"name":      "Adi",
						"email":     "456@456.com",
						"password":  "456",
					},
				},
			},
		},
		{
			Name: "Get users with empty users",
			Args: map[string]interface{}{
				"method": "GET",
				"users":  []models.User{},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success getting users",
				"data":    []interface{}{},
			},
		},
		{
			Name: "Get users with error",
			Args: map[string]interface{}{
				"method": "GET",
				"users":  []models.User{},
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error getting users"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(tc.Args["method"].(string), "/", nil)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)

			mockUserService := new(MockUserService)
			mockUserService.On("GetUsers").Return(tc.Args["users"], tc.ExpectedErr)

			userController := NewUserController(mockUserService)

			// Execution
			err := userController.GetUsersController(c)

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

func TestUserController_GetUserController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedBody interface{}
		ExpectedCode int
		ExpectedErr  error
	}{
		{
			Name: "Get user success",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "1",
				"user": models.User{
					Name:     "Surya",
					Email:    "123@123.com",
					Password: "123",
				},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success getting user",
				"data": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": nil,
					"ID":        float64(0),
					"name":      "Surya",
					"email":     "123@123.com",
					"password":  "123",
				},
			},
		},
		{
			Name: "Get user with id 1 not found",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "1",
				"user":   models.User{},
			},
			ExpectedCode: 404,
			ExpectedErr:  gorm.ErrRecordNotFound,
		},
		{
			Name: "Get user with error",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "1",
				"user":   models.User{},
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error getting user"),
		},
		{
			Name: "Get user with invalid id",
			Args: map[string]interface{}{
				"method": "GET",
				"id":     "invalid",
				"user":   models.User{},
			},
			ExpectedCode: 500,
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
			mockUserService := new(MockUserService)
			mockUserService.On("GetUserByID", args).Return(tc.Args["user"], tc.ExpectedErr)

			userController := NewUserController(mockUserService)

			// Execution
			err := userController.GetUserController(c)

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

func TestUserController_CreateUserController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedBody interface{}
		ExpectedCode int
		ExpectedErr  error
	}{
		{
			Name: "Create user success",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.User{
					Name:     "Surya",
					Email:    "123@123.com",
					Password: "123",
				},
				"mime": "application/json",
			},
			ExpectedCode: 201,
			ExpectedBody: map[string]interface{}{
				"message": "success creating user",
				"data": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": nil,
					"ID":        float64(0),
					"name":      "Surya",
					"email":     "123@123.com",
					"password":  "123",
				},
			},
		},
		{
			Name: "Create user with error",
			Args: map[string]interface{}{
				"method": "POST",
				"body":   models.User{},
				"mime":   "application/json",
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error creating user"),
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

			mockUserService := new(MockUserService)
			mockUserService.On("CreateUser", tc.Args["body"].(models.User)).Return(tc.Args["body"], tc.ExpectedErr)

			userController := NewUserController(mockUserService)

			// Execution
			err := userController.CreateUserController(c)

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

func TestUserController_DeleteUserController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedBody interface{}
		ExpectedCode int
		ExpectedErr  error
	}{
		{
			Name: "Delete user with id 1 success",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "1",
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success deleting user",
			},
		},
		{
			Name: "Delete user with id 1 not found",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "1",
			},
			ExpectedCode: 404,
			ExpectedErr:  database.ErrIDNotFound,
		},
		{
			Name: "Delete user with general error",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "1",
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error deleting user"),
		},
		{
			Name: "Delete user with invalid id",
			Args: map[string]interface{}{
				"method": "DELETE",
				"id":     "a",
			},
			ExpectedCode: 400,
			ExpectedErr:  ErrInvalidID,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			r := httptest.NewRequest(tc.Args["method"].(string), "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(r, w)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.Args["id"].(string))

			args, _ := strconv.Atoi(tc.Args["id"].(string))
			mockUserService := new(MockUserService)
			mockUserService.On("DeleteUser", args).Return(tc.ExpectedErr)

			userController := NewUserController(mockUserService)

			// Execution
			err := userController.DeleteUserController(c)

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

func TestUserController_UpdateUserController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedBody interface{}
		ExpectedCode int
		ExpectedErr  error
	}{
		{
			Name: "Update user with id 1 success without password",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.User{
					Name: "Surya",
				},
				"mime": "application/json",
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success updating user",
			},
		},
		{
			Name: "Update user with id 1 success with password",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.User{
					Name:     "Surya",
					Password: "123456",
				},
				"mime": "application/json",
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success updating user",
			},
		},
		{
			Name: "Update user with id 1 not found",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.User{
					Name: "Surya",
				},
				"mime": "application/json",
			},
			ExpectedCode: 404,
			ExpectedErr:  database.ErrIDNotFound,
		},
		{
			Name: "Update user with general error",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "1",
				"body": models.User{
					Name: "Surya",
				},
				"mime": "application/json",
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error updating user"),
		},
		{
			Name: "Update user with invalid id",
			Args: map[string]interface{}{
				"method": "PUT",
				"id":     "a",
				"body": models.User{
					Name: "Surya",
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
			mockUserService := new(MockUserService)
			if tc.Args["body"].(models.User).Password == "" {
				mockUserService.On("UpdateUser", args, tc.Args["body"].(models.User)).Return(tc.ExpectedErr)
			} else {
				mockUserService.On("UpdateUser", args, mock.AnythingOfType("models.User")).Return(tc.ExpectedErr)
			}

			userController := NewUserController(mockUserService)

			// Execution
			err := userController.UpdateUserController(c)

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

func TestUserController_LoginController(t *testing.T) {
	testCase := []struct {
		Name         string
		Args         map[string]interface{}
		ExpectedBody interface{}
		ExpectedCode int
		ExpectedErr  error
	}{
		{
			Name: "Login success",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.User{
					Email:    "123@123.com",
					Password: "456",
				},
				"mime": "application/json",
				"saved": models.User{
					Name:     "Surya",
					Email:    "123@123.com",
					Password: "$2a$10$i8ga2m4ii.adbP66lG3NL.pJEgZ7lD2D/I2cEwlvcI5NDB8CROMF2", // Hash of 456
				},
			},
			ExpectedCode: 200,
			ExpectedBody: map[string]interface{}{
				"message": "success login",
				"token":   "Some Token",
			},
		},
		{
			Name: "Login failed with invalid credentials",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.User{
					Email:    "123@123.com",
					Password: "123",
				},
				"mime": "application/json",
				"saved": models.User{
					Name:     "Surya",
					Email:    "123@123.com",
					Password: "$2a$10$i8ga2m4ii.adbP66lG3NL.pJEgZ7lD2D/I2cEwlvcI5NDB8CROMF2", // Hash of 456
				},
			},
			ExpectedCode: 401,
			ExpectedErr:  ErrInvalidCredentials,
		},
		{
			Name: "Login failed with general error",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.User{
					Email:    "123@123.com",
					Password: "123",
				},
				"mime": "application/json",
				"saved": models.User{
					Name:     "Surya",
					Email:    "123@123.com",
					Password: "$2a$10$i8ga2m4ii.adbP66lG3NL.pJEgZ7lD2D/I2cEwlvcI5NDB8CROMF2", // Hash of 456
				},
				"ThrownErr": fmt.Errorf("error login"),
			},
			ExpectedCode: 500,
			ExpectedErr:  fmt.Errorf("error login"),
		},
		{
			Name: "Login failed for user not exist",
			Args: map[string]interface{}{
				"method": "POST",
				"body": models.User{
					Email:    "123@123.com",
					Password: "123",
				},
				"mime":      "application/json",
				"saved":     models.User{},
				"ThrownErr": gorm.ErrRecordNotFound,
			},
			ExpectedCode: 401,
			ExpectedErr:  ErrInvalidCredentials,
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

			mockUserService := new(MockUserService)
			mockUserService.On("LoginUser", tc.Args["body"].(models.User)).Return(tc.Args["saved"].(models.User), tc.Args["ThrownErr"])

			userController := NewUserController(mockUserService)

			// Execution
			err := userController.LoginController(c)

			// Validation
			if tc.ExpectedErr != nil {
				assert.Error(t, echo.NewHTTPError(tc.ExpectedCode, tc.ExpectedErr.Error()))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, w.Result().StatusCode)

				var response map[string]interface{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tc.ExpectedBody.(map[string]interface{})["message"], response["message"])
			}
		})
	}
}
