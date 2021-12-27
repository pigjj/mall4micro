package mall_category

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	cm "github.com/jianghaibo12138/mall4micro/mall4micro-common/models"
)

//
// MallShopCategory
// @Description: 商户产品类目
//
type MallShopCategory struct {
	cm.MallBase
	ShopId       *int   `gorm:"column:shop_id;comment:店铺ID;not null" json:"shop_id"`
	ParentId     *int   `gorm:"column:parent_id;comment:父节点;not null" json:"parent_id"`
	CategoryName string `gorm:"column:category_name;size:50;comment:产品类目名称;default:null" json:"category_name"`
	Icon         string `gorm:"column:icon;size:50;comment:类目图标;default:null" json:"icon"`
	Pic          *int   `gorm:"column:pic;comment:类目的显示图片;default:null" json:"pic"`
	Status       *int   `gorm:"column:status;default:1;comment:默认是1，表示正常状态,0为下线状态;default:null" json:"status"`
}

const (
	TableMallShopCategory = "mall_shop_category"
)

var (
	StatusOffline = 0
	StatusOnline  = 1
)

func (MallShopCategory) TableName() string {
	return TableMallShopCategory
}

func (MallShopCategory) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallShopCategory{}) {
		_ = session.AutoMigrate(&MallShopCategory{})
	} else {
		_ = session.Migrator().CreateTable(&MallShopCategory{})
	}
}
