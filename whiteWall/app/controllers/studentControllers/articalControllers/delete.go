package articalControllers

import (
	"whiteWall/app/models"
	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

type DeleteArticalData struct {
	//Name      string `json:"name" binding:"required"`
	ArticalID uint `json:"artical_id" binding:"required"`
	UserID    uint `json:"user_id" binding:"required"`
}

// 删除文章
func DeleteArtical(c *gin.Context) {
	var data DeleteArticalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断是否是同一个人，获取文章信息
	var artical *models.Artical
	artical, err = articalServices.GetArticalByUserIDAndArticalID(data.UserID, data.ArticalID)
	if err != nil {
		utils.JsonErrorResponse(c, 406, "参数错误（不是同一个人或文章不存在）")
		return
	}

	//删除
	err = articalServices.DeleteArtical(artical.ArticalID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
