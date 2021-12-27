package main

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/dao/mall_category"
)

func init() {
	conf.ReloadConf("mall4micro-product")
}

func main() {
	mall_category.MallShopCategory{}.Migrate()
}
