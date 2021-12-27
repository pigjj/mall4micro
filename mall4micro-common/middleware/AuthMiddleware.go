package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	authGrpcDto "github.com/jianghaibo12138/mall4micro/mall4micro-auth/grpc_dto/mall4micro-auth/protos"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

var (
	ErrAuthorizedFailed = errors.New("authorized failed")
)

func authorized(token string, logger *log.ZapLogger) (*authGrpcDto.RpcAuthenticateReply, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.Settings.GrpcClient.GrpcAuth.Host, conf.Settings.GrpcClient.GrpcAuth.Port), grpc.WithInsecure())
	if err != nil {
		logger.Errorf("[authorized] connect gRpc server err: %s", err.Error())
		return nil, err
	}
	client := authGrpcDto.NewRpcAuthenticateSrvClient(conn)
	request := authGrpcDto.RpcAuthenticateRequest{
		Token: token,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := client.CallRpcAuthenticateSrv(ctx, &request)
	if err != nil {
		logger.Errorf("[authorized] call gRpc auth service err: %s", err.Error())
		return nil, err
	}
	if res.Reply.Code != int64(response.SuccessResponse.Code) {
		return nil, ErrAuthorizedFailed
	}
	return res, nil
}

func AuthMiddleWare(logger *log.ZapLogger) gin.HandlerFunc {
	// 当客户端有请求来之后, 先执行这个函数
	return func(c *gin.Context) {
		if utils.StringContained(conf.Settings.Authorized.IgnoreUrls, c.Request.URL.Path) != -1 {
			c.Next()
		} else {
			token := c.GetHeader("Authorization")
			userInfo, err := authorized(token, logger)
			if err != nil {
				logger.Errorf("[AuthMiddleWare] auth user err: %s", err.Error())
				c.JSON(http.StatusOK, response.UserNotLoginResponse)
				c.Abort()
				return
			}
			c.Set("user", http_dto.HttpAuthenticateDTO{
				Username: userInfo.Username,
				Email:    userInfo.Email,
				Mobile:   userInfo.Mobile,
				Status:   int(userInfo.Status),
			})
			// 执行完对应的回调函数之后, 继续回到这个地方进行执行(但是响应还没有返回给客户端)
			c.Next()
		}
	}
}
