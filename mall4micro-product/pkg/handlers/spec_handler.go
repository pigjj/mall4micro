package handlers

import (
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/pkg/services"
)

func SpecListHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err.Error())
		return
	}
	fmt.Println(user)
	services.SpecList(gtx, user)
	gtx.Json(response.SuccessResponse)
}

func SpecCreateHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err.Error())
		return
	}
	fmt.Println(user)
	gtx.Json(response.SuccessResponse)
}
