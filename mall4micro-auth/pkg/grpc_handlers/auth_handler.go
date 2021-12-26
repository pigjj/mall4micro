package grpc_handlers

import (
	"context"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/grpc_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
)

type RpcAuthenticateSrvServer struct {
	grpc_dto.UnimplementedRpcAuthenticateSrvServer
}

func (s RpcAuthenticateSrvServer) CallRpcAuthenticateSrv(c context.Context, in *grpc_dto.RpcAuthenticateRequest) (*grpc_dto.RpcAuthenticateReply, error) {
	var out = &grpc_dto.RpcAuthenticateReply{
		Reply: &grpc_dto.RpcReply{
			Code:    int64(response.SuccessResponse.Code),
			Message: response.SuccessResponse.Message,
		},
	}
	gtx := ctx.GetGrpcCtx(c)
	gtx.Logger.Infof("[CallRpcAuthenticateSrv] in: %+v", in)
	var t utils.TokenUtil
	err := t.Parse(in.Token)
	if err != nil {
		out.Reply.Code = int64(response.ParseTokenResponse.Code)
		out.Reply.Message = response.ParseTokenResponse.Message
	}
	return out, nil
}