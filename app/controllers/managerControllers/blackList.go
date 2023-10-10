package managerControllers

import (
	"log"
	"whiteWall/app/services/managerServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type BlackListData struct {
	UserID uint `json:"user_id" binding:"required"`
}

func BlackList(c *gin.Context) {
	var data BlackListData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//鉴别权限
	_, err = userServices.GetManagerSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}

	//加入黑名单
	err = managerServices.BlackDeleteUserByUserID(data.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

func PutBlackList(c *gin.Context) {
	var data BlackListData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//鉴别权限
	_, err = userServices.GetManagerSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}

	//恢复
	err = managerServices.BlackPutUserByUserID(data.UserID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
