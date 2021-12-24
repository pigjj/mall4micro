package handlers

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/pkg/services"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
)

func LoginHandler(gtx *ctx.GinContext) {
	var d dto.LoginDTO
	err := gtx.Context.ShouldBind(&d)
	if err != nil || d.UsernameValidate() != nil || d.PasswordValidate() != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err)
		return
	}
	token, res, err := services.LoginService(&d)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, token)
}
