package userServices

import (
	"errors"
	"net/http"
	"whiteWall/app/services/userServices/configServices"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUserSession(c *gin.Context) error {
	webSession := sessions.Default(c)
	webSession.Options(sessions.Options{MaxAge: 3600 * 24 * 7, Path: "/api", Secure: true, SameSite: http.SameSiteNoneMode})
	webSession.Set("id", configServices.GetConfig("admin")) //configService//关于id和admin如何·在哪定义，获得
	return webSession.Save()
}

func GetUserSession(c *gin.Context) error {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return errors.New("")
	}
	if id != configServices.GetConfig("admin") {
		ClearUserSession(c)
		return errors.New("")
	}
	return nil
}

func UpdateUserSession(c *gin.Context) error { //为什么没有被调用
	err := GetUserSession(c)
	if err != nil {
		return err
	}
	err = SetUserSession(c)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserSession(c *gin.Context) bool {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return false
	}
	return true
}

func ClearUserSession(c *gin.Context) {
	webSession := sessions.Default(c)
	webSession.Delete("id")
	webSession.Save()
	return
}
