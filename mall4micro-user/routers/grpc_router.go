package routers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	userGrpcDto "github.com/pigjj/mall4micro/mall4micro-user/grpc_dto/mall4micro-user/protos"
	"github.com/pigjj/mall4micro/mall4micro-user/pkg/grpc_handlers"
	"google.golang.org/grpc"
)

func InitGrpcRouter() *grpc.Server {
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		ctx.NewGrpcContext(logger),
	))
	userGrpcDto.RegisterRpcUserInfoByUsernameSrvServer(s, grpc_handlers.RpcUserInfoByUsernameSrvServer{})
	return s
}
