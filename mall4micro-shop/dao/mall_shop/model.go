package mall_shop

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	cm "github.com/pigjj/mall4micro/mall4micro-common/models"
)

//
// MallShop
// @Description: 商户
//
type MallShop struct {
	cm.MallBase
	UserId   *uint  `gorm:"column:user_id;comment:所有者ID;not null" json:"user_id"`
	ShopName string `gorm:"column:shop_name;size:50;comment:商铺名称;default:null" json:"shop_name"`
	ShopDesc string `gorm:"column:shop_desc;size:50;comment:商铺描述;default:null" json:"shop_desc"`
	ShopPic  string `gorm:"column:shop_pic;size:100;comment:商铺的显示图片;default:null" json:"shop_pic"`
	Status   *int   `gorm:"column:status;default:1;comment:默认是1，表示正常状态,0为下线状态;default:null" json:"status"`
}

type MallShopList []MallShop

const (
	TableMallShop = "mall_shop"
)

var (
	StatusOffline = 0
	StatusOnline  = 1
)

func (MallShop) TableName() string {
	return TableMallShop
}

func (MallShop) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallShop{}) {
		_ = session.AutoMigrate(&MallShop{})
	} else {
		_ = session.Migrator().CreateTable(&MallShop{})
	}
}
