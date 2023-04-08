package models

import (
	"gorm.io/gorm"
)

type Blogs struct {
	gorm.Model
	UserID  int    `json:"user_id" form:"user_id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
