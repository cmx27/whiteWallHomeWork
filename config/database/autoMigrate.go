package database

import (
	"whiteWall/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Manager{},
		&models.Student{},
		&models.Artical{},
		&models.Comment{},
	)

	return err
}
