package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	commonHandlers "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
)

const MicroServiceName = "mall4micro-product"

var logger *log.ZapLogger

func init() {
	logger = log.InitZapLogger(MicroServiceName, conf.Settings.Server.Debug)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	url := r.Group("/api/product")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandlers.PingHandler, logger))
	}
	return r
}
