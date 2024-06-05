package zaplogs

import (
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
LumberjackConfig 在 json 格式的配置文件里的默认配置:

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

type LumberjackZapCfg struct {
	SyncX *lumberjack.Logger
	Level zapcore.Level
}

func NewLumberjackZapCFG(cfg *LumberjackConfig) *LumberjackZapCfg {
	syncX := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	return NewLumberjackZapCfg(syncX, NewLevelFromString(cfg.Level))
}

func NewLumberjackZapCfg(syncX *lumberjack.Logger, level zapcore.Level) *LumberjackZapCfg {
	return &LumberjackZapCfg{
		SyncX: syncX,
		Level: level,
	}
}

func (cfg *LumberjackZapCfg) Close() error {
	if err := cfg.SyncX.Close(); err != nil { //这里的close在底层是允许多次调用的，只有首次关闭是生效的
		return errors.WithMessage(err, "close lumberjack is wrong")
	}
	return nil
}

// NewLumberjackZapLOG 能够打印一组日志，各个日志使用不同的Level级别，而且能自动分页，当某个日志超过大小时自动切割
// 当然为了在标准输出中打印日志，也支持 stdout 和 stderr 的输出
// 起名好难
func NewLumberjackZapLOG(cfgs []*LumberjackZapCfg) *zap.Logger {
	return NewLumberjackZapLog(cfgs, true, 0)
}

// NewLumberjackZapLog skip在内部已经+1因此外部通常传0即可
func NewLumberjackZapLog(cfgs []*LumberjackZapCfg, debug bool, skip int) *zap.Logger {
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
