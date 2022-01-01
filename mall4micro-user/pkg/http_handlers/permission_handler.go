package http_handlers

import (
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"github.com/pigjj/mall4micro/mall4micro-user/pkg/services"
)

func PermissionCreateHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err)
		return
	}
	var permissionDto http_dto.PermissionDTO
	err = gtx.ShouldBind(&permissionDto)
	if err != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err)
		return
	}
	res, err := services.CreatePermissionSrv(&permissionDto, user, gtx)
	gtx.JsonWithData(res, err)
}

func UserRelatedPermissionListHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err)
		return
	}
	res, err := services.UserRelatedPermissionListSrv(user, gtx)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, user)
}
