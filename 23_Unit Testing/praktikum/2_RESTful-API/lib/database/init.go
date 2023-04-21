package database

import (
	"fmt"

	"RESTFUL_API/config"
	"RESTFUL_API/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(db)

	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Book{})
}
