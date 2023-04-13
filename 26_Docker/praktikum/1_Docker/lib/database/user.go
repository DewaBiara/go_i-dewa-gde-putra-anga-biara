package database

import (
	"1_JWT_Auth/config"
	"1_JWT_Auth/lib/utils"
	"1_JWT_Auth/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByID(id int) (interface{}, error) {
	var user models.User

	err := config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hash

	err = config.DB.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id int) error {
	var user models.User

	result := config.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInvalidID
	}

	return nil
}

func UpdateUser(id int, user models.User) (interface{}, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hash

	result := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, ErrInvalidID
	}

	return user, nil
}

func LoginUser(requestUser models.User) (models.User, error) {
	var user models.User

	err := config.DB.Where("email = ?", requestUser.Email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
