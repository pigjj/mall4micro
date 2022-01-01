package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/log"
	"github.com/pigjj/mall4micro/mall4micro-common/middleware"
	commonHandlers "github.com/pigjj/mall4micro/mall4micro-common/pkg/handlers"
	"github.com/pigjj/mall4micro/mall4micro-user/constant"
	"github.com/pigjj/mall4micro/mall4micro-user/pkg/http_handlers"
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

		url.POST("/register", ctx.NewGinContext(http_handlers.RegisterHandler, logger))

		url.GET("/user_related_group_list", ctx.NewGinContext(http_handlers.UserRelatedGroupHandler, logger))
		url.POST("/group_create", ctx.NewGinContext(http_handlers.GroupCreateHandler, logger))

		url.GET("/user_related_permission_list", ctx.NewGinContext(http_handlers.UserRelatedPermissionListHandler, logger))
		url.POST("/permission_create", ctx.NewGinContext(http_handlers.PermissionCreateHandler, logger))
	}
	return r
}
