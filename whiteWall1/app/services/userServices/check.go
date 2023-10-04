package userServices

import (
	"whiteWall/app/models"
	"whiteWall/app/utils"
	"whiteWall/config/database"
)

func CheckUserExistByAccount(account string) error {
	result := database.DB.Where(models.User{
		Account: account,
	}).First(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func CheckUserBYAccountAndPassword(account, password string) bool {
	pass := utils.Encryrpt(password)
	println(pass)
	result := database.DB.Where(
		models.User{
			Account:  account,
			Password: pass,
		}).First(&models.User{})
	if result.Error != nil {
		return result.Error != nil
	}
	return result.Error == nil

}

// func CheckUserExistByAccount(account string) error {
// 	result := database.DB.Where(models.User{
// 		Account: account,
// 	}).First(&models.User{})
// 	if result.Error != nil {
// 		result2 := database.DB.Where(models.Manager{
// 			Account: account,
// 		}).First(&models.Manager{})
// 		return result2.Error
// 	} else {
// 		return result.Error
// 	}
// }

// func CheckUserBYAccountAndPassword(account, password string) bool {
// 	pass := utils.Encryrpt(password)
// 	println(pass)
// 	result := database.DB.Where(
// 		models.User{
// 			Account:  account,
// 			Password: pass,
// 		}).First(&models.User{})
// 	if result.Error != nil {
// 		result2 := database.DB.Where(models.Manager{
// 			Account:  account,
// 			Password: pass,
// 		}).First(&models.Manager{})
// 		return result2.Error == nil
// 	} else {
// 		return result.Error == nil
// 	}
// }
