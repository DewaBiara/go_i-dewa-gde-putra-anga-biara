package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email    string
	Password string
}

type Users []User
