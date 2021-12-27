package mall_sys_group_permission_relation

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	cm "github.com/jianghaibo12138/mall4micro/mall4micro-common/models"
)

//
// MallSysGroupPermissionRelation
// @Description: 商城组权限关系
//
type MallSysGroupPermissionRelation struct {
	cm.MallBase
	GroupId      *uint `gorm:"column:group_id;comment:组ID;not null" json:"group_id"`
	PermissionId *uint `gorm:"column:permission_id;comment:权限ID;not null" json:"permission_id"`
}

type MallSysGroupPermissionRelationList []MallSysGroupPermissionRelation

const (
	TableMallSysGroupPermissionRelation = "mall_sys_group_permission_relation"
)

func (MallSysGroupPermissionRelation) TableName() string {
	return TableMallSysGroupPermissionRelation
}

func (MallSysGroupPermissionRelation) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallSysGroupPermissionRelation{}) {
		_ = session.AutoMigrate(&MallSysGroupPermissionRelation{})
	} else {
		_ = session.Migrator().CreateTable(&MallSysGroupPermissionRelation{})
	}
}
