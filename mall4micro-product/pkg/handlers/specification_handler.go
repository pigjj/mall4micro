package handlers

import (
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
)

func SpecListHandler(gtx *ctx.GinContext) {
	userItl, ok := gtx.Get("user")
	if !ok {
		gtx.Json(response.UserNotLoginResponse)
		return
	}
	user, ok := userItl.(http_dto.HttpAuthenticateDTO)
	if !ok {
		gtx.Json(response.UserNotLoginResponse)
		return
	}
	fmt.Println(user)
	gtx.Json(response.SuccessResponse)
}
