package services

import (
	"errors"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	cm "github.com/jianghaibo12138/mall4micro/mall4micro-common/models"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/dao/mall_product"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/http_dto"
	userHttpDto "github.com/jianghaibo12138/mall4micro/mall4micro-user/http_dto"
	"gorm.io/gorm"
)

func ProductListSrv(shopId int, user *userHttpDto.UserDTO, gtx *ctx.GinContext) (*[]http_dto.ProductDTO, *response.Response, error) {
	if user == nil {
		return nil, response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return nil, response.DBConnResponse, errors.New(response.DBConnResponse.Message)
	}
	var productList mall_product.MallProductList
	err = session.Transaction(func(tx *gorm.DB) error {
		return productList.SelectProductByShopId(tx, uint(shopId))
	})
	if err != nil {
		gtx.Logger.Errorf("[ProductListSrv] user: %s, find product list err: %s", user.Username, err.Error())
		return nil, response.SQLExecResponse, err
	}
	var productDtoList []http_dto.ProductDTO
	for _, product := range productList {
		productDtoList = append(productDtoList, http_dto.ProductDTO{
			ShopId:      *product.ShopId,
			ProductName: product.ProductName,
			OriPrice:    *product.OriPrice,
			Price:       *product.Price,
			Brief:       product.Brief,
			Content:     product.Content,
			Pic:         product.Pic,
			Status:      *product.Status,
			CategoryId:  *product.CategoryId,
			SoldNum:     *product.SoldNum,
			TotalStocks: *product.TotalStocks,
			PutOnTime:   product.PutOnTime.UnixMicro(),
		})
	}
	return &productDtoList, response.SuccessResponse, err
}

func ProductCreateSrv(productDto *http_dto.ProductDTO, user *userHttpDto.UserDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if productDto == nil {
		return response.PayloadParseResponse, errors.New(response.PayloadParseResponse.Message)
	}
	if user == nil {
		return response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, errors.New(response.DBConnResponse.Message)
	}
	err = session.Transaction(func(tx *gorm.DB) error {
		var product = mall_product.MallProduct{
			MallBase: cm.MallBase{
				CreateUserId: user.ID,
			},
			ShopId:      &productDto.ShopId,
			ProductName: productDto.ProductName,
			OriPrice:    &productDto.OriPrice,
			Price:       &productDto.Price,
			Brief:       productDto.Brief,
			Content:     productDto.Content,
			Pic:         productDto.Pic,
			CategoryId:  &productDto.CategoryId,
			TotalStocks: &productDto.TotalStocks,
			Status:      &mall_product.StatusOnline,
		}
		return product.Create(tx)
	})
	if err != nil {
		gtx.Logger.Errorf("[CategoryCreateSrv] user: %s, create category err: %s", err.Error())
		return response.SQLExecResponse, err
	}
	return response.SuccessResponse, nil
}
