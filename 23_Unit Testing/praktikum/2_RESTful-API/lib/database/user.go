package database

import (
	"RESTFUL_API/lib/utils"
	"RESTFUL_API/models"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}

func (u *UserServiceImpl) GetUsers() (interface{}, error) {
	var users []models.User

	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserServiceImpl) GetUserByID(id int) (interface{}, error) {
	var user models.User

	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserServiceImpl) CreateUser(user models.User) (interface{}, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hash

	err = u.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserServiceImpl) DeleteUser(id int) error {
	var user models.User

	result := u.db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func (u *UserServiceImpl) UpdateUser(id int, user models.User) error {
	result := u.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func (u *UserServiceImpl) LoginUser(requestUser models.User) (models.User, error) {
	var user models.User

	err := u.db.Where("email = ?", requestUser.Email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
