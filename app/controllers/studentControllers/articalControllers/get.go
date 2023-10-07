package articalControllers

import (
	"log"
	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

func GetArtical(c *gin.Context) {
	//鉴别权限
	student, err := userServices.GetStudentSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}
	artical_list, err := articalServices.GetArticalByUserID(student.StudentID)
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
