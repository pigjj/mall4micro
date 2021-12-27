package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/middleware"
	commonHandlers "github.com/jianghaibo12138/mall4micro/mall4micro-common/pkg/handlers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/constant"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/pkg/handlers"
)

var logger *log.ZapLogger

func init() {
	logger = log.InitZapLogger(constant.MicroServiceName, conf.Settings.HttpServer.Debug)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(middleware.AuthMiddleWare(logger))

	url := r.Group("/api/product")
	{
		url.GET("/ping", ctx.NewGinContext(commonHandlers.PingHandler, logger))

		url.GET("/category/list/:shop_id", ctx.NewGinContext(handlers.CategoryListHandler, logger))
		url.POST("/category/create", ctx.NewGinContext(handlers.CategoryCreateHandler, logger))

		url.GET("/product/list/:shop_id", ctx.NewGinContext(handlers.ProductListHandler, logger))
		url.POST("/product/create", ctx.NewGinContext(handlers.ProductCreateHandler, logger))
	}
	return r
}
