package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/server/grpcserver"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/server/httpserver"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/app"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/urfave/cli/v2"
)

const (
	kDefaultLoggerPath  = "conf/logger.json"
	kDefaultUmsConfPath = "conf/config-dev.toml"
)

func main() {
	umsApp := app.App{
		Name:    "ums",
		Usage:   "user management system",
		RunFunc: run,
	}
	umsApp.StartService(
		&cli.StringFlag{
			Name:  "conf",
			Usage: "config file path",
			Value: "conf/config.toml",
		}, &cli.StringFlag{
			Name:  "log-conf",
			Usage: "log config file path",
			Value: "conf/logger.json",
		},
	)
}

func run(flags map[string]interface{}, ctx context.Context, errChan chan error, appWg *sync.WaitGroup) error {
	// 初始化日志
	var loggerConfPath string = kDefaultLoggerPath
	if val, exists := flags["log-conf"]; exists {
		loggerConfPath = fmt.Sprintf("%v", val)
	}
	if err := logger.InitLogger(loggerConfPath); err != nil {
		fmt.Printf("init logger [%s] fail with err [%v]\n", loggerConfPath, err)
		return fmt.Errorf("init logger [%s] fail with err [%v]", loggerConfPath, err)
	}

	// 初始化配置文件
	var umsConfPath string = kDefaultUmsConfPath
	if val, exists := flags["conf"]; exists {
		umsConfPath = fmt.Sprintf("%v", val)
	}
	if err := config.ParseConfig(umsConfPath); err != nil {
		logger.Error("parse ums config [%s] fail with err [%v]", umsConfPath, err)
		return err
	}

	// metric && pprof http service
	appWg.Add(1)
	go func() {
		defer appWg.Done()
		logger.Info("start metric && pprof server with port [%d]", config.GlobalUmsConfig.HttpServer.Port)
		if err := httpserver.Start(ctx, config.GlobalUmsConfig.HttpServer.Port); err != nil {
			errChan <- fmt.Errorf("http server closed with err [%v]", err)
		} else {
			logger.Info("http server shutdown")
		}
	}()

	// 启动 grpc 服务
	logger.Info("start grpc server with port [%d]", config.GlobalUmsConfig.GrpcServer.Port)
	return grpcserver.NewServer().Start(ctx, config.GlobalUmsConfig.GrpcServer.Port)
}
