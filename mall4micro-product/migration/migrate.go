package main

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-product/dao/mall_category"
	"github.com/pigjj/mall4micro/mall4micro-product/dao/mall_product"
)

func init() {
	conf.ReloadConf("mall4micro-product")
}

func main() {
	mall_category.MallCategory{}.Migrate()
	mall_product.MallProduct{}.Migrate()
}
