package session

import (
	"whiteWall/config/redis"

	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func setRedis(r *gin.Engine, name string) {
	Info := redis.RedisInfo
	store, _ := sessionRedis.NewStore(10, "tcp", Info.Host+":"+Info.Port, Info.Password, []byte("secret"))
	r.Use(sessions.Sessions(name, store))
}
