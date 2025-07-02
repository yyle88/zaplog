package zaplogs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestParseLevel(t *testing.T) {
	require.Equal(t, zap.DebugLevel, ParseLevel("debug"))
	require.Equal(t, zap.InfoLevel, ParseLevel("info"))
	require.Equal(t, zap.WarnLevel, ParseLevel("warn"))
	require.Equal(t, zap.ErrorLevel, ParseLevel("error"))
	require.Equal(t, zap.PanicLevel, ParseLevel("panic"))
	require.Equal(t, zap.InfoLevel, ParseLevel("unknown"))
}

func TestNewEncoderUsage(t *testing.T) {
	encoder := NewEncoder(false)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core)
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}

func TestNewEncoderDebug(t *testing.T) {
	encoder := NewEncoder(true)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core)
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}

func TestNewDevelopmentEncoder(t *testing.T) {
	encoder := NewDevelopmentEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core)
	logger.Debug("This is a debug message with development encoder")
	logger.Info("This is an info message with development encoder")
	logger.Warn("This is a warning message with development encoder")
	logger.Error("This is an error message with development encoder")
}

func TestNewProductionsEncoder(t *testing.T) {
	encoder := NewProductionsEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core)
	logger.Debug("This is a debug message with production encoder")
	logger.Info("This is an info message with production encoder")
	logger.Warn("This is a warning message with production encoder")
	logger.Error("This is an error message with production encoder")
}

func TestNewCallerEncoderTrimPath(t *testing.T) {
	encoderFunc := NewCallerEncoderTrimPath()
	require.NotNil(t, encoderFunc)

	encConfig := zap.NewDevelopmentEncoderConfig()
	encConfig.EncodeCaller = encoderFunc
	enc := zapcore.NewConsoleEncoder(encConfig)
	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Debug("This is a debug message with custom caller encoder")
	logger.Info("This is an info message with custom caller encoder")
	logger.Warn("This is a warning message with custom caller encoder")
	logger.Error("This is an error message with custom caller encoder")
}

func TestNewCallerEncoderFullPath(t *testing.T) {
	encoderFunc := NewCallerEncoderFullPath()
	require.NotNil(t, encoderFunc)

	encConfig := zap.NewDevelopmentEncoderConfig()
	encConfig.EncodeCaller = encoderFunc
	enc := zapcore.NewConsoleEncoder(encConfig)
	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Debug("This is a debug message with custom caller encoder")
	logger.Info("This is an info message with custom caller encoder")
	logger.Warn("This is a warning message with custom caller encoder")
	logger.Error("This is an error message with custom caller encoder")
}

func TestNewLoggerOptionsWithSkip(t *testing.T) {
	options := NewLoggerOptionsWithSkip(true, 1)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)
	logger := zap.New(core, options...)
	require.NotNil(t, logger)
	logger.Debug("This is a debug message with custom options")
	logger.Info("This is an info message with custom options")
	logger.Warn("This is a warning message with custom options")
	logger.Error("This is an error message with custom options")
}
