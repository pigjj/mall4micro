package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	commonHandler "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
)

const MicroServiceName = "mall4micro-auth"

var logger *log.ZapLogger

func init() {
	logger = log.InitZapLogger(MicroServiceName, conf.Settings.Server.Debug)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	url := r.Group("/api/auth")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandler.PingHandler, logger))
	}
	return r
}
