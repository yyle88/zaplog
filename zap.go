package zaplog

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog/zaplogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOGGER = MustNewZapTuple(NewConfig())
var LOG = LOGGER.LOG //最常用的日志
var SUG = LOGGER.SUG

type ZapTuple struct {
	LOG *zap.Logger
	SUG *zap.SugaredLogger //比较慢但也是种简单的调用接口
}

func NewZapTuple(zlg *zap.Logger) *ZapTuple {
	return &ZapTuple{
		LOG: zlg,
		SUG: zlg.Sugar(),
	}
}

func NewZapTupleSkip(zlg *zap.Logger, skip int) *ZapTuple {
	return NewZapTuple(zlg.WithOptions(zap.AddCallerSkip(skip)))
}

func MustNewZapTuple(cfg *Config) *ZapTuple {
	zlg, err := NewZapLog(cfg)
	if err != nil {
		panic(errors.Wrap(err, "ERROR WHEN NEW LOG"))
	}
	return NewZapTuple(zlg)
}

func (T *ZapTuple) SubLog(module string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String("module", module)).With(fields...)
}

func (T *ZapTuple) SubZap(module string, fields ...zap.Field) *ZapTuple {
	return NewZapTuple(T.SubLog(module, fields...))
}

func (T *ZapTuple) SubLog2(k, v string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String(k, v)).With(fields...)
}

func (T *ZapTuple) SubZap2(k, v string, fields ...zap.Field) *ZapTuple {
	return NewZapTuple(T.SubLog2(k, v, fields...))
}

func (T *ZapTuple) Close() error {
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
		config.EncoderConfig.EncodeCaller = zaplogs.NewCallerEncoderSimple()
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
