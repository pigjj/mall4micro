package conf

import "github.com/snownd/cake"

type CakeRequestConf struct {
	cake.RequestConfig
	Key string `param:"key"`
}

type AuthConf struct {
}
