package main

import (
	"context"
	"fmt"
	"sync"
	"syscall"
	"time"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/db"
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
			Value: "conf/config-dev.toml",
		}, &cli.StringFlag{
			Name:  "log-conf",
			Usage: "log config file path",
			Value: "conf/logger.json",
		},
	)
}

func run(flags map[string]interface{}, ctx context.Context, errChan chan error, appWg *sync.WaitGroup) error {
	// 初始化日志
	var loggerConfPath = fmt.Sprintf("%v", flags["log-conf"])
	if err := logger.InitLogger(loggerConfPath); err != nil {
		errmsg := fmt.Sprintf("init logger [%s] fail with err [%v]", loggerConfPath, err)
		fmt.Printf("%s\n", errmsg)
		return fmt.Errorf(errmsg)
	}

	// 初始化配置文件
	var umsConfPath = fmt.Sprintf("%v", flags["conf"])
	if err := config.ParseConfig(umsConfPath); err != nil {
		logger.Error("parse ums config [%s] fail with err [%v]", umsConfPath, err)
		return err
	}

	// 初始化 redis 客户端
	db.InitRedisClient()

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

	// 启动 5 秒后自动退出
	go func() {
		for i := 0; i < 5; i++ {
			logger.Info("ums server is running...")
			time.Sleep(time.Second * 1)
		}
		logger.Info("master service update, going to quit and restart!")
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	return grpcserver.NewServer().Start(ctx, config.GlobalUmsConfig.GrpcServer.Port)
}
