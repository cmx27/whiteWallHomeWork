package commentServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func CreateComment(name, comment string, user_id, artical_id uint) error {
	content := &models.Comment{
		Name:      name,
		Comment:   comment,
		UserID:    user_id,
		ArticalID: artical_id,
	}
	result := database.DB.Create(content)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
