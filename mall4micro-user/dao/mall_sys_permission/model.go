package mall_sys_permission

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	cm "github.com/jianghaibo12138/mall4micro/mall4micro-common/models"
)

//
// MallSysPermission
// @Description: 商城权限
//
type MallSysPermission struct {
	cm.MallBase
	PermissionName string `gorm:"column:permission_name;size:50;comment:权限名称;not null;index:idx_permission_name,unique" json:"permission_name"`
	PermissionDesc string `gorm:"column:permission_desc;size:100;comment:权限描述;not null" json:"permission_desc"`
}

type MallSysPermissionList []MallSysPermission

const (
	TableMallSysPermission = "mall_sys_permission"
)

func (MallSysPermission) TableName() string {
	return TableMallSysPermission
}

func (MallSysPermission) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallSysPermission{}) {
		_ = session.AutoMigrate(&MallSysPermission{})
	} else {
		_ = session.Migrator().CreateTable(&MallSysPermission{})
	}
}
