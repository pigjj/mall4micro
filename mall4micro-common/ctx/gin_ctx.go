package ctx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GinContext struct {
	*gin.Context
	Logger *zap.Logger
}

func NewGinContext(fn func(gtx GinContext), logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(GinContext{ctx, logger})
	}
}
