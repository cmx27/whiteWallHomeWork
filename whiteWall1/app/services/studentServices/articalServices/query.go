package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

// func GetUserByUserID(id uint) (*models.User, error) {
// 	var user models.User
// 	result := database.DB.Where(&models.User{
// 		UserID: id,
// 	}).First(&user)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }

func GetArticalByUserIDAndArticalID(user_id uint, artical_id uint) (*models.Artical, error) {
	var artical models.Artical
	result := database.DB.Where(&models.Artical{
		UserID:    user_id,
		ArticalID: artical_id,
	}).First(&artical)
	if result.Error != nil {
		return nil, result.Error
	}
	return &artical, nil
}

func GetArticalByUserID(user_id uint) ([]models.Artical, error) {
	var artical_list []models.Artical
	result := database.DB.Where(&models.Artical{
		UserID: user_id,
	}).Find(&artical_list)
	if result.Error != nil {
		return nil, result.Error
	}
	return artical_list, nil
}
