package handlers

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"time"
)

func PingHandler(gtx *ctx.GinContext) {
	gtx.Logger.Infof("[PingHandler] now: %+v", time.Now())
	gtx.JsonWithData(response.SuccessResponse, "ping")
}
