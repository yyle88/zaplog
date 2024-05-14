package zaplogs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
LumberjackConfig 的默认配置

	{
		"filename": "stdout",
		"max_size": 500, //megabytes. Example: 500M means 0.5G
		"maxbackups": 5,
		"maxage": 3650, //days
		"compress": false,
		"level": "debug"
	}
*/
type LumberjackConfig struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxage"`
	Compress   bool   `json:"compress"`
	Level      string `json:"level"`
}

func NewLumberjackSyncX(cfg *LumberjackConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
}

func NewLumberjackZapLogConfig(cfg *LumberjackConfig) *LumberjackZapLogConfig {
	return &LumberjackZapLogConfig{
		SyncX: NewLumberjackSyncX(cfg),
		Level: NewLevelFromString(cfg.Level),
	}
}

type LumberjackZapLogConfig struct {
	SyncX *lumberjack.Logger
	Level zapcore.Level
}

// NewLumberjackZapLog 能够打印一组日志，各个日志使用不同的Level级别，而且能自动分页，当某个日志超过大小时自动切割
// 当然为了在标准输出中打印日志，也支持 stdout 和 stderr 的输出
func NewLumberjackZapLog(cfgs []*LumberjackZapLogConfig) *zap.Logger {
	return NewLumberjackZapLogX(cfgs, true, 0)
}

// NewLumberjackZapLogX skip在内部已经+1因此外部通常传0即可
func NewLumberjackZapLogX(cfgs []*LumberjackZapLogConfig, debug bool, skip int) *zap.Logger {
	if len(cfgs) <= 0 {
		panic("no cfgs")
	}
	cores := make([]zapcore.Core, 0)

	encSimple := NewEncoderSimple(debug)

	for _, cfg := range cfgs {
		switch cfg.SyncX.Filename {
		case "stdout":
			cores = append(cores, zapcore.NewCore(encSimple, os.Stdout, cfg.Level))
		case "stderr":
			cores = append(cores, zapcore.NewCore(encSimple, os.Stderr, cfg.Level))
		default:
			cores = append(cores, zapcore.NewCore(encSimple, zapcore.AddSync(cfg.SyncX), cfg.Level))
		}
	}
	tee := zapcore.NewTee(cores...)

	options := NewOptionsSimple(debug, skip)

	zlg := zap.New(tee, options...)
	return zlg
}
