package database

import (
	"whiteWall/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Artical{},
		&models.Manager{},
	)

	return err
}
