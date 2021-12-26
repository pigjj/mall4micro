package ctx

import (
	"context"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"google.golang.org/grpc"
)

type GrpcContext struct {
	context.Context
	Logger *log.ZapLogger
}

const cKey = "GRPC_CTX"

func NewGrpcContext(logger *log.ZapLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		commonCtx := newContext(ctx, logger)
		return handler(commonCtx, req)
	}
}

func newContext(c context.Context, logger *log.ZapLogger) *GrpcContext {
	gRpcCtx := &GrpcContext{}
	gRpcCtx.Context = storeContext(c, gRpcCtx)
	gRpcCtx.Logger = logger
	return gRpcCtx
}

// 拦截器之间直接只能通过context.Context传递, 所以需要将自定义context存到go的context里向下传
func storeContext(c context.Context, ctx *GrpcContext) context.Context {
	return context.WithValue(c, cKey, ctx)
}

func GetGrpcCtx(c context.Context) *GrpcContext {
	return c.Value(cKey).(*GrpcContext)
}
