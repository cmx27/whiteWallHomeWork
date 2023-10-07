package managerServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func GetStudent() ([]models.Student, error) {
	var student_list []models.Student
	result := database.DB.Where(&models.Student{}).Find(&student_list)
	if result.Error != nil {
		return nil, result.Error
	}
	return student_list, nil
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
