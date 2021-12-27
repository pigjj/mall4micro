package mall_product

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	cm "github.com/jianghaibo12138/mall4micro/mall4micro-common/models"
	"time"
)

//
// MallProduct
// @Description: 商户产品
//
type MallProduct struct {
	cm.MallBase
	ShopId      *uint      `gorm:"column:shop_id;comment:店铺ID;not null" json:"shop_id"`
	ProductName string     `gorm:"column:product_name;size:50;comment:产品类目名称;not null" json:"product_name"`
	OriPrice    *float64   `gorm:"column:ori_price;comment:产品原价;default:0" json:"ori_price"`
	Price       *float64   `gorm:"column:price;comment:产品现价;default:0" json:"price"`
	Brief       string     `gorm:"column:brief;size:500;comment:简要描述,卖点等;default:null" json:"brief"`
	Content     string     `gorm:"column:content;size:100;comment:详细描述;default:null" json:"content"`
	Pic         string     `gorm:"column:pic;size:100;comment:商品主图;default:null" json:"pic"`
	Status      *int       `gorm:"column:status;default:1;comment:默认是1，表示正常状态, 0下架;default:null" json:"status"`
	CategoryId  *int       `gorm:"column:category_id;comment:产品分类;default:0" json:"category_id"`
	SoldNum     *int       `gorm:"column:sold_num;default:0;comment:总销量;default:0" json:"sold_num"`
	TotalStocks *int       `gorm:"column:total_stocks;default:0;comment:总库存;default:0" json:"total_stocks"`
	PutOnTime   *time.Time `gorm:"column:put_on_time;comment:上架时间;default:null" json:"put_on_time"`
}

type MallProductList []MallProduct

const (
	TableMallProduct = "mall_product"
)

var (
	StatusOffline = 0
	StatusOnline  = 1
)

func (MallProduct) TableName() string {
	return TableMallProduct
}

func (MallProduct) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallProduct{}) {
		_ = session.AutoMigrate(&MallProduct{})
	} else {
		_ = session.Migrator().CreateTable(&MallProduct{})
	}
}
