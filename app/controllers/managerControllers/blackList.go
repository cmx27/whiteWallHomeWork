package managerControllers

import (
	"log"
	"whiteWall/app/services/managerServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
	//"github.com/robfig/cron/v3"
)

type BlackListData struct {
	Time   int  `json:"time" binding:"required"`
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
	_, err = userServices.GetUserSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}

	//删除
	err = managerServices.BlackDeleteUserByUserID(data.UserID, data.Time)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
