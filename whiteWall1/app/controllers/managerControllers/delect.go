package managerControllers

import (
	"log"
	"whiteWall/app/services/managerServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type MDeleteArticalData struct {
	ArticalID uint `json:"artical_id" binding:"required"`
}

// 删除文章
func MDeleteArtical(c *gin.Context) {
	var data MDeleteArticalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}

	//删除
	err = managerServices.DeleteArticalByArticalID(data.ArticalID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type MDeleteUserData struct {
	UserID uint `json:"user_id" binding:"required"`
}

// 删除用户
func MDeleteUser(c *gin.Context) {
	var data MDeleteUserData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}

	//删除
	err = managerServices.DeleteUserByUserID(data.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
