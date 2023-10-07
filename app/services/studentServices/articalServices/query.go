package articalServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func GetStusentByUserID(id uint) (*models.Student, error) {
	var student models.Student
	result := database.DB.Where(&models.Student{
		StudentID: id,
	}).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

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

func GetArticalByArticalID(artical_id uint) (models.Artical, error) {
	var artical models.Artical
	result := database.DB.Where(&models.Artical{
		ArticalID: artical_id,
	}).Find(&artical)
	if result.Error != nil {
		return artical, result.Error
	}
	if artical.Namestate {
		artical.Name = "匿名"
	}
	return artical, nil
}
