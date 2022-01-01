package http_handlers

import (
	"github.com/pigjj/mall4micro/mall4micro-auth/http_dto"
	"github.com/pigjj/mall4micro/mall4micro-auth/pkg/services"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
)

func LoginHandler(gtx *ctx.GinContext) {
	var d http_dto.HttpLoginDTO
	err := gtx.Context.ShouldBind(&d)
	if err != nil || d.UsernameValidate() != nil || d.PasswordValidate() != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err)
		return
	}
	token, res, err := services.LoginService(&d, gtx)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, token)
}

func AuthenticateHandler(gtx *ctx.GinContext) {
	token := gtx.GetHeader("Authorization")
	res, userInfo, err := services.AuthorizedService(token)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, userInfo)
}
