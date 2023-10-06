package userServices

import (
	"whiteWall/app/models"
	"whiteWall/app/utils"
	"whiteWall/config/config"
	"whiteWall/config/database"
)

// func CreateUser(password, account, name, sex, major string) error {
// 	pass := utils.Encryrpt(password)
// 	//var manager_state bool
// 	manager_state := false
// 	user := &models.User{
// 		Account:      account,
// 		Password:     pass,
// 		Name:         name,
// 		Sex:          sex,
// 		Major:        major,
// 		ManagerState: manager_state,
// 	}
// 	result := database.DB.Create(user)
// 	if result.Error != nil {
// 		return result.Error
// 	} else {
// 		return nil
// 	}
// }

func CreateUser(password, account, name, sex, major, token string) error {
	pass := utils.Encryrpt(password)
	var manager_state bool
	value := config.Config.GetString("key.value")
	if token == value {
		manager_state = true

		manager := &models.Manager{
			Account:      account,
			Password:     pass,
			Name:         name,
			ManagerState: manager_state,
		}
		result2 := database.DB.Create(manager)
		if result2.Error != nil {
			return result2.Error
		}
		user := &models.User{
			UserID:       manager.ManagerID,
			Account:      account,
			Password:     pass,
			Name:         name,
			ManagerState: manager_state,
		}
		result1 := database.DB.Create(user)
		if result1.Error != nil {
			return result1.Error
		} else {
			return nil
		}

	} else {
		manager_state = false

		student := &models.Student{

			Sex:          sex,
			Major:        major,
			Account:      account,
			Password:     password,
			Name:         name,
			ManagerState: manager_state,
		}
		result2 := database.DB.Create(student)
		if result2.Error != nil {
			return result2.Error
		}
		user := &models.User{
			UserID:       student.StudentID,
			Account:      account,
			Password:     pass,
			Name:         name,
			ManagerState: manager_state,
		}
		result1 := database.DB.Create(user)
		if result1.Error != nil {
			return result1.Error
		} else {
			return nil
		}
	}

}
