package database

import (
	"RESTFUL_API/models"

	"gorm.io/gorm"
)

type BookServiceImpl struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookServiceImpl {
	return &BookServiceImpl{
		db: db,
	}
}

func (b *BookServiceImpl) GetBooks() (interface{}, error) {
	var books []models.Book

	err := b.db.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (b *BookServiceImpl) GetBookByID(id int) (interface{}, error) {
	var book models.Book

	err := b.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *BookServiceImpl) CreateBook(book models.Book) (interface{}, error) {
	err := b.db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *BookServiceImpl) DeleteBook(id int) error {
	var book models.Book

	result := b.db.Where("id = ?", id).Delete(&book)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func (b *BookServiceImpl) UpdateBook(id int, book models.Book) error {
	result := b.db.Model(&models.Book{}).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
