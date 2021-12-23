package conf

import (
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

//
// YmlLocalConf
// @Description: 本地配置文件读取, 目的是找到远程consul连接信息
//
type YmlLocalConf struct {
	Conf struct {
		Consul struct {
			Url      string `yaml:"url"`
			FileName string `yaml:"file_name"`
		} `yaml:"consul"`
	} `yaml:"conf"`
}

//
// loadEnv
// @Description: 通过环境变量判断是dev, uat还是prod环境, 从而读取不同配置文件
// @Document:
// @return string
//
func loadEnv() string {
	var confFileName = "dev.yml"
	env := os.Getenv("mall4micro")
	switch env {
	case "PROD":
		confFileName = "conf-prod.yml"
	case "UAT":
		confFileName = "conf-uat.yml"
	default:
		confFileName = "conf-dev.yml"
	}
	return confFileName
}

//
// loadLocalConf
// @Description: 记载本地配置数据
// @Document:
//
func loadLocalConf() {
	projBasePath := utils.ProjectBasePath()
	confFileName := loadEnv()
	projBasePath = fmt.Sprintf("%s%small4micro-auth%sconf%s%s", projBasePath, utils.PathSplitFlag, utils.PathSplitFlag, utils.PathSplitFlag, confFileName)
	yamlFile, err := ioutil.ReadFile(projBasePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, localSettings)
	if err != nil {
		panic(err)
	}
}
