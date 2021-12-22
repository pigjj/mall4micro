package conf

import (
	"context"
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/snownd/cake"
)

const ConsulUrl = "http://127.0.0.1:8500/v1"

type ConsulConfApi struct {
	GetConsulConf    func(ctx context.Context, config *CakeRequestConf) (*dto.ConsulConfDTO, error) `url:"/kv/:key" headers:"x-request-name=users;x-request-app=cake-example"`
	UpsertConsulConf func(ctx context.Context, config *CakeRequestConf) (*dto.ConsulConfDTO, error) `url:"/kv/:key" headers:"x-request-name=users;x-request-app=cake-example"`
	DeleteConsulConf func(ctx context.Context, config *CakeRequestConf) (*dto.ConsulConfDTO, error) `url:"/kv/:key" headers:"x-request-name=users;x-request-app=cake-example"`
}

func (af *AuthConf) Reload() {
	ck := cake.New()
	defer ck.Close()
	apiItf, err := ck.Build(&ConsulConfApi{}, cake.WithBaseURL(ConsulUrl))
	if err != nil {
		panic(err)
	}
	api := apiItf.(*ConsulConfApi)
	conf, err := api.GetConsulConf(context.Background(), &CakeRequestConf{
		Key: "my-key",
	})
	fmt.Println("get conf", conf)
}
