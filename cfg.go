package zaplog

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog/internal/utils"
	"github.com/yyle88/zaplog/zaplogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Debug       bool
	Level       string
	OutputPaths []string
	Skip        int
}

func NewConfig() *Config {
	return &Config{
		Debug:       true,
		Level:       "DEBUG",
		OutputPaths: []string{"stdout"},
		Skip:        0,
	}
}

func NewZapLog(cfg *Config) (*zap.Logger, error) {
	config := NewZapConfig(cfg.Debug, cfg.Level, cfg.OutputPaths)

	var opts []zap.Option
	if cfg.Skip > 0 {
		opts = append(opts, zap.AddCallerSkip(cfg.Skip))
	}

	zapLog, err := config.Build(opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "new zap log is wrong")
	}
	return zapLog, nil
}

func NewZapConfig(debug bool, level string, outputPaths []string) *zap.Config {
	var config *zap.Config
	if debug {
		config = utils.PtrX(zap.NewDevelopmentConfig())
		// config.DisableStacktrace = true //认为说还是需要打印错误的调用栈的，保持和默认值相同吧
		config.EncoderConfig.EncodeCaller = zaplogs.NewCallerEncoderTrimmed()
	} else {
		config = utils.PtrX(zap.NewProductionConfig())
		// config.DisableCaller = true //是否在日志中展示文件的路径和代码行号，保持和默认值相同吧
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}
	if len(outputPaths) > 0 {
		config.OutputPaths = outputPaths
	}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.Level = zap.NewAtomicLevelAt(zaplogs.ParseLevel(level))
	return config
}
