package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/pkg/handlers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	commonHandlers "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
)

const MicroServiceName = "mall4micro-auth"

var logger *log.ZapLogger

func init() {
	logger = log.InitZapLogger(MicroServiceName, conf.Settings.Server.Debug)
}

//
// InitRouter
// @Description: 初始化路由函数
// @return *gin.Engine
//
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	url := r.Group("/api/auth")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandlers.PingHandler, logger))
		
		url.POST("/login", ctx.NewGinContext(handlers.LoginHandler, logger))

	}
	return r
}
