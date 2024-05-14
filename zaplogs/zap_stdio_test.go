package zaplogs

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewStdLog(t *testing.T) {
	{
		zlg := NewStdLog(zapcore.DebugLevel)
		zlg.Debug("abc")
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Info("123")
		zlg.Warn("abc")
	}
	{
		zlg := NewStdLogX(false, zapcore.DebugLevel, 0)
		zlg.Debug("abc")
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Info("123")
		zlg.Warn("abc")
	}
}
