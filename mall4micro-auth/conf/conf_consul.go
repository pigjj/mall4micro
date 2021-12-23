package conf

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/http_client"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

//
// AuthConf
// @Description: Auth服务配置接收
//
type AuthConf struct {
	Server struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Debug bool   `yaml:"debug"`
	} `yaml:"server"`
}

//
// downloadConf
// @Description: 从consul下载配置
// @Document:
// @receiver af
// @param client
// @return error
//
func (af *AuthConf) downloadConf(client *http_client.Client) error {
	response, err := client.Request(nil)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var kvDTOs []dto.ConsulKvDTO
	err = json.Unmarshal(bytes, &kvDTOs)
	if err != nil {
		return err
	}
	if len(kvDTOs) != 1 {
		return errors.New("kv invalid")
	}
	data, err := base64.StdEncoding.DecodeString(kvDTOs[0].Value)
	err = yaml.Unmarshal(data, af)
	if err != nil {
		return err
	}
	return nil
}

//
// LoadConf
// @Description: 从consul中获取最新的配置信息
// @Document:
// @receiver af
// @return error
//
func (af *AuthConf) LoadConf() error {
	client := http_client.NewHttpClient(KvGetMethod, fmt.Sprintf("%s%s/%s", LocalSettings.Conf.Consul.Url, KvGetUrl, LocalSettings.Conf.Consul.FileName), "application/json", nil)
	return af.downloadConf(client)
}
