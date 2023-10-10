package articalControllers

import (
	"log"
	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type CreateArticalData struct {
	Namestate *bool  `json:"name_state" binding:"required"`
	Artical   string `json:"artical" binding:"required"`
}

// 发布文章
func CreateArtical(c *gin.Context) {
	var data CreateArticalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//鉴别权限
	student, err := userServices.GetStudentSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}
	//判断是否在黑名单
	black := student.BlackState
	if black {
		utils.JsonResponse(c, 405, 400, "您在小黑屋", nil)
		return
	}

	//存入数据库
	err = articalServices.CreateArtical(student.Name, data.Artical, *data.Namestate, student.UserID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
