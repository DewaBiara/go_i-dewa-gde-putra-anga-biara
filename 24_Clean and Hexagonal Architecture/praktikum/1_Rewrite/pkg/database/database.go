package database

import (
	"rewrite/pkg/config"
	"rewrite/pkg/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.DB), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		entity.User{},
	)
}
