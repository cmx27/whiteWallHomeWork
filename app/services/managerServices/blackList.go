package managerServices

import (
	//"time"

	"whiteWall/app/models"
	"whiteWall/config/database"
	//"gorm.io/gorm"
)

func BlackDeleteUserByUserID(user_id uint) error {
	result := database.DB.Where("user_id = ?", user_id).Where("manager_state <> ?", true).Delete(&models.User{})
	return result.Error

	//time.Sleep(time.Duration(t) * time.Minute)

}

func BlackPutUserByUserID(user_id uint) error {
	result := database.DB.Where("user_id = ?", user_id).Where("manager_state <> ?", true).Update("deleted_at", nil)
	return result.Error
}
