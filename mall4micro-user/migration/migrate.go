package main

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_group"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_group_permission_relation"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_permission"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user_group_relation"
)

func init() {
	conf.ReloadConf("mall4micro-user")
}

func main() {
	mall_sys_user.MallSysUser{}.Migrate()
	mall_sys_group.MallSysGroup{}.Migrate()
	mall_sys_permission.MallSysPermission{}.Migrate()
	mall_sys_user_group_relation.MallSysUserGroupRelation{}.Migrate()
	mall_sys_group_permission_relation.MallSysGroupPermissionRelation{}.Migrate()
}
