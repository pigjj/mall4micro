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
	"os"
)

type server struct {
	AutoRegister   bool        `yaml:"auto_register"`
	Host           string      `yaml:"host"`
	Debug          bool        `yaml:"debug"`
	Port           int         `yaml:"port"`
	ServerId       string      `yaml:"server_id"`
	ServerName     string      `yaml:"server_name"`
	ServerTags     []string    `yaml:"server_tags"`
	Address        string      `yaml:"address"`
	ServiceCheck   dto.Check   `yaml:"service_check"`
	ServiceWeights dto.Weights `yaml:"service_weights"`
}

//
// MicroConf
// @Description: Auth服务配置接收
//
type MicroConf struct {
	HttpServer server `yaml:"http_server"`
	GrpcServer server `yaml:"grpc_server"`
	GrpcClient struct {
		GrpcAuth    server `yaml:"grpc_auth"`
		GrpcUser    server `yaml:"grpc_user"`
		GrpcProduct server `yaml:"grpc_product"`
	} `yaml:"grpc_client"`
	Mysql struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Database    string `yaml:"database"`
		MaxIdleConn int    `yaml:"max_idle_conn"`
		MaxConn     int    `yaml:"max_conn"`
	} `yaml:"mysql"`
	Authorized struct {
		IgnoreUrls []string `yaml:"ignore_urls"`
	} `yaml:"authorized"`
}

//
// downloadConf
// @Description: 从consul下载配置
// @Document:
// @receiver af
// @param client
// @return error
//
func (af *MicroConf) downloadConf(client *http_client.Client) error {
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
func (af *MicroConf) LoadConf() error {
	client := http_client.NewHttpClient(KvGetMethod, fmt.Sprintf("%s%s/%s", LocalSettings.Conf.Consul.Url, KvGetUrl, LocalSettings.Conf.Consul.FileName), "application/json", LoadConsulAclHeader())
	return af.downloadConf(client)
}

//
// LoadConsulAclHeader
// @Description: 通过环境变量组装consul acl请求头
// @Document:
// @return map[string]string
//
func LoadConsulAclHeader() map[string]string {
	headers := make(map[string]string)
	headers["X-Consul-Token"] = os.Getenv("CONSUL_TOKEN")
	return headers
}
