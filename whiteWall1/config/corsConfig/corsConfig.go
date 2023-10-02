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
	corsConfig.AddAllowMethods("OPTIONS")
	return cors.New(corsConfig)
}
