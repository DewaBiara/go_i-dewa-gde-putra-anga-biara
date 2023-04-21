package repository

import "rewrite/pkg/entity"

type UserRepository interface {
	FindAll() (entity.Users, error)
	CreateUser(user *entity.User) error
}
