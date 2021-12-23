package conf

import (
	"context"
	"encoding/base64"
	commonDTO "github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/snownd/cake"
	"gopkg.in/yaml.v3"
)

//
// ConsulConfApi
// @Description: cake连接consul配置结构体
//
type ConsulConfApi struct {
	GetConsulConf    func(ctx context.Context, config *commonDTO.CakeRequestConf) ([]commonDTO.ConsulKvDTO, error) `method:"GET" url:"/kv/:file_name" headers:"Content-Type:application/json"`
	UpsertConsulConf func(ctx context.Context, config *commonDTO.CakeRequestConf) error                            `method:"PUT" url:"/kv/:file_name" headers:"Content-Type:application/json"`
	DeleteConsulConf func(ctx context.Context, config *commonDTO.CakeRequestConf) ([]commonDTO.ConsulKvDTO, error) `method:"DELETE" url:"/kv/:file_name" headers:"Content-Type:application/json"`
}

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
// @param fileName
// @param api
// @return error
//
func (af *AuthConf) downloadConf(fileName string, api *ConsulConfApi) error {
	kvs, err := api.GetConsulConf(context.Background(), &commonDTO.CakeRequestConf{
		FileName: fileName,
	})
	if err != nil {
		return err
	}
	if len(kvs) != 1 {
		return err
	}
	value := kvs[0].Value
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}
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
	ck := cake.New()
	defer ck.Close()
	apiItf, err := ck.Build(&ConsulConfApi{}, cake.WithBaseURL(LocalSettings.Conf.Consul.Url))
	if err != nil {
		return err
	}
	api := apiItf.(*ConsulConfApi)
	return af.downloadConf(LocalSettings.Conf.Consul.FileName, api)
}

func (af *AuthConf) UploadConf() error {
	return nil
}
