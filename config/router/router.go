package router

import (
	"whiteWall/app/controllers/managerControllers"
	"whiteWall/app/controllers/studentControllers/articalControllers"
	"whiteWall/app/controllers/studentControllers/commentControllers.go"
	"whiteWall/app/controllers/userControllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/user/login", userControllers.Login)
		api.POST("/user/reg", userControllers.Register)

		artical := api.Group("/student/wall-artical")
		{
			artical.POST("", articalControllers.CreateArtical)
			artical.PUT("", articalControllers.PUTArtical)
			artical.DELETE("", articalControllers.DeleteArtical)
			artical.GET("", articalControllers.GetArtical)
			artical.GET("/others", articalControllers.GetOthersArtical)
		}

		comment := api.Group("/student/wall-comment")
		{
			comment.POST("", commentControllers.CreateComment)
			comment.PUT("", commentControllers.PUTComment)
			comment.DELETE("", commentControllers.DeleteComment)
			comment.GET("", commentControllers.GetMyComment)
			comment.GET("/artical-comment", commentControllers.GetArticalComment)
		}

		//api.POST("/student/wall-image", imageControllers.CreateImage)

		manager := api.Group("/manager")
		{
			manager.DELETE("/artical", managerControllers.MDeleteArtical)
			manager.DELETE("/user", managerControllers.MDeleteUser)
			manager.DELETE("/black-list", managerControllers.BlackList)
			manager.PUT("/black-list", managerControllers.PutBlackList)
			manager.GET("/user", managerControllers.GetUser)
			manager.GET("/artical", managerControllers.GetArtical)
			manager.GET("/comment", managerControllers.GetComment)
			manager.DELETE("/comment", managerControllers.MDeleteComment)
		}
	}
}
