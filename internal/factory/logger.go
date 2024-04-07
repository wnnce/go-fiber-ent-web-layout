package factory

import (
	"io"
	"log/slog"
	"os"
	"sync"
)

// LoggerOptions 自定义全局日志对象参数
type LoggerOptions struct {
	Level  slog.Level        // 输入的日志等级
	Args   map[string]string // 日志携带的自定义参数
	Output io.Writer         // 日志的输出位置
}

var (
	logger  *slog.Logger
	once    sync.Once
	options = &LoggerOptions{
		Level:  slog.LevelDebug,
		Args:   map[string]string{"appName": "fiber-example"},
		Output: os.Stdout,
	}
)

func initLogger() {
	handler := slog.NewJSONHandler(options.Output, &slog.HandlerOptions{
		Level: options.Level,
	})
	logger = slog.New(handler)
	for k, v := range options.Args {
		logger = logger.With(k, v)
	}
}

// SetLoggerOptions 应用启动时更新全局日志配置
func SetLoggerOptions(op *LoggerOptions) {
	options = op
}

// GetLogger 获取全局日志对象
// traceName 参数会在日志中输入 可以用来区别不同的模块或文件
func GetLogger(traceName string) *slog.Logger {
	once.Do(initLogger)
	return logger.With("trace-name", traceName)
}
