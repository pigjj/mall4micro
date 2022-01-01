package discovery

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-common/dto"
)

//
// ServiceRegister
// @Description: 服务自注册
// @Document:
// @return *cron.AliveRegister
// @return error
//
func ServiceRegister() (*AliveRegister, error) {
	serviceRegisterDto := &dto.ConsulServiceDTO{
		ID:             conf.Settings.HttpServer.ServerId,
		Name:           conf.Settings.HttpServer.ServerName,
		Tags:           conf.Settings.HttpServer.ServerTags,
		Address:        conf.Settings.HttpServer.Address,
		Port:           conf.Settings.HttpServer.Port,
		ServiceCheck:   conf.Settings.HttpServer.ServiceCheck,
		ServiceWeights: conf.Settings.HttpServer.ServiceWeights,
	}
	ar := NewAliveRegister(serviceRegisterDto)
	return ar, ar.Register()
}
