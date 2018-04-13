package handler

import (
	"github.com/go-redis/redis"
)

var (
	AppId       = "" // 填上自己的 AppId
	AppSecret   = "" // 填上自己的 AppSecret
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
