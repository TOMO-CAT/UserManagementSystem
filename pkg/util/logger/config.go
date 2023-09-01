package logger

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/TOMO-CAT/UserManagementSystem/proto/config"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	kStdioLogFile = "stdout.log"
)

// InitLogger 根据配置文件初始化日志模块
func InitLogger(loggerConfPath string) (err error) {

	if !isFileExist(loggerConfPath) {
		return fmt.Errorf("logger conf [%s] don't exist", loggerConfPath)
	}

	var confContent []byte
	if confContent, err = ioutil.ReadFile(loggerConfPath); err != nil {
		panic(err)
	}

	var confPbMsg config.LoggerConfig
	if err = protojson.Unmarshal(confContent, &confPbMsg); err != nil {
		panic(err)
	}

	return initLoggerWithConf(&confPbMsg)
}

// InitLoggerDefault 使用默认配置初始化日志模块
func InitLoggerDefault() (err error) {
	confPbMsg := config.LoggerConfig{
		FileWriterConfig:    &config.LoggerConfig_FileWriterConfig{},
		ConsoleWriterConfig: &config.LoggerConfig_ConsoleWriterConfig{},
	}

	return initLoggerWithConf(&confPbMsg)
}

// RedirectStdoutAndStderr 重定向标准输出和标准错误
func RedirectStdoutAndStderr(filePath string) (err error) {
	var file *os.File
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
	if err != nil {
		fmt.Printf("[Error]: create [%s] file fail with err [%v]\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	syscall.Dup2(int(file.Fd()), 1)
	syscall.Dup2(int(file.Fd()), 2)
	return
}

func initLoggerWithConf(conf *config.LoggerConfig) (err error) {
	if conf.FileWriterConfig != nil {
		// INFO 日志
		if len(conf.FileWriterConfig.GetInfoLogPath()) > 0 {
			w := NewFileWriter()
			w.SetFileName(conf.FileWriterConfig.GetInfoLogPath())
			w.SetLogLevelFloor(LogLevelDebug)
			if len(conf.FileWriterConfig.GetWfLogPath()) > 0 {
				w.SetLogLevelCeiling(LogLevelInfo)
			} else {
				w.SetLogLevelCeiling(LogLevelFatal)
			}
			w.SetRetainHours(int(conf.FileWriterConfig.GetRetainHours()))
			Register(w)
		}

		// WF 日志
		if len(conf.FileWriterConfig.GetWfLogPath()) > 0 {
			w := NewFileWriter()
			w.SetFileName(conf.FileWriterConfig.GetWfLogPath())
			w.SetLogLevelFloor(LogLevelWarn)
			w.SetLogLevelCeiling(LogLevelFatal)
			w.SetRetainHours(int(conf.FileWriterConfig.GetRetainHours()))
			Register(w)
		}
	}

	// 如果是 DAEMON 进程则关闭控制台输出并重定向 stdout 和 stderr
	if _, isDaemon := os.LookupEnv("DAEMON"); isDaemon {
		*conf.ConsoleWriterConfig.Enable = false

		// stdout.log 文件存放在 info 日志的文件夹中, 前面 Register(w) 已经保证了文件夹存在
		logFileDir := filepath.Dir(conf.FileWriterConfig.GetInfoLogPath())
		stdoutLogFilePath := filepath.Join(logFileDir, kStdioLogFile)
		fmt.Printf("[Info] stdout && stderr will redirect to file [%s]\n", stdoutLogFilePath)

		RedirectStdoutAndStderr(stdoutLogFilePath)
	}

	// 控制台日志
	if conf.ConsoleWriterConfig.GetEnable() {
		w := NewConsoleWriter()
		w.SetColor(conf.ConsoleWriterConfig.GetEnableColor())
		if consoleLogLevel, ok := string2logLevel[conf.ConsoleWriterConfig.GetLogLevel().String()]; !ok {
			err = errors.New("invalid log level: " + conf.ConsoleWriterConfig.GetLogLevel().String())
			return
		} else {
			w.SetLevel(consoleLogLevel)
		}
		Register(w)
	}

	if fileLogLevel, ok := string2logLevel[conf.FileWriterConfig.GetLogLevel().String()]; !ok {
		err = errors.New("invalid log level: " + conf.FileWriterConfig.GetLogLevel().String())
		return
	} else {
		SetLevel(fileLogLevel)
	}

	return
}

func isFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
