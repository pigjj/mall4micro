package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/middleware"
	commonHandlers "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/constant"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/pkg/handlers"
)

var logger *log.ZapLogger

func init() {
	logger = log.InitZapLogger(constant.MicroServiceName, conf.Settings.HttpServer.Debug)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(middleware.AuthMiddleWare(logger))

	url := r.Group("/api/shop")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandlers.PingHandler, logger))

		url.GET("/user_related_list", ctx.NewGinContext(handlers.ShopListHandler, logger))

		url.POST("/create", ctx.NewGinContext(handlers.ShopCreateHandler, logger))
	}
	return r
}
