package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	cli "github.com/urfave/cli/v2"
)

type App struct {
	Name        string
	Usage       string
	PrepareFunc cli.BeforeFunc
	RunFunc     func(map[string]interface{}, context.Context, chan error, *sync.WaitGroup) error

	implement cli.App

	// exit gracefully
	wg      sync.WaitGroup
	errChan chan error
}

var commonFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "print help message",
	},
	&cli.BoolFlag{
		Name:    "info",
		Aliases: []string{"i"},
		Usage:   "print build info",
	},
	&cli.StringFlag{
		Name:  "control",
		Usage: "send command to control the running service [start | stop | restart]",
	},
	&cli.StringFlag{
		Name:  "pid-dir",
		Usage: "where to put the pid file",
		Value: ".",
	},
}

func (a *App) StartService(customFlags ...cli.Flag) error {
	a.errChan = make(chan error, 1)

	flags := append(commonFlags, customFlags...)
	sort.Sort(cli.FlagsByName(flags))

	a.implement = cli.App{
		Name:        a.Name,
		Usage:       a.Usage,
		Version:     fmt.Sprintf("%s - %s [%s]", Version, Commit, Branch),
		HideHelp:    true, // 覆盖 cli 包原生的 help 命令行参数
		HideVersion: true, // 覆盖 cli 包原生的 version 命令行参数
		Flags:       flags,
		Before: func(ctx *cli.Context) error {
			if a.PrepareFunc != nil {
				return a.PrepareFunc(ctx)
			}
			return nil
		},
		Action: a.runFuncWrapper(),
	}

	if err := a.implement.Run(os.Args); err != nil {
		logger.Error("run service fail with err [%v]", err)
		return err
	}

	return nil
}

func (a *App) runFuncWrapper() cli.ActionFunc {
	return func(c *cli.Context) error {
		// 在程序退出时确保打印出所有的异步日志
		defer logger.Close()

		if c.Bool("help") {
			return cli.ShowAppHelp(c)
		}

		if c.Bool("info") {
			printBuildInfo()
			return nil
		}

		pidFileDir := c.String("pid-dir")
		if strings.TrimSpace(pidFileDir) == "" {
			pidFileDir, _ = os.Getwd()
		}

		// 处理启停控制指令 (-c start | stop | restart)
		controlCmd := c.String("control")
		if controlCmd != "" {
			switch controlCmd {
			case "stop":
				controlStopHandler(pidFileDir, a.Name)
				return nil
			case "start":
				controlStartHandler(pidFileDir, a.Name)
			case "restart":
				controlRestartHandler(pidFileDir, a.Name)
			default:
				fmt.Printf("[Error] unsupported control command [%s]\n", controlCmd)
				return cli.ShowAppHelp(c)
			}
		}

		// 处理 pid 文件
		if _, isDaemon := os.LookupEnv("DAEMON"); isDaemon {
			if controlCmd == "start" || controlCmd == "restart" {
				logger.Info("DAEMON process [%d] start running in the background", os.Getpid())
				if err := writeServicePid(pidFileDir, a.Name); err != nil {
					logger.Error("write pid to file fail with err [%v]", err)
					fmt.Printf("[Error] write pid to file fail with err [%v]\n", err)
					syscall.Exit(1)
				}
				defer deletePidFile(pidFileDir, a.Name)
			}
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// 运行主逻辑
		a.wg.Add(1)
		// 获取所有的 flag 透传出去给用户 run 方法

		var flagMap = make(map[string]interface{})
		// 如果命令行没指定的话，这里 c.FlagNames () 拿不到任何 flags
		// for _, name := range c.FlagNames() {
		// 	flagMap[name] = c.Generic(name)
		// }
		// 换成 a.implement.Flags 获取参数
		for _, cliFlag := range a.implement.Flags {
			flagName := cliFlag.Names()[0]
			flagMap[flagName] = c.Generic(flagName)
		}

		go func() {
			defer func() {
				defer a.wg.Done()
				if err := recover(); err != nil {
					// 避免因为 logger 挂了导致无法打印堆栈信息
					fmt.Printf("panic with err [%v], stack:\n%s\n", err, string(debug.Stack()))
					logger.Error("panic with err [%v], stack:\n%s", err, string(debug.Stack()))
					a.errChan <- fmt.Errorf("panic with err [%v]", err)
				}
			}()

			logger.Info("app [%s] start", a.Name)
			err := a.RunFunc(flagMap, ctx, a.errChan, &a.wg)
			a.errChan <- err
			if err != nil {
				logger.Error("app [%s] quit with err: %v", a.Name, err)
			} else {
				logger.Info("app [%s] quit successfully!", a.Name)
			}
		}()

		// 优雅退出
		a.wait(cancel)

		return nil
	}
}

func (a *App) wait(cancel context.CancelFunc) {
	// 一直阻塞直到接收到信号或者抛出错误
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case sig := <-sigChan:
		logger.Info("receive signal [%d], process [%d] quit", sig, os.Getpid())
	case err := <-a.errChan:
		if err != nil {
			logger.Error("service exit with err [%s]", err)
		}
	}

	// 等待两秒以实现优雅退出，如果还是不能退出则强制退出
	cancel()
	waitCtx, waitCancel := context.WithTimeout(context.Background(), time.Second*2)
	defer waitCancel()

	isWgDone := make(chan struct{})
	go func() {
		a.wg.Wait()
		close(isWgDone)
	}()

	select {
	case <-waitCtx.Done():
		fmt.Println("force quit!")
		logger.Error("force quit!")
	case <-isWgDone:
		return
	}
}
