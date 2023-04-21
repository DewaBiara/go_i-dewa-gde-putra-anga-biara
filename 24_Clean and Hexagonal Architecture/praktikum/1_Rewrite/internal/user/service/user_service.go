package service

import (
	"rewrite/internal/user/dto"
)

type UserService interface {
	FindAll() (dto.Users, error)
	CreateUser(user dto.User) error
}
