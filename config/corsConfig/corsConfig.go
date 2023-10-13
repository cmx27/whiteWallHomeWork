package corsConfig

import (
	"whiteWall/config/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Config.GetString("allow_origins")}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "POST", "GET", "PUT", "DELETE", "UPDATE")
	// 设置允许的请求头
	//corsConfig.AddAllowHeaders("Origin", "X-Requested-With", "X-Extra-Header", "Content-Type", "Accept", "Authorization")

	// 设置暴露的响应头
	//corsConfig.ExposeHeaders = []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type"}

	return cors.New(corsConfig)
}
