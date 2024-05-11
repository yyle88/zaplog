package zaplog

import (
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOGGER = NewZapTuple(true, "DEBUG", []string{"stdout"}, 0)
var LOG = LOGGER.LOG //最常用的日志
var SUG = LOGGER.SUG

type ZapTuple struct {
	LOG *zap.Logger
	SUG *zap.SugaredLogger //比较慢但也是种简单的调用接口
}

func NewZapTuple(debug bool, level string, outputNames []string, skip int) *ZapTuple {
	var opts []zap.Option
	if skip > 0 {
		opts = append(opts, zap.AddCallerSkip(skip))
	}
	config := NewZapConfig(debug, level, outputNames)
	xlg, e := config.Build(opts...)
	if e != nil {
		panic(errors.Wrap(e, "ERROR WHEN NEW LOG"))
	}
	return &ZapTuple{
		LOG: xlg,
		SUG: xlg.Sugar(),
	}
}

func (T *ZapTuple) SubLog(module string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String("module", module)).With(fields...)
}

func (T *ZapTuple) SubZap(module string, fields ...zap.Field) *ZapTuple {
	zlg := T.SubLog(module, fields...)
	return &ZapTuple{
		LOG: zlg,
		SUG: zlg.Sugar(),
	}
}

func (T *ZapTuple) Close() error {
	return T.LOG.Sync()
}

func NewZapConfig(debug bool, level string, outputPaths []string) (zpc *zap.Config) {
	if debug {
		config := zap.NewDevelopmentConfig()
		zpc = &config
	} else {
		config := zap.NewProductionConfig()
		config.DisableCaller = true //是否在日志中展示文件的路径和代码行号
		zpc = &config
	}
	if len(outputPaths) > 0 {
		zpc.OutputPaths = outputPaths
	}
	zpc.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeCaller := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
	}
	zpc.EncoderConfig.EncodeCaller = encodeCaller
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "panic":
		zapLevel = zap.PanicLevel
	default:
		zapLevel = zap.InfoLevel
	}
	zpc.Level = zap.NewAtomicLevelAt(zapLevel)
	return zpc
}
