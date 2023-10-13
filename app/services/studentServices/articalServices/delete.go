package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func DeleteArtical(artical_id uint) error {
	result4 := database.DB.Where(&models.Comment{
		ArticalID: artical_id,
	}).Delete(&models.Comment{})
	if result4.Error != nil {
		return result4.Error
	}
	result := database.DB.Where(&models.Artical{
		ArticalID: artical_id,
	}).Delete(&models.Artical{})
	return result.Error
}
