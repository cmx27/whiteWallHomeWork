package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func CreateArtical(name, artical string, name_state bool, user_id uint) error {
	content := &models.Artical{
		Name:      name,
		Namestate: name_state,
		Artical:   artical,
		UserID:    user_id,
	}
	result := database.DB.Create(content)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
