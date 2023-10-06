package managerServices

import (
	"time"
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func BlackDeleteUserByUserID(user_id uint, t int) error {
	result := database.DB.Where(&models.User{
		UserID: user_id,
	}).Delete(&models.User{})
	if result == nil {
		return result.Error
	}
	time.Sleep(time.Duration(t) * time.Minute)
	result2 := database.DB.Model(&models.User{}).Where(&models.User{
		UserID: user_id,
	}).Update("deleted_at", nil)

	return result2.Error
}
