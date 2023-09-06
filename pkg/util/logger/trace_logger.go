package logger

import (
	"context"
	"fmt"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/common"
)

// https://stackoverflow.com/questions/56343599/how-to-add-trace-id-to-each-logs-in-go-micro-service
// 在 Logger 中通过 context 中的 TraceID 追踪不同的请求链路

var (
	genLogPrefix = func(ctx context.Context) string {
		var traceID = ctx.Value(common.ContextKeyTraceID)
		if traceID == nil {
			return "[0] "
		}
		return fmt.Sprintf("[%s] ", traceID)

	}
)

func DebugTrace(ctx context.Context, fmt string, args ...interface{}) {
	Debug(genLogPrefix(ctx)+fmt, args...)
}

func InfoTrace(ctx context.Context, fmt string, args ...interface{}) {
	Info(genLogPrefix(ctx)+fmt, args...)
}

func WarnTrace(ctx context.Context, fmt string, args ...interface{}) {
	Warn(genLogPrefix(ctx)+fmt, args...)
}

func ErrorTrace(ctx context.Context, fmt string, args ...interface{}) {
	Error(genLogPrefix(ctx)+fmt, args...)
}
