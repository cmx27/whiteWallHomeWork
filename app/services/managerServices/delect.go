package managerServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func DeleteArticalByArticalID(artical_id uint) error {
	result := database.DB.Delete(&models.Artical{}, artical_id)
	return result.Error
}

func DeleteUserByUserID(user_id uint) error {
	result := database.DB.Unscoped().Where("user_id = ?", user_id).Where("manager_state <> ?", true).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	result3 := database.DB.Where(&models.Student{
		StudentID: user_id,
	}).Delete(&models.Student{})
	if result3.Error != nil {
		return result3.Error
	}
	result2 := database.DB.Where(&models.Artical{
		UserID: user_id,
	}).Delete(&models.Artical{})

	return result2.Error

}
