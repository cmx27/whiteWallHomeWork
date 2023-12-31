package articalControllers

import (
	"log"

	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type PUTArticalData struct {
	ArticalID uint   `json:"artical_id" binding:"required"`
	Artical   string `json:"artical" binding:"required"`
}

//更新文章

func PUTArtical(c *gin.Context) {
	var data PUTArticalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
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

	// 判断是否是同一个人，获取文章信息

	_, err = articalServices.GetArticalByUserIDAndArticalID(student.UserID, data.ArticalID)
	if err != nil {
		utils.JsonErrorResponse(c, 406, "参数错误（不是同一个人或文章不存在）")
		return
	}
	//更新
	err = articalServices.UpdateArtical(data.Artical, data.ArticalID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
