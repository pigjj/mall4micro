package mall_sys_user_group_relation

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	cm "github.com/pigjj/mall4micro/mall4micro-common/models"
)

//
// MallSysUserGroupRelation
// @Description: 商城用户组关系
//
type MallSysUserGroupRelation struct {
	cm.MallBase
	GroupId *uint `gorm:"column:group_id;comment:组ID;not null" json:"group_id"`
	Userid  *uint `gorm:"column:user_id;comment:用户ID;not null" json:"user_id"`
}

type MallSysUserGroupRelationList []MallSysUserGroupRelation

const (
	TableMallSysUserGroupRelation = "mall_sys_user_group_relation"
)

func (MallSysUserGroupRelation) TableName() string {
	return TableMallSysUserGroupRelation
}

func (MallSysUserGroupRelation) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallSysUserGroupRelation{}) {
		_ = session.AutoMigrate(&MallSysUserGroupRelation{})
	} else {
		_ = session.Migrator().CreateTable(&MallSysUserGroupRelation{})
	}
}
