package zaplumberjack

import (
	"os"

	"github.com/pkg/errors"
	"github.com/yyle88/zaplog/zaplogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
Config 在 json 格式的配置文件里的默认配置:

	{
		"filename": "stdout",
		"max_size": 500, //megabytes. Example: 500M means 0.5G
		"maxbackups": 5,
		"maxage": 3650, //days
		"compress": false,
		"level": "debug"
	}
*/
type Config struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxage"`
	Compress   bool   `json:"compress"`
	Level      string `json:"level"`
}

// NewConfig 返回个默认的配置
func NewConfig(filename string, level string) *Config {
	return &Config{
		Filename:   filename,
		MaxSize:    500, //megabytes. Example: 500M means 0.5G
		MaxBackups: 5,
		MaxAge:     3650, //days. Example: 3650 days means 10 years
		Compress:   false,
		Level:      level,
	}
}

func NewZapLumberjackFromConfig(cfg *Config) *ZapLumberjack {
	logger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	return NewZapLumberjack(logger, zaplogs.ParseLevel(cfg.Level))
}

type ZapLumberjack struct {
	Logger *lumberjack.Logger
	Level  zapcore.Level
}

func NewZapLumberjack(logger *lumberjack.Logger, level zapcore.Level) *ZapLumberjack {
	return &ZapLumberjack{
		Logger: logger,
		Level:  level,
	}
}

func (cfg *ZapLumberjack) Close() error {
	if err := cfg.Logger.Close(); err != nil { //这里的close在底层是允许多次调用的，只有首次关闭是生效的
		return errors.WithMessage(err, "close lumberjack is wrong")
	}
	return nil
}

// GetLogger 能够打印一组日志，各个日志使用不同的Level级别，而且能自动分页，当某个日志超过大小时自动切割
// 当然为了在标准输出中打印日志，也支持 stdout 和 stderr 的输出
func GetLogger(configs []*ZapLumberjack) *zap.Logger {
	return NewLogger(configs, true, 0)
}

// NewLogger skip在内部已经+1因此外部通常传0即可
func NewLogger(configs []*ZapLumberjack, debug bool, skipDepth int) *zap.Logger {
	if len(configs) <= 0 {
		panic("no configs")
	}
	tee := NewZapTee(configs, zaplogs.NewEncoder(debug))
	options := zaplogs.NewLoggerOptionsWithSkip(debug, skipDepth)
	zapLog := zap.New(tee, options...)
	return zapLog
}

// NewZapTee 返回一个 Tee Core，能够将日志输出到多个 Lumberjack Logger 里面，实现分级别分页的日志
func NewZapTee(configs []*ZapLumberjack, encoder zapcore.Encoder) zapcore.Core {
	cores := make([]zapcore.Core, 0)
	for _, cfg := range configs {
		switch cfg.Logger.Filename {
		case "stdout":
			cores = append(cores, zapcore.NewCore(encoder, os.Stdout, cfg.Level))
		case "stderr":
			cores = append(cores, zapcore.NewCore(encoder, os.Stderr, cfg.Level))
		default:
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(cfg.Logger), cfg.Level))
		}
	}
	tee := zapcore.NewTee(cores...)
	return tee
}
