package repository

import (
	"rewrite/pkg/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) FindAll() (entity.Users, error) {
	var users entity.Users

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) CreateUser(user *entity.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
