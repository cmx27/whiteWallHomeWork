package main

import (
	"log"
	"whiteWall/app/midwares"

	//"whiteWall/config/corsConfig"
	"whiteWall/config/database"
	"whiteWall/config/router"
	"whiteWall/config/session"

	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
)

func main() {
	database.Init()
	r := gin.Default()
	//r.Use(corsConfig.GetCors())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	session.Init(r)
	router.Init(r)
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	err := r.Run()
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
