package managerControllers

import (
	"log"
	"whiteWall/app/services/managerServices"
	"whiteWall/app/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user_list, err := managerServices.GetUser()
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user_list": user_list,
	})
}

func GetArtical(c *gin.Context) {
	artical_list, err := managerServices.GetArtical()
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
	}
	utils.JsonSuccessResponse(c, gin.H{
		"artical_list": artical_list,
	})
}
