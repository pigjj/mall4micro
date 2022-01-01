package handlers

import (
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"time"
)

func PingHandler(gtx *ctx.GinContext) {
	gtx.Logger.Infof("[PingHandler] now: %+v", time.Now())
	gtx.JsonWithData(response.SuccessResponse, "ping")
}
