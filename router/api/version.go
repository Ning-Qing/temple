package api

import (
	"github.com/Ning-Qing/temple/global"
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H{"verson": global.Version})
}
