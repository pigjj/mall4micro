package handlers

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/pkg/services"
	"strconv"
)

func CategoryListHandler(gtx *ctx.GinContext) {
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
	list, res, err := services.CategoryListSrv(shioId, user, gtx)
	if err != nil {
		gtx.JsonWithData(res, err)
		return
	}
	gtx.JsonWithData(res, list)
}

func CategoryCreateHandler(gtx *ctx.GinContext) {
	user, err := gtx.GetUser()
	if err != nil {
		gtx.JsonWithData(response.UserNotLoginResponse, err.Error())
		return
	}
	var categoryDto http_dto.CategoryDTO
	err = gtx.ShouldBind(&categoryDto)
	if err != nil {
		gtx.JsonWithData(response.PayloadParseResponse, err.Error())
		return
	}
	res, err := services.CategoryCreateSrv(&categoryDto, user, gtx)
	gtx.JsonWithData(res, err)
}
