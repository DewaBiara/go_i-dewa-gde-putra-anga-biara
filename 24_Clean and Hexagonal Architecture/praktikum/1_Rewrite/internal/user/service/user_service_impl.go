package service

import (
	"rewrite/internal/user/dto"
	"rewrite/internal/user/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (u *UserServiceImpl) FindAll() (dto.Users, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var dtoUsers dto.Users
	dtoUsers.FromEntity(users)
	return dtoUsers, nil
}

func (u *UserServiceImpl) CreateUser(user dto.User) error {
	userEntity := user.ToEntity()
	return u.userRepository.CreateUser(userEntity)
}
