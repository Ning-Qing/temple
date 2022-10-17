package router

import (
	"github.com/Ning-Qing/temple/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.GlobalSettings.Server.Mode)
	r := gin.Default()

	return r
}
