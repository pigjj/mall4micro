package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	commonHandler "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	url := r.Group("/api/auth")
	{
		url.GET("/ping", commonHandler.PingHandler)
	}
	return r
}
