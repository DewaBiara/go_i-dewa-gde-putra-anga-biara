package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"1_JWT_Auth/models"
)

const (
	JWT_SECRET = "ThisIsSecret"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_USERNAME": "root",
		"DB_PASSWORD": "root",
		"DB_PORT":     "3306",
		"DB_HOST":     "mysql-db-1",
		"DB_NAME":     "layered_api",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
