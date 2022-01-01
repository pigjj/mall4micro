package routers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	authGrpcDto "github.com/pigjj/mall4micro/mall4micro-auth/grpc_dto/mall4micro-auth/protos"
	"github.com/pigjj/mall4micro/mall4micro-auth/pkg/grpc_handlers"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"google.golang.org/grpc"
)

func InitGrpcRouter() *grpc.Server {
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		ctx.NewGrpcContext(logger),
	))
	authGrpcDto.RegisterRpcAuthenticateSrvServer(s, grpc_handlers.RpcAuthenticateSrvServer{})
	return s
}
