package handlers

import (
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-product/http_dto"
	"github.com/pigjj/mall4micro/mall4micro-product/pkg/services"
	"strconv"
)

func ProductListHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err.Error())
		return
	}
	shopIdStr := gtx.Param("shop_id")
	shioId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		gtx.JsonWithData(response.ParamParseResponse, err.Error())
		return
	}
	list, res, err := services.ProductListSrv(shioId, user, gtx)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, list)
}

func ProductCreateHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err.Error())
		return
	}
	var productDto http_dto.ProductDTO
	err = gtx.ShouldBind(&productDto)
	if err != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err.Error())
		return
	}
	res, err := services.ProductCreateSrv(&productDto, user, gtx)
	gtx.JsonWithData(res, err)
}
