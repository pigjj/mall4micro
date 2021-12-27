package main

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/dao/mall_category"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/dao/mall_product"
)

func init() {
	conf.ReloadConf("mall4micro-product")
}

func main() {
	mall_category.MallCategory{}.Migrate()
	mall_product.MallProduct{}.Migrate()
}
