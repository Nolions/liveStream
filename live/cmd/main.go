package main

import (
	"context"
	"flag"
	"fmt"
	"live/conf"
	"live/internal/api"
	log "live/internal/helper"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var confPath string

func main() {
	flag.StringVar(&confPath, "c", "conf.yaml", "default conf path")
	flag.Parse()

	ctx := context.Background()

	// loading conf
	config, err := conf.New(confPath)
	if err != nil {
		panic(fmt.Sprintf("new conf err: %v", err))
	}

	// logger設定
	log.InitLogger(config.App.Debug)

	log.Infof("Config loaded: %+v", config)
	log.Info("App Info", zap.String("Project", config.App.Name))

	serv := api.New(ctx, &config.App)
	serv.Run()

	shutdown(config)
}

func shutdown(config *conf.Conf) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	s := <-quit
	log.Debugf("get a signal %s. %s Server is shutting down ...", s.String(), config.App.Name)

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Infof("%s Server is exiting...", config.App.Name)
	log.CloseLogger()
}
