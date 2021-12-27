package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/services/discovery"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/constant"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/routers"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf.ReloadConf(constant.MicroServiceName)
	if !conf.Settings.HttpServer.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	logger := log.InitZapLogger(constant.MicroServiceName, conf.Settings.HttpServer.Debug)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Settings.GrpcServer.Host, conf.Settings.GrpcServer.Port))
	if err != nil {
		logger.Fatalf("[mall4micro-user] Failed to listen: %v", err)
		return
	}
	go func() {
		s := routers.InitGrpcRouter()
		reflection.Register(s)
		logger.Infof("[mall4micro-user] gRpcServer start on %s:%d", conf.Settings.GrpcServer.Host, conf.Settings.GrpcServer.Port)

		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	r := routers.InitRouter()

	serverUrl := fmt.Sprintf("%s:%d", conf.Settings.HttpServer.Host, conf.Settings.HttpServer.Port)
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: r,
	}
	go func() {
		if conf.Settings.HttpServer.AutoRegister {
			_, err := discovery.ServiceRegister()
			if err != nil {
				logger.Fatalf("[mall4micro-user] Service Discovery failed: %+v", err)
				return
			}
		}
		logger.Infof("[mall4micro-user] HttpServer start on: %s", serverUrl)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("[mall4micro-user] HttpServer listen: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("[mall4micro-user] Start Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("[mall4micro-user] HttpServer forced to shutdown: %s", err.Error())
	}
	logger.Info("[mall4micro-user] Shutdown server...")
}
