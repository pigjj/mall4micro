package mall_sys_group

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	cm "github.com/jianghaibo12138/mall4micro/mall4micro-common/models"
)

//
// MallSysGroup
// @Description: 商城用户组
//
type MallSysGroup struct {
	cm.MallBase
	UserId    *uint  `gorm:"column:user_id;comment:用户ID;not null" json:"user_id"`
	ShopId    *uint  `gorm:"column:shop_id;comment:商铺ID;not null" json:"shop_id"`
	GroupName string `gorm:"column:group_name;size:50;comment:组名称;not null" json:"group_name"`
	GroupDesc string `gorm:"column:group_desc;size:100;comment:组描述;not null" json:"group_desc"`
}

type MallSysGroupList []MallSysGroup

const (
	TableMallSysGroup = "mall_sys_group"
)

func (MallSysGroup) TableName() string {
	return TableMallSysGroup
}

func (MallSysGroup) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallSysGroup{}) {
		_ = session.AutoMigrate(&MallSysGroup{})
	} else {
		_ = session.Migrator().CreateTable(&MallSysGroup{})
	}
}
