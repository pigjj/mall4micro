package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/middleware"
	commonHandlers "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/constant"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/pkg/handlers"
)

var logger *log.ZapLogger

func init() {
	logger = log.InitZapLogger(constant.MicroServiceName, conf.Settings.HttpServer.Debug)
}

//
// InitRouter
// @Description: 初始化路由函数
// @return *gin.Engine
//
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(middleware.AuthMiddleWare(logger))
	
	url := r.Group("/api/user")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandlers.PingHandler, logger))

		url.POST("/register", ctx.NewGinContext(handlers.RegisterHandler, logger))
	}
	return r
}
