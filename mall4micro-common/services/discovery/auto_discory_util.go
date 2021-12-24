package discovery

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
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
		ID:             conf.Settings.Server.ServerId,
		Name:           conf.Settings.Server.ServerName,
		Tags:           conf.Settings.Server.ServerTags,
		Address:        conf.Settings.Server.Address,
		Port:           conf.Settings.Server.Port,
		ServiceCheck:   conf.Settings.Server.ServiceCheck,
		ServiceWeights: conf.Settings.Server.ServiceWeights,
	}
	ar := NewAliveRegister(serviceRegisterDto)
	return ar, ar.Register()
}
