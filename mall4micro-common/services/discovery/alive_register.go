package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-common/dto"
	"github.com/pigjj/mall4micro/mall4micro-common/http_client"
	"io/ioutil"
)

type AliveRegister struct {
	*dto.ConsulServiceDTO
}

func NewAliveRegister(d *dto.ConsulServiceDTO) *AliveRegister {
	return &AliveRegister{
		d,
	}
}

func (ar *AliveRegister) Register() error {
	client := http_client.NewHttpClient(ServiceRegisterMethod, fmt.Sprintf("%s%s", conf.LocalSettings.Conf.Consul.Url, ServiceRegisterUrl), "application/json", conf.LoadConsulAclHeader())
	return ar.uploadService(client)
}

func (ar *AliveRegister) DeRegister() error {
	client := http_client.NewHttpClient(ServiceDeRegisterMethod, fmt.Sprintf("%s%s/%s", conf.LocalSettings.Conf.Consul.Url, ServiceDeRegisterUrl, ar.ID), "application/json", conf.LoadConsulAclHeader())
	_, err := client.Request(nil)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AliveRegister) uploadService(client *http_client.Client) error {
	buf, err := json.Marshal(ar.ConsulServiceDTO)
	if err != nil {
		return err
	}
	response, err := client.Request(buf)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if len(bytes) != 0 {
		return errors.New(string(bytes))
	}
	return nil
}
