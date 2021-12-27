package services

import (
	"errors"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_group"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_group_permission_relation"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_permission"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/http_dto"
	"gorm.io/gorm"
)

//
// GroupCreateSrv
// @Description: 创建权限组
// @Document:
// @param groupInfo
// @param user
// @param gtx
// @return *response.Response
// @return error
//
func GroupCreateSrv(groupInfo *http_dto.GroupDTO, user *http_dto.UserDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if groupInfo == nil {
		return response.PayloadParseResponse, errors.New(response.PayloadParseResponse.Message)
	}
	if user == nil {
		return response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, errors.New(response.DBConnResponse.Message)
	}
	var permissionIds []uint
	for _, permission := range groupInfo.PermissionList {
		permissionIds = append(permissionIds, permission.ID)
	}
	var permissionMap = make(map[uint]mall_sys_permission.MallSysPermission)
	err = session.Transaction(func(tx *gorm.DB) error {
		var group = mall_sys_group.MallSysGroup{
			UserId:    &user.ID,
			ShopId:    &groupInfo.ShopId,
			GroupName: groupInfo.GroupName,
			GroupDesc: groupInfo.GroupDesc,
		}
		err = group.Create(tx)
		if err != nil {
			return err
		}
		var permissionList mall_sys_permission.MallSysPermissionList
		if len(permissionIds) != 0 {
			err := permissionList.SelectPermissionByIds(tx, permissionIds)
			if err != nil {
				return err
			}
		}
		for _, permission := range permissionList {
			permissionMap[permission.ID] = permission
		}

		var relationList mall_sys_group_permission_relation.MallSysGroupPermissionRelationList
		for _, permission := range groupInfo.PermissionList {
			if _, ok := permissionMap[permission.ID]; !ok {
				gtx.Logger.Errorf("[GroupCreateSrv] user: %s, permission info not found: %d", user.Username, permission.ID)
				continue
			}
			relationList = append(relationList, mall_sys_group_permission_relation.MallSysGroupPermissionRelation{
				GroupId:      &group.ID,
				PermissionId: &permission.ID,
			})
		}
		err = relationList.Create(tx)
		if err != nil {
			gtx.Logger.Errorf("[GroupCreateSrv] user: %s, group permission relation create err: %s", user.Username, err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		gtx.Logger.Errorf("[GroupCreateSrv] user: %s, create group err: %s", user.Username, err.Error())
		return response.SQLExecResponse, err
	}
	return response.SuccessResponse, nil
}

//
// UserRelatedGroupListSrv
// @Description: 用户关联权限组列表接口
// @Document:
// @param user
// @param gtx
// @return *response.Response
// @return error
//
func UserRelatedGroupListSrv(user *http_dto.UserDTO, gtx *ctx.GinContext) (*response.Response, error) {
	if user == nil {
		return response.UserNotLoginResponse, errors.New(response.UserNotLoginResponse.Message)
	}
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, errors.New(response.DBConnResponse.Message)
	}
	var groupList mall_sys_group.MallSysGroupList

	err = session.Transaction(func(tx *gorm.DB) error {
		err = groupList.SelectGroupByUserId(tx, user.ID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		gtx.Logger.Errorf("[UserRelatedGroupListSrv] user: %s, get err: %+v", user.Username, err)
		return response.SQLExecResponse, errors.New(response.SQLExecResponse.Message)
	}
	for _, g := range groupList {
		user.GroupList = append(user.GroupList, http_dto.GroupDTO{
			ID:        g.ID,
			GroupName: g.GroupName,
			GroupDesc: g.GroupDesc,
			ShopId:    *g.ShopId,
		})
	}
	return response.SuccessResponse, nil
}
