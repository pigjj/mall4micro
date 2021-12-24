package service

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/services/discovery"
)

func RedisRegister() {
	serviceRegisterDto := &dto.ConsulServiceDTO{
		ID:      "d183d513-74d2-4dec-bd8b-fb3e6cb5b834",
		Name:    "redis",
		Tags:    []string{"redis"},
		Address: "192.168.0.115",
		Port:    6377,
		ServiceCheck: dto.Check{
			DeregisterCriticalServiceAfter: "30s",
			Args:                           []string{"/usr/bin/redis-cli", "-h", "192.168.0.115", "-p", "6379", "-a", "bitorobotics", "ping"},
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
