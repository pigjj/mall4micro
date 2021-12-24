package main

import "github.com/jianghaibo12138/mall4micro/cmd/common-service/service"

func main() {
	service.RedisRegister()
	service.MysqlRegister()
}
