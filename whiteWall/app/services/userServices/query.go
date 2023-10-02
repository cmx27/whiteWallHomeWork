package userServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func GetUserByAccount(account string) (interface{}, error) {
	var user models.User
	result := database.DB.Where(&models.User{
		Account: account,
	}).First(&user)

	if result.Error != nil {
		var manager models.Manager
		result2 := database.DB.Where(models.Manager{
			Account: account,
		}).First(&manager)
		if result2.Error != nil {
			return nil, result2.Error
		}
		return manager, nil
	}
	return user, nil
}
