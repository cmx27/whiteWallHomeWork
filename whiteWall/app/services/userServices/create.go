package userServices

import (
	"whiteWall/app/models"
	"whiteWall/app/utils"
	"whiteWall/config/database"
)

func CreateUser(password, account, name, sex, major, token string) error {
	pass := utils.Encryrpt(password)
	var manager_state bool
	if token == "五花肉串串" {
		manager_state = true
		manager := &models.Manager{
			Account:      account,
			Password:     pass,
			Name:         name,
			ManagerState: manager_state,
		}
		result := database.DB.Create(manager)
		if result.Error != nil {
			return result.Error
		} else {
			return nil
		}

	} else {
		manager_state = false
		user := &models.User{
			Account:      account,
			Password:     pass,
			Name:         name,
			Sex:          sex,
			Major:        major,
			ManagerState: manager_state,
		}
		result := database.DB.Create(user)
		if result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}

}
