package main

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-shop/dao/mall_shop"
)

func init() {
	conf.ReloadConf("mall4micro-shop")
}

func main() {
	mall_shop.MallShop{}.Migrate()
}
