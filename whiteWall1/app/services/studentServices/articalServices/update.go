package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func UpdateArtical(artical string, artical_id uint) error {
	result := database.DB.Model(&models.Artical{
		ArticalID: artical_id,
	}).Updates(models.Artical{
		Artical: artical})
	return result.Error
}

func UpdateImage(artical_id uint, imageAddr string) error {

	result := database.DB.Where(&models.Artical{
		ArticalID: artical_id,
	}).Updates(models.Artical{
		ImageAddr: imageAddr,
	})
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}

}
