package userServices

import (
	"errors"
	"net/http"
	"whiteWall/app/models"
	"whiteWall/app/services/managerServices"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUserSession(c *gin.Context, user *models.User) error {
	webSession := sessions.Default(c)
	webSession.Options(sessions.Options{MaxAge: 3600 * 24 * 7, Path: "/api", Secure: false, SameSite: http.SameSiteNoneMode})
	webSession.Set("id", user.UserID)
	webSession.Set("user_id", user.UserID)
	webSession.Set("state", user.ManagerState)

	return webSession.Save()
}

func GetManagerSession(c *gin.Context) (*models.Manager, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("user_id")
	state := webSession.Get("state")
	if id == nil {
		return nil, errors.New("id")
	}
	if state == 0 {
		return nil, errors.New("stste")
	}
	manager, _ := managerServices.GetManagerByID(id.(uint))
	if manager == nil {
		ClearUserSession(c)
		return nil, errors.New("manager")
	}
	return manager, nil
}

func GetStudentSession(c *gin.Context) (*models.User, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("user_id")
	if id == nil {
		return nil, errors.New("id")
	}
	student, _ := GetUserByUserID(id.(uint))
	if student == nil {
		ClearUserSession(c)
		return nil, errors.New("user")
	}
	return student, nil
}

// func UpdateUserSession(c *gin.Context) error { //为什么没有被调用
// 	err := GetManagerSession(c)
// 	if err != nil {
// 		return err
// 	}
// 	err = SetUserSession(c)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func CheckUserSession(c *gin.Context) bool {
	webSession := sessions.Default(c)
	id := webSession.Get("user_id")
	if id == nil {
		return false
	}
	return true
}

func ClearUserSession(c *gin.Context) {
	webSession := sessions.Default(c)
	webSession.Delete("user_id")
	webSession.Save()
	return
}
