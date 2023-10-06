package articalControllers

import (
	"log"
	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type CreateArticalData struct {
	//Name      string `json:"name" binding:"required"`
	Namestate *bool  `json:"name_state" binding:"required"`
	Artical   string `json:"artical" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
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
	//getUserByID
	student,err := articalServices.GetStusentByUserID(data.UserID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//存入数据库
	err = articalServices.CreateArtical(student.Name, data.Artical, *data.Namestate, data.UserID)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
