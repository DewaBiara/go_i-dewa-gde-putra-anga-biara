package database

import (
	"03_layered_api_blog/config"
	"03_layered_api_blog/models"
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
	err := config.DB.Create(&user).Error

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
	result := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, ErrInvalidID
	}

	return user, nil
}
