package userControllers

import (
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

type RegisterData struct {
	Account  string `json:"account" binding:"required" `
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Sex      string `json:"sex"`
	Major    string `json:"major"`
	Token    string `json:"token"`
}

// 注册
func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断号是否已经注册
	err = userServices.CheckUserExistByAccount(data.Account)
	if err == nil {
		utils.JsonErrorResponse(c, 200504, "账号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//限制密码长度
	if len(data.Password) < 8 || len(data.Password) > 20 {
		utils.JsonResponse(c, 200, 401, "密码长度必须在8~20位之间", nil)
		return
	}

	//创建用户
	err = userServices.CreateUser(data.Password, data.Account, data.Name, data.Sex, data.Major, data.Token)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
