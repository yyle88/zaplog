package zaplogs

import (
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLevelFromString(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}

func NewEncoderSimple(debug bool) zapcore.Encoder {
	if debug {
		return NewDevelopmentEncoderSimple()
	} else {
		return NewProductionEncoderSimple()
	}
}

func NewDevelopmentEncoderSimple() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = NewCallerEncoderSimple()
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func NewProductionEncoderSimple() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder //这个其实是默认值，但这里依然是赋值下以便于使用
	return zapcore.NewJSONEncoder(encoderConfig)
}

func NewCallerEncoderSimple() func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	return func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
	}
}

func NewOptionsSimple(debug bool, skip int) []zap.Option {
	var options = []zap.Option{zap.AddCaller(), zap.AddCallerSkip(skip)}

	if debug {
		options = append(options, zap.AddStacktrace(zapcore.WarnLevel))
	} else {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	return options
}
