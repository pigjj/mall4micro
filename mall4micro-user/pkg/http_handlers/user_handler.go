package http_handlers

import (
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"github.com/pigjj/mall4micro/mall4micro-user/pkg/services"
	"time"
)

func RegisterHandler(gtx *ctx.GinContext) {
	gtx.Logger.Infof("[RegisterHandler] now: %+v", time.Now())
	var d http_dto.RegisterDTO
	err := gtx.Context.ShouldBind(&d)
	if err != nil || d.UsernameValidate() != nil || d.PasswordValidate() != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err)
		return
	}
	res, err := services.RegisterSrv(&d)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.Json(res)
}
