package cron_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/http_client"
	"io/ioutil"
)

type AliveRegister struct {
	*dto.ConsulServiceDTO
}

//
// ConsulServiceApi
// @Description: cake连接consul处理service结构体
//
type ConsulServiceApi struct {
	UpsertConsulService func(ctx context.Context, config *dto.CakeRequestService) error `method:"PUT" url:"" headers:"Content-Type:application/json"`
}

func NewAliveRegister(d *dto.ConsulServiceDTO) *AliveRegister {
	return &AliveRegister{
		d,
	}
}

func (ar *AliveRegister) Register() error {
	client := http_client.NewHttpClient(ServiceRegisterMethod, fmt.Sprintf("%s/%s", conf.LocalSettings.Conf.Consul.Url, ServiceRegisterUrl), "application/json", nil)
	_, err := ar.uploadService(client)
	return err
}

func (ar *AliveRegister) DeRegister() {

}

func (ar *AliveRegister) uploadService(client *http_client.Client) (*[]byte, error) {
	buf, err := json.Marshal(ar.ConsulServiceDTO)
	if err != nil {
		panic(err)
	}
	response, err := client.Request(buf)
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}
