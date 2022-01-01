package http_handlers

import (
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"github.com/pigjj/mall4micro/mall4micro-user/pkg/services"
)

func GroupCreateHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err)
		return
	}
	var groupDto http_dto.GroupDTO
	err = gtx.ShouldBind(&groupDto)
	if err != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err)
		return
	}
	res, err := services.GroupCreateSrv(&groupDto, user, gtx)
	gtx.JsonWithData(res, err)
}

func UserRelatedGroupHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err)
		return
	}
	res, err := services.UserRelatedGroupListSrv(user, gtx)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, user)
}
