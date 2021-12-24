package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/services/discovery"
	"github.com/jianghaibo12138/mall4micro/mall4micro-product/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf.ReloadConf(routers.MicroServiceName)
	if !conf.Settings.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	logger := log.InitZapLogger(routers.MicroServiceName, conf.Settings.Server.Debug)
	r := routers.InitRouter()

	serverUrl := fmt.Sprintf("%s:%d", conf.Settings.Server.Host, conf.Settings.Server.Port)
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: r,
	}

	_, err := discovery.ServiceRegister()
	if err != nil {
		logger.Fatalf("[mall4micro-product] Service Discovery failed: %+v", err)
		return
	}

	logger.Infof("[mall4micro-product] Server start on: %s", serverUrl)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("[mall4micro-product] Server listen: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("[mall4micro-product] Start Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("[mall4micro-product] Server forced to shutdown: %s", err.Error())
	}
	logger.Info("[mall4micro-product] Shutdown server...")
}