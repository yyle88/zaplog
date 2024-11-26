package zaplogs

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewStdoutLog(t *testing.T) {
	zapLog := NewStdoutLog(zapcore.DebugLevel)
	zapLog.Debug("abc")
	zapLog.Error("abc", zap.String("xyz", "uvw"))
	zapLog.Info("123")
	zapLog.Warn("abc")
}

func TestNewStdoutZap(t *testing.T) {
	zapLog := NewStdoutZap(false, zapcore.DebugLevel, 0)
	zapLog.Debug("abc")
	zapLog.Error("abc", zap.String("xyz", "uvw"))
	zapLog.Info("123")
	zapLog.Warn("abc")
}
