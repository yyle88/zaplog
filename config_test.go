package zaplog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func TestNewZapConfig(t *testing.T) {
	{
		config := zaplog.NewZapConfig(true, "debug", []string{"stdout"})
		zapLog, err := config.Build()
		require.NoError(t, err)
		zapLog.Info("abc", zap.String("xyz", "uvw"))
		zapLog.Error("abc", zap.String("xyz", "uvw"))
		zapLog.Debug("abc", zap.String("xyz", "uvw"))
		zapLog.Warn("abc", zap.String("xyz", "uvw"))
	}
	{
		config := zaplog.NewZapConfig(false, "debug", []string{"stdout"})
		zapLog, err := config.Build()
		require.NoError(t, err)
		zapLog.Info("abc", zap.String("xyz", "uvw"))
		zapLog.Error("abc", zap.String("xyz", "uvw"))
		zapLog.Debug("abc", zap.String("xyz", "uvw"))
		zapLog.Warn("abc", zap.String("xyz", "uvw"))
	}
}
