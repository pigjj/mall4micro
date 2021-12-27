package services

import (
	"errors"
	authHttpDto "github.com/jianghaibo12138/mall4micro/mall4micro-auth/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/dao/mall_shop"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/http_dto"
	"gorm.io/gorm"
)

//
// ShopList
// @Description: 获取商铺列表
// @Document:
// @param user
// @param gtx
// @return mall_shop.MallShopList
// @return *response.Response
// @return error
//
func ShopList(user *authHttpDto.HttpAuthenticateDTO, gtx *ctx.GinContext) (mall_shop.MallShopList, *response.Response, error) {
	if user == nil {
		return nil, response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return nil, response.DBConnResponse, err
	}
	var shopList mall_shop.MallShopList
	err = session.Transaction(func(tx *gorm.DB) error {
		err = shopList.FindShopByUser(tx, user.ID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		gtx.Logger.Errorf("[ShopList] user: %s, find shop err: %s", err.Error())
		return nil, nil, err
	}
	return shopList, response.SuccessResponse, nil
}

//
// ShopCreate
// @Description: 注册商铺
// @Document:
// @param shopDTO
// @param user
// @param gtx
// @return *response.Response
// @return error
//
func ShopCreate(shopDTO *http_dto.ShopDTO, user *authHttpDto.HttpAuthenticateDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if shopDTO == nil {
		return response.PayloadParseResponse, errors.New(response.PayloadParseResponse.Message)
	}
	if user == nil {
		return response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, err
	}
	var shop = mall_shop.MallShop{
		UserId:   &user.ID,
		ShopName: shopDTO.ShopName,
		ShopDesc: shopDTO.ShopDesc,
		ShopPic:  shopDTO.ShopPic,
		Status:   &mall_shop.StatusOnline,
	}
	err = session.Transaction(func(tx *gorm.DB) error {
		return shop.Create(tx)
	})
	if err != nil {
		gtx.Logger.Errorf("[ShopCreate] user: %s, create shop err: %s", user.Username, err.Error())
		return response.SQLExecResponse, err
	}
	return response.SuccessResponse, nil
}
