package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/routers"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/services/cron_service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//
// ServiceRegister
// @Description: 服务自注册
// @Document:
// @return *cron_service.AliveRegister
// @return error
//
func ServiceRegister() (*cron_service.AliveRegister, error) {
	serviceRegisterDto := &dto.ConsulServiceDTO{
		ID:             conf.Settings.Server.ServerId,
		Name:           conf.Settings.Server.ServerName,
		Tags:           conf.Settings.Server.ServerTags,
		Address:        conf.Settings.Server.Address,
		Port:           conf.Settings.Server.Port,
		ServiceCheck:   conf.Settings.Server.ServiceCheck,
		ServiceWeights: conf.Settings.Server.ServiceWeights,
	}
	ar := cron_service.NewAliveRegister(serviceRegisterDto)
	return ar, ar.Register()
}

func main() {
	if !conf.Settings.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	logger := log.InitZapLogger("mall4micro-auth", conf.Settings.Server.Debug)
	r := routers.InitRouter()

	serverUrl := fmt.Sprintf("%s:%d", conf.Settings.Server.Host, conf.Settings.Server.Port)
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: r,
	}

	_, err := ServiceRegister()
	if err != nil {
		logger.Fatalf("[mall4micro-auth] Service Discovery failed: %+v", err)
		return
	}

	logger.Infof("[mall4micro-auth] Server start on: %s", serverUrl)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("[mall4micro-auth] Server listen: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("[mall4micro-auth] Start Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("[mall4micro-auth] Server forced to shutdown: %s", err.Error())
	}
	logger.Info("[mall4micro-auth] Shutdown server...")
}
