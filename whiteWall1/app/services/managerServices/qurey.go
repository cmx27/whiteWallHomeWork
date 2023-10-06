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
	//fmt.Println(models.Num)
	//artical_list := make([]models.Artical, models.Num)
	result := database.DB.Find(&artical_list)
	if result.Error != nil {
		return nil, result.Error
	}

	// for i := 0; i < int(models.Num); i++ {
	// 	if artical_list[i].Namestate {
	// 		log.Println(i)
	// 		var user []models.User
	// 		result2 := database.DB.Where(&models.User{
	// 			UserID: artical_list[i].UserID,
	// 		}).First(&user[i])
	// 		if result2.Error != nil {
	// 			artical_list[i].Name = user[i].Name
	// 		} else {
	// 			return nil, result2.Error
	// 		}
	// 	}
	// }
	return artical_list, nil
}
