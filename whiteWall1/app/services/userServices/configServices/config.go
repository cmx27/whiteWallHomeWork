package configServices

import (
	"context"
	"whiteWall/config/redis"
)

var ctx = context.Background() //有什么用

func GetConfig(key string) string {
	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err == nil {
		return val
	}
	return ""
}
