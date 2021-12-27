package handlers

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/pkg/services"
)

func ShopListHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err)
		return
	}
	shopList, res, err := services.ShopList(user, gtx)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, shopList)
}

func ShopCreateHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err)
		return
	}
	var shopDTO http_dto.ShopDTO
	err = gtx.ShouldBind(&shopDTO)
	if err != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err)
		return
	}
	res, err := services.ShopCreate(&shopDTO, user, gtx)
	gtx.JsonWithData(res, err)
}
