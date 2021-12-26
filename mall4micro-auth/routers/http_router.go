package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/constant"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/pkg/http_handlers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/middleware"
	commonHandlers "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
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

	r.Use(middleware.AuthMiddleWare())

	url := r.Group("/api/auth")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandlers.PingHandler, logger))

		url.POST("/login", ctx.NewGinContext(http_handlers.LoginHandler, logger))
		url.POST("/authenticate", ctx.NewGinContext(http_handlers.AuthenticateHandler, logger))

	}
	return r
}
