package userServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func GetUserByAccount(account string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(&models.User{
		Account: account,
	}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
