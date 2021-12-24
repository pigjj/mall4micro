package main

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user"
)

func init() {
	conf.ReloadConf("mall4micro-user")
}

func main() {
	mall_sys_user.MallSysUser{}.Migrate()
}
