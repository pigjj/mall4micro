package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"time"
)

func PingHandler(context ctx.GinContext) {
	context.Logger.Infof("[PingHandler] now: %+v", time.Now())
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
