package services

import (
	"errors"
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	cm "github.com/pigjj/mall4micro/mall4micro-common/models"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-product/dao/mall_category"
	"github.com/pigjj/mall4micro/mall4micro-product/http_dto"
	userHttpDto "github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"gorm.io/gorm"
)

//
// CategoryListSrv
// @Description: 产品类目列表接口
// @Document:
// @param shopId
// @param user
// @param gtx
// @return *[]http_dto.CategoryDTO
// @return *response.Response
// @return error
//
func CategoryListSrv(shopId int, user *userHttpDto.UserDTO, gtx *ctx.GinContext) (*[]http_dto.CategoryDTO, *response.Response, error) {
	if user == nil {
		return nil, response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return nil, response.DBConnResponse, errors.New(response.DBConnResponse.Message)
	}
	var categoryList mall_category.MallCategoryList
	err = session.Transaction(func(tx *gorm.DB) error {
		return categoryList.SelectCategoryByShopId(tx, uint(shopId))
	})
	if err != nil {
		gtx.Logger.Errorf("[CategoryListSrv] user: %s, find category list err: %s", user.Username, err.Error())
		return nil, response.SQLExecResponse, err
	}
	var categoryDtoList []http_dto.CategoryDTO
	for _, category := range categoryList {
		categoryDtoList = append(categoryDtoList, http_dto.CategoryDTO{
			ID:           category.ID,
			ShopId:       *category.ShopId,
			CategoryName: category.CategoryName,
			Icon:         category.Icon,
			Pic:          category.Pic,
			Status:       *category.Status,
		})
	}
	return &categoryDtoList, response.SuccessResponse, nil
}

//
// CategoryCreateSrv
// @Description: 创建产品类目接口
// @Document:
// @param categoryDto
// @param user
// @param gtx
// @return *response.Response
// @return error
//
func CategoryCreateSrv(categoryDto *http_dto.CategoryDTO, user *userHttpDto.UserDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if categoryDto == nil {
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
		var category = mall_category.MallCategory{
			MallBase: cm.MallBase{
				CreateUserId: user.ID,
			},
			ShopId:       &categoryDto.ShopId,
			CategoryName: categoryDto.CategoryName,
			Icon:         categoryDto.Icon,
			Pic:          categoryDto.Pic,
			Status:       &mall_category.StatusOnline,
		}
		return category.Create(tx)
	})
	if err != nil {
		gtx.Logger.Errorf("[CategoryCreateSrv] user: %s, create category err: %s", err.Error())
		return response.SQLExecResponse, err
	}
	return response.SuccessResponse, nil
}
