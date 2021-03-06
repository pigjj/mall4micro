package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-common/log"
	"github.com/pigjj/mall4micro/mall4micro-common/services/discovery"
	"github.com/pigjj/mall4micro/mall4micro-shop/constant"
	"github.com/pigjj/mall4micro/mall4micro-shop/routers"
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
	r := routers.InitRouter()

	serverUrl := fmt.Sprintf("%s:%d", conf.Settings.HttpServer.Host, conf.Settings.HttpServer.Port)
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: r,
	}
	if conf.Settings.HttpServer.AutoRegister {
		_, err := discovery.ServiceRegister()
		if err != nil {
			logger.Fatalf("[mall4micro-shop] Service Discovery failed: %+v", err)
			return
		}
	}

	logger.Infof("[mall4micro-shop] HttpServer start on: %s", serverUrl)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("[mall4micro-shop] HttpServer listen: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("[mall4micro-shop] Start Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("[mall4micro-shop] HttpServer forced to shutdown: %s", err.Error())
	}
	logger.Info("[mall4micro-shop] Shutdown server...")
}
