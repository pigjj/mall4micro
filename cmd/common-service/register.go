package main

import "github.com/pigjj/mall4micro/cmd/common-service/service"

func main() {
	service.RedisRegister()
	service.MysqlRegister()
}
