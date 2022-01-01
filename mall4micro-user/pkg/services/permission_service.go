package services

import (
	"errors"
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	cm "github.com/pigjj/mall4micro/mall4micro-common/models"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-user/dao/mall_sys_group_permission_relation"
	"github.com/pigjj/mall4micro/mall4micro-user/dao/mall_sys_permission"
	"github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"gorm.io/gorm"
)

//
// CreatePermissionSrv
// @Description: 创建权限数据
// @Document:
// @param permissionDto
// @param user
// @param gtx
// @return *response.Response
// @return error
//
func CreatePermissionSrv(permissionDto *http_dto.PermissionDTO, user *http_dto.UserDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if permissionDto == nil {
		return response.PayloadParseResponse, errors.New(response.PayloadParseResponse.Message)
	}
	if user == nil {
		return response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, err
	}
	err = session.Transaction(func(tx *gorm.DB) error {
		var permission = mall_sys_permission.MallSysPermission{
			MallBase: cm.MallBase{
				CreateUserId: user.ID,
			},
			PermissionName: permissionDto.PermissionName,
			PermissionDesc: permissionDto.PermissionDesc,
		}
		return permission.Create(tx)
	})
	if err != nil {
		gtx.Logger.Errorf("[CreatePermissionSrv] user: %s, create permission err: %+v", user.Username, permissionDto)
		return response.SQLExecResponse, err
	}
	return response.SuccessResponse, nil
}

//
// UserRelatedPermissionListSrv
// @Description: 获取用户关联的权限列表接口
// @Document:
// @param user
// @param gtx
// @return *response.Response
// @return error
//
func UserRelatedPermissionListSrv(user *http_dto.UserDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if user == nil {
		return response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	res, err := UserRelatedGroupListSrv(user, gtx)
	if err != nil {
		return res, err
	}
	var groupIds []uint
	for _, group := range user.GroupList {
		groupIds = append(groupIds, group.ID)
	}
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, errors.New(response.DBConnResponse.Message)
	}
	var relationList mall_sys_group_permission_relation.MallSysGroupPermissionRelationList
	var permissionList mall_sys_permission.MallSysPermissionList
	err = session.Transaction(func(tx *gorm.DB) error {
		err = relationList.SelectRelationByGroupIds(tx, groupIds)
		if err != nil {
			return err
		}
		var permissionIds []uint
		for _, relation := range relationList {
			permissionIds = append(permissionIds, *relation.PermissionId)
		}
		err = permissionList.SelectPermissionByIds(tx, permissionIds)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		gtx.Logger.Errorf("[UserRelatedPermissionListSrv] user: %s get permission err: %s", user.Username, err.Error())
		return response.SQLExecResponse, err
	}
	for _, permission := range permissionList {
		user.PermissionList = append(user.PermissionList, http_dto.PermissionDTO{
			ID:             permission.ID,
			PermissionName: permission.PermissionName,
			PermissionDesc: permission.PermissionDesc,
		})
	}
	return response.SuccessResponse, nil
}
