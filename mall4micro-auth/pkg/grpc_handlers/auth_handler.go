package grpc_handlers

import (
	"context"
	authGrpcDto "github.com/jianghaibo12138/mall4micro/mall4micro-auth/grpc_dto/mall4micro-auth/protos"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	commonGrpcDto "github.com/jianghaibo12138/mall4micro/mall4micro-common/grpc_dto/mall4micro-common/protos"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
)

type RpcAuthenticateSrvServer struct {
	authGrpcDto.UnimplementedRpcAuthenticateSrvServer
}

//
// CallRpcAuthenticateSrv
// @Description: Token认证gRpc接口
// @Document:
// @receiver s
// @param c
// @param in
// @return *authGrpcDto.RpcAuthenticateReply
// @return error
//
func (s RpcAuthenticateSrvServer) CallRpcAuthenticateSrv(c context.Context, in *authGrpcDto.RpcAuthenticateRequest) (*authGrpcDto.RpcAuthenticateReply, error) {
	var out = &authGrpcDto.RpcAuthenticateReply{
		Reply: &commonGrpcDto.RpcReply{
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
		out.Reply.Data = err.Error()
		return out, nil
	}
	out.Id = uint32(t.ID)
	out.Username = t.Username
	out.Email = t.Email
	out.Mobile = t.Mobile
	out.Status = int64(t.Status)
	return out, nil
}
