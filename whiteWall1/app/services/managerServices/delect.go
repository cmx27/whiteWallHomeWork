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
	result := database.DB.Where(&models.User{
		UserID: user_id,
	}).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	result2 := database.DB.Where(&models.Artical{
		UserID: user_id,
	}).Delete(&models.Artical{})
	//models.Num--
	return result2.Error

}
