package database

import (
	"RESTFUL_API/models"
)

type BookService interface {
	GetBooks() (interface{}, error)
	GetBookByID(id int) (interface{}, error)
	CreateBook(book models.Book) (interface{}, error)
	DeleteBook(id int) error
	UpdateBook(id int, book models.Book) error
}

type UserService interface {
	GetUsers() (interface{}, error)
	GetUserByID(id int) (interface{}, error)
	CreateUser(user models.User) (interface{}, error)
	DeleteUser(id int) error
	UpdateUser(id int, user models.User) error
	LoginUser(requestUser models.User) (models.User, error)
}
