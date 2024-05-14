package zaplog

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog/zaplogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOGGER = MustCreateZapTuple(true, "DEBUG", []string{"stdout"}, 0)
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

func NewZapTupleWithSkip(zlg *zap.Logger, skip int) *ZapTuple {
	return NewZapTuple(zlg.WithOptions(zap.AddCallerSkip(skip)))
}

func MustCreateZapTuple(debug bool, level string, outputPaths []string, skip int) *ZapTuple {
	var opts []zap.Option
	if skip > 0 {
		opts = append(opts, zap.AddCallerSkip(skip))
	}
	config := NewZapConfig(debug, level, outputPaths)
	zlg, err := config.Build(opts...)
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

func (T *ZapTuple) Close() error {
	return T.LOG.Sync()
}

func NewZapConfig(debug bool, level string, outputPaths []string) *zap.Config {
	var config *zap.Config
	if debug {
		config = newConfig(zap.NewDevelopmentConfig())
		// config.DisableStacktrace = true //认为说还是需要打印错误的调用栈的，保持和默认值相同吧
		config.EncoderConfig.EncodeCaller = zaplogs.NewCallerEncoderSimple()
	} else {
		config = newConfig(zap.NewProductionConfig())
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

func newConfig(cfg zap.Config) *zap.Config {
	return &cfg
}
