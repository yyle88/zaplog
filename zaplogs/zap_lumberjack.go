package zaplogs

import (
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
LumberjackLogConfig 在 json 格式的配置文件里的默认配置:

	{
		"filename": "stdout",
		"max_size": 500, //megabytes. Example: 500M means 0.5G
		"maxbackups": 5,
		"maxage": 3650, //days
		"compress": false,
		"level": "debug"
	}
*/
type LumberjackLogConfig struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxage"`
	Compress   bool   `json:"compress"`
	Level      string `json:"level"`
}

// NewLumberjackLogConfig 返回个默认的配置
func NewLumberjackLogConfig(filename string, level string) *LumberjackLogConfig {
	return &LumberjackLogConfig{
		Filename:   filename,
		MaxSize:    500, //megabytes. Example: 500M means 0.5G
		MaxBackups: 5,
		MaxAge:     3650, //days. Example: 3650 days means 10 years
		Compress:   false,
		Level:      level,
	}
}

type LumberjackLoggerConfig struct {
	Logger *lumberjack.Logger
	Level  zapcore.Level
}

func NewLumberjackLoggerConfigFromConfig(cfg *LumberjackLogConfig) *LumberjackLoggerConfig {
	logger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	return NewLumberjackLoggerConfigFromWriter(logger, ParseLevel(cfg.Level))
}

func NewLumberjackLoggerConfigFromWriter(logger *lumberjack.Logger, level zapcore.Level) *LumberjackLoggerConfig {
	return &LumberjackLoggerConfig{
		Logger: logger,
		Level:  level,
	}
}

func (cfg *LumberjackLoggerConfig) Close() error {
	if err := cfg.Logger.Close(); err != nil { //这里的close在底层是允许多次调用的，只有首次关闭是生效的
		return errors.WithMessage(err, "close lumberjack is wrong")
	}
	return nil
}

// NewLumberjackZapSimple 能够打印一组日志，各个日志使用不同的Level级别，而且能自动分页，当某个日志超过大小时自动切割
// 当然为了在标准输出中打印日志，也支持 stdout 和 stderr 的输出
// 起名好难
func NewLumberjackZapSimple(configs []*LumberjackLoggerConfig) *zap.Logger {
	return NewLumberjackZapLogger(configs, true, 0)
}

// NewLumberjackZapLogger skip在内部已经+1因此外部通常传0即可
func NewLumberjackZapLogger(configs []*LumberjackLoggerConfig, debug bool, skipDepth int) *zap.Logger {
	if len(configs) <= 0 {
		panic("no configs")
	}
	cores := make([]zapcore.Core, 0)

	coEnc := NewEncoder(debug)

	for _, cfg := range configs {
		switch cfg.Logger.Filename {
		case "stdout":
			cores = append(cores, zapcore.NewCore(coEnc, os.Stdout, cfg.Level))
		case "stderr":
			cores = append(cores, zapcore.NewCore(coEnc, os.Stderr, cfg.Level))
		default:
			cores = append(cores, zapcore.NewCore(coEnc, zapcore.AddSync(cfg.Logger), cfg.Level))
		}
	}
	tee := zapcore.NewTee(cores...)

	options := NewLoggerOptions(debug, skipDepth)

	zapLog := zap.New(tee, options...)
	return zapLog
}
