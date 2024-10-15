package zaplog

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog/zaplogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOGGER = MustNewZap(NewConfig())
var ZAP = LOGGER
var LOG = LOGGER.LOG //最常用的日志
var SUG = LOGGER.SUG

type Zap struct {
	LOG *zap.Logger
	SUG *zap.SugaredLogger //比较慢但也是种简单的调用接口
}

func NewZap(zlg *zap.Logger) *Zap {
	return &Zap{
		LOG: zlg,
		SUG: zlg.Sugar(),
	}
}

func NewZapSkip(zlg *zap.Logger, skip int) *Zap {
	return NewZap(zlg.WithOptions(zap.AddCallerSkip(skip)))
}

func MustNewZap(cfg *Config) *Zap {
	zlg, err := NewZapLog(cfg)
	if err != nil {
		panic(errors.Wrap(err, "ERROR WHEN NEW ZAP LOG"))
	}
	return NewZap(zlg)
}

func (T *Zap) SubLog(zp zap.Field, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zp).With(fields...)
}

func (T *Zap) SubZap(zp zap.Field, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog(zp, fields...))
}

func (T *Zap) Sub(zp zap.Field, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog(zp, fields...))
}

func (T *Zap) SubLog2(k, v string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String(k, v)).With(fields...)
}

func (T *Zap) SubZap2(k, v string, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog2(k, v, fields...))
}

func (T *Zap) Sub2(k, v string, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog2(k, v, fields...))
}

func (T *Zap) SubLog3(module string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String("module", module)).With(fields...)
}

func (T *Zap) SubZap3(module string, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog3(module, fields...))
}

func (T *Zap) Sub3(module string, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog3(module, fields...))
}

func (T *Zap) Close() error {
	return T.LOG.Sync()
}

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

	zlg, err := config.Build(opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "new zap log is wrong")
	}
	return zlg, nil
}

func NewZapConfig(debug bool, level string, outputPaths []string) *zap.Config {
	var config *zap.Config
	if debug {
		config = newCfg(zap.NewDevelopmentConfig())
		// config.DisableStacktrace = true //认为说还是需要打印错误的调用栈的，保持和默认值相同吧
		config.EncoderConfig.EncodeCaller = zaplogs.NewCallerEncoderTrimPC()
	} else {
		config = newCfg(zap.NewProductionConfig())
		// config.DisableCaller = true //是否在日志中展示文件的路径和代码行号，保持和默认值相同吧
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}
	if len(outputPaths) > 0 {
		config.OutputPaths = outputPaths
	}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.Level = zap.NewAtomicLevelAt(zaplogs.NewLevelFromString(level))
	return config
}

func newCfg(cfg zap.Config) *zap.Config {
	return &cfg //就是把结构体类型的配置转换为指针类型
}
