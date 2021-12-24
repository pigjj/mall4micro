package service

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/services/discovery"
)

func MysqlRegister() {
	serviceRegisterDto := &dto.ConsulServiceDTO{
		ID:      "73816d00-6954-43ca-ac64-ce8f1558a7ff",
		Name:    "mysql",
		Tags:    []string{"mysql"},
		Address: "192.168.0.115",
		Port:    3306,
		ServiceCheck: dto.Check{
			DeregisterCriticalServiceAfter: "30s",
			Args:                           []string{"/usr/bin/mysql", "--user=root", "--host=192.168.0.115", "--port=3306", "--user=root", "", "--password=123456", "--execute=SHOW DATABASES;"},
			Interval:                       "2s",
			Timeout:                        "2s",
		},
		ServiceWeights: dto.Weights{
			Passing: 10,
			Warning: 1,
		},
	}
	ar := discovery.NewAliveRegister(serviceRegisterDto)
	err := ar.Register()
	if err != nil {
		panic(err)
	}
}
