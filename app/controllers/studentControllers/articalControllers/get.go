package articalControllers

import (
	"log"
	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type GetArticalData struct {
	UserID uint `json:"user_id" binding:"required"`
}

func GetArtical(c *gin.Context) {
	var data GetArticalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}

	artical_list, err := articalServices.GetArticalByUserID(data.UserID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}
	utils.JsonSuccessResponse(c, gin.H{
		"artical_list": artical_list,
	})
}

type GetOtherArticalData struct {
	ArticalID uint `json:"artical_id" binding:"required"`
}

func GetOthersArtical(c *gin.Context) {
	var data GetOtherArticalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}

	artical, err := articalServices.GetArticalByArticalID(data.ArticalID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}
	utils.JsonSuccessResponse(c, artical)
}
