package commentControllers

import (
	"log"
	"whiteWall/app/services/studentServices/commentServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type PUTCommentData struct {
	CommentID uint   `json:"comment_id" binding:"required"`
	Comment   string `json:"comment" binding:"required"`
}

//更新文章

func PUTComment(c *gin.Context) {
	var data PUTCommentData
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

	// 判断是否是同一个人
	_, err = commentServices.GetCommmentByUserIDAndCommentID(student.UserID, data.CommentID)
	if err != nil {
		utils.JsonErrorResponse(c, 406, "参数错误（不是同一个人或评论不存在）")
		return
	}
	//更新
	err = commentServices.UpdateComment(data.Comment, data.CommentID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
