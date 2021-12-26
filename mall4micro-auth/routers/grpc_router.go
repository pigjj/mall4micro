package routers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/grpc_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/pkg/grpc_handlers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"google.golang.org/grpc"
)

func InitGrpcRouter() *grpc.Server {
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		ctx.NewGrpcContext(logger),
	))
	grpc_dto.RegisterRpcAuthenticateSrvServer(s, grpc_handlers.RpcAuthenticateSrvServer{})
	return s
}
