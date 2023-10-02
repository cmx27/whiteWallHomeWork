package userControllers

import (
	"log"
	//"whiteWall/app/models"
	"whiteWall/app/services/userServices"
	"whiteWall/app/services/userServices/configServices"
	"whiteWall/app/utils"

	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录
func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Panicln(err)
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断用户是否存在,获取用户信息
	user, err := userServices.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 404, "用户不存在")
		return
	}

	//session,转移到userServices

	if configServices.GetConfig("password") != data.Password {

		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userServices.SetUserSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// session := sessions.Default(c)
	// v := session.Get("id")
	// if v == nil {
	// 	session.Set("id", id)
	// 	session.Set("manager_state", manager_state)
	// 	session.Save()
	// }
	// c.JSON(200, gin.H{"id": id, "manager_state": manager_state})

	//判断密码是否正确
	// flag := userServices.CheckUserBYAccountAndPassword(data.Account, data.Password)
	// if !flag {
	// 	utils.JsonResponse(c, 405, 400, "密码错误", nil)
	// 	return
	// }

	// 返回用户信息
	utils.JsonSuccessResponse(c, gin.H{
		"list": user,
	})
}
