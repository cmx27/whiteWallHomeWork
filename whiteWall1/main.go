package main

import (
	"log"
	"whiteWall/app/midwares"
	"whiteWall/config/database"
	"whiteWall/config/router"

	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
)

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	//store := cookie.NewStore([]byte("secret"))//
	//r.Use(sessions.Sessions("mysession", store))
	router.Init(r)
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// err := r.Run(":" + config.Config.GetString("server.port"))
	// if err != nil {
	// 	log.Fatal("ServerStartFailed", err)
	// }
	err := r.Run()
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}

}
