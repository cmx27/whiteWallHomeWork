package managerServices

import (
	//"fmt"
	//"log"
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func GetUser() ([]models.User, error) {
	var user_list []models.User
	result := database.DB.Where(&models.User{}).Find(&user_list)
	if result.Error != nil {
		return nil, result.Error
	}
	return user_list, nil
}

func GetManagerByID(id uint) (*models.Manager, error) {
	manager := models.Manager{}
	result := database.DB.Where(&models.Manager{
		ManagerID: id,
	}).First(&manager)

	if result.Error != nil {
		return nil, result.Error
	}
	return &manager, nil
}

func GetArtical() ([]models.Artical, error) {
	var artical_list []models.Artical
	result := database.DB.Find(&artical_list)
	if result.Error != nil {
		return nil, result.Error
	}
	return artical_list, nil
}
