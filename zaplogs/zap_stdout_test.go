package zaplogs

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestGetStdoutZapLogger(t *testing.T) {
	zapLog := GetStdoutZapLogger(zapcore.DebugLevel)
	zapLog.Debug("abc")
	zapLog.Error("abc", zap.String("xyz", "uvw"))
	zapLog.Info("123")
	zapLog.Warn("abc")
}

func TestNewStdoutZapLogger(t *testing.T) {
	zapLog := NewStdoutZapLogger(false, zapcore.DebugLevel, 0)
	zapLog.Debug("abc")
	zapLog.Error("abc", zap.String("xyz", "uvw"))
	zapLog.Info("123")
	zapLog.Warn("abc")
}
