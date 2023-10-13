package commentControllers

import (
	"log"
	"whiteWall/app/services/studentServices/commentServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type CreateCommentData struct {
	Comment   string `json:"comment" binding:"required"`
	ArticalID uint   `json:"artical_id" binding:"required"`
}

// 发布评论
func CreateComment(c *gin.Context) {
	var data CreateCommentData
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
	err = commentServices.CreateComment(student.Name, data.Comment, student.UserID, data.ArticalID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
