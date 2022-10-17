package router

import (
	"github.com/Ning-Qing/temple/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.GlobalSettings.GetServerMode())
	r := gin.Default()

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/version")
	}
	r.GET("/ping", Ping)
	return r
}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}
