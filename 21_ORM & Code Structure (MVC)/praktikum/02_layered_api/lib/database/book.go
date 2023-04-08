package database

import (
	"layered_api/config"
	"layered_api/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	err := config.DB.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func GetBookByID(id int) (interface{}, error) {
	var book models.Book

	err := config.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	err := config.DB.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func DeleteBook(id int) error {
	var book models.Book

	result := config.DB.Where("id = ?", id).Delete(&book)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInvalidID
	}

	return nil
}

func UpdateBook(id int, book models.Book) (interface{}, error) {
	result := config.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		return book, result.Error
	}

	if result.RowsAffected == 0 {
		return book, ErrInvalidID
	}

	return book, nil
}
