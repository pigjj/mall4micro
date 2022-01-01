package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-common/log"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	userGrpcDto "github.com/pigjj/mall4micro/mall4micro-user/grpc_dto/mall4micro-user/protos"
	"google.golang.org/grpc"
	"time"
)

//
// userInfoByUsernameService
// @Description: 通过username获取userinfo接口
// @Document:
// @param username
// @param logger
// @return *userGrpcDto.RpcUserInfoReply
// @return error
//
func userInfoByUsernameService(username string, logger *log.ZapLogger) (*userGrpcDto.RpcUserInfoReply, error) {
	c, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.Settings.GrpcClient.GrpcUser.Host, conf.Settings.GrpcClient.GrpcUser.Port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := userGrpcDto.NewRpcUserInfoByUsernameSrvClient(c)
	request := userGrpcDto.RpcUserInfoByUsernameRequest{
		Username: username,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := client.CallRpcUserInfoByUsernameSrv(ctx, &request)
	if err != nil {
		logger.Errorf("[userInfoByUsernameService] call gRpc auth service err: %s", err.Error())
		return nil, err
	}
	if res.Reply.Code != int64(response.SuccessResponse.Code) {
		return nil, errors.New(response.RPCExecResponse.Message)
	}
	return res, nil
}
