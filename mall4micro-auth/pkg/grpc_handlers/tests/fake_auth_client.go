package main

import (
	"context"
	"fmt"
	"github.com/pigjj/mall4micro/mall4micro-auth/grpc_dto"
	"google.golang.org/grpc"
	"time"
)

func Auth() {
	conn, err := grpc.Dial("0.0.0.0:8090", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := grpc_dto.NewRpcAuthenticateSrvClient(conn)
	request := grpc_dto.RpcAuthenticateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTY0MTEyNTA3NiwiaWF0IjoxNjQwNTIwMjc2LCJpc3MiOiJtYWxsNG1pY3JvIiwibW9iaWxlIjoiIiwibmJmIjoxNDQ0NDc4NDAwLCJzdGF0dXMiOjAsInVzZXJuYW1lIjoiYWRtaW4ifQ.IsYxbKIfSYplvZWbqww3x5L_apk3ywC35Lt2qIdMaB4",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := client.CallRpcAuthenticateSrv(ctx, &request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func main() {
	Auth()
}
