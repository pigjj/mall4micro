package mall_category

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	cm "github.com/pigjj/mall4micro/mall4micro-common/models"
)

//
// MallCategory
// @Description: 商户产品类目
//
type MallCategory struct {
	cm.MallBase
	ShopId       *uint  `gorm:"column:shop_id;comment:店铺ID;not null" json:"shop_id"`
	ParentId     *uint  `gorm:"column:parent_id;comment:父节点" json:"parent_id"`
	CategoryName string `gorm:"column:category_name;size:50;comment:产品类目名称;not null" json:"category_name"`
	Icon         string `gorm:"column:icon;size:50;comment:类目图标;default:null" json:"icon"`
	Pic          string `gorm:"column:pic;size:50;comment:类目的显示图片;default:null" json:"pic"`
	Status       *int   `gorm:"column:status;default:1;comment:默认是1，表示正常状态,0为下线状态;default:null" json:"status"`
}

type MallCategoryList []MallCategory

const (
	TableMallCategory = "mall_category"
)

var (
	StatusOffline = 0
	StatusOnline  = 1
)

func (MallCategory) TableName() string {
	return TableMallCategory
}

func (MallCategory) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallCategory{}) {
		_ = session.AutoMigrate(&MallCategory{})
	} else {
		_ = session.Migrator().CreateTable(&MallCategory{})
	}
}
