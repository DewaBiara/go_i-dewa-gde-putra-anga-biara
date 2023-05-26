package entity

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          string `gorm:"primaryKey; type:varchar(36)"`
	Name        string `gorm:"type:varchar(255);not null;uniqueIndex"`
	CategoryID  int
	Price       int
	Stock       int
	Description string `gorm:"type:varchar(255)"`
	CreatedBy   string `gorm:"type:varchar(255)"`
	UpdatedBy   string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Items []Item
