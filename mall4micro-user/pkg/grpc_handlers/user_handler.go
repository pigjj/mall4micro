package grpc_handlers

import (
	"context"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/ctx"
	commonGrpcDto "github.com/jianghaibo12138/mall4micro/mall4micro-common/grpc_dto/mall4micro-common/protos"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user"
	userGrpcDto "github.com/jianghaibo12138/mall4micro/mall4micro-user/grpc_dto/mall4micro-user/protos"
	"gorm.io/gorm"
)

type RpcUserInfoByUsernameSrvServer struct {
	userGrpcDto.UnimplementedRpcUserInfoByUsernameSrvServer
}

//
// CallRpcUserInfoByUsernameSrv
// @Description:  通过用户名获取用户信息grpc接口
// @Document:
// @receiver s
// @param c
// @param in
// @return *grpc_dto.RpcAuthenticateReply
// @return error
//
func (s RpcUserInfoByUsernameSrvServer) CallRpcUserInfoByUsernameSrv(c context.Context, in *userGrpcDto.RpcUserInfoByUsernameRequest) (*userGrpcDto.RpcUserInfoReply, error) {
	var out = &userGrpcDto.RpcUserInfoReply{
		Reply: &commonGrpcDto.RpcReply{
			Code:    int64(response.SuccessResponse.Code),
			Message: response.SuccessResponse.Message,
		},
	}
	gtx := ctx.GetGrpcCtx(c)
	gtx.Logger.Infof("[CallRpcUserInfoByUsernameSrv] in: %+v", in)
	var user mall_sys_user.MallSysUser
	session, err := conn.Conn()
	if err != nil {
		out.Reply.Code = int64(response.DBConnResponse.Code)
		out.Reply.Message = response.DBConnResponse.Message
		out.Reply.Data = err.Error()
		return out, nil
	}
	err = session.Transaction(func(tx *gorm.DB) error {
		err = user.SelectUserByUsername(tx, in.Username)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		gtx.Logger.Errorf("[CallRpcUserInfoByUsernameSrv] find user err: %s, username: %s", err.Error(), in.Username)
		out.Reply.Code = int64(response.SQLExecResponse.Code)
		out.Reply.Message = response.SQLExecResponse.Message
		out.Reply.Data = err.Error()
		return out, nil
	}
	out.Id = uint32(user.ID)
	out.Username = user.Username
	out.Password = user.Password
	out.Email = user.Email
	out.Mobile = user.Mobile
	out.Status = int64(*user.Status)
	out.SaltStr = user.SaltStr
	return out, nil
}
