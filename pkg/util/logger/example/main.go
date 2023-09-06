package main

import (
	"context"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/util"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/common"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
)

// 运行: go run pkg/util/logger/example/main.go
func main() {
	if err := logger.InitLogger(util.DirAbsPath() + "/" + "logger.json"); err != nil {
		panic(err)
	}

	defer logger.Close()

	logger.Debug("%s log", "debug")
	logger.Info("%s log", "info")
	logger.Warn("%s log", "warn")
	logger.Error("%s log", "error")

	var (
		ctx = context.WithValue(context.Background(), common.ContextKeyTraceID, common.NewTraceId())
	)
	logger.DebugTrace(ctx, "%s log", "debug")
	logger.InfoTrace(ctx, "%s log", "info")
	logger.WarnTrace(ctx, "%s log", "warn")
	logger.ErrorTrace(ctx, "%s log", "error")

	logger.Fatal("%s log", "fatal")
}
