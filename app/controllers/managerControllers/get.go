package managerControllers

import (
	"log"
	"whiteWall/app/services/managerServices"
	"whiteWall/app/services/userServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	//鉴别权限
	_, err := userServices.GetManagerSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}
	student_list, err := managerServices.GetStudent()
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user_list": student_list,
	})
}

func GetArtical(c *gin.Context) {
	//鉴别权限
	_, err := userServices.GetManagerSession(c)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 500, "session")
		return
	}
	artical_list, err := managerServices.GetArtical()
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}
	utils.JsonSuccessResponse(c, gin.H{
		"artical_list": artical_list,
	})
}
