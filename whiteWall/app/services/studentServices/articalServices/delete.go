package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func DeleteArtical(artical_id uint) error {
	result := database.DB.Where(&models.Artical{
		ArticalID: artical_id,
	}).Delete(&models.Artical{})
	//models.Num--
	return result.Error
}
