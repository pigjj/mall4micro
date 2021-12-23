package ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
)

type GinContext struct {
	*gin.Context
	Logger *log.ZapLogger
}

func NewGinContext(fn func(gtx GinContext), logger *log.ZapLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(GinContext{ctx, logger})
	}
}
