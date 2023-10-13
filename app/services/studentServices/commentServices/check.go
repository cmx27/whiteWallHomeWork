package commentServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func CheckArtical(artical_id uint) error {
	result := database.DB.Where("id = ?", artical_id).First(&models.Artical{})
	return result.Error
}
