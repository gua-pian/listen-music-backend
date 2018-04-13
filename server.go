package main

import (
	"github.com/gin-gonic/gin"
	"littleprogram/handler"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Any("/api/search", handler.SearchSong)
	r.Any("/api/jscode2session", handler.Jscode2Session)
	r.Any("/api/getUserInfo", handler.GetUserInfo)

	r.Run(":4000")
}
