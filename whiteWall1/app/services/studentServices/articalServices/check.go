package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func CheckArtical(artical_id uint) error {
	result := database.DB.Where("id = ?", artical_id).First(&models.Artical{})
	return result.Error
}

func CompareUser(user1 string, user2 string) bool {
	return user1 == user2
}
