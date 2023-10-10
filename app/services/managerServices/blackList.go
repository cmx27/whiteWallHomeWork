package managerServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func BlackDeleteUserByUserID(user_id uint) error {
	result := database.DB.Where("user_id = ?", user_id).Where("manager_state <> ?", true).Updates(models.User{
		BlackState: true,
	})
	return result.Error

}

func BlackPutUserByUserID(user_id uint) error {
	user := &models.User{}
	result := database.DB.Where("user_id = ?", user_id).Where("manager_state <> ?", true).First(&user).Updates(map[string]interface{}{
		"black_state": false,
	})
	return result.Error

}
