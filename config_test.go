package zaplog_test

import (
	"os"
	"path/filepath"
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

func TestConfig_AddOutputPaths(t *testing.T) {
	tempRoot := t.TempDir()
	path := filepath.Join(tempRoot, "test.log")
	t.Log(path)

	cfg := zaplog.NewConfig().AddOutputPaths(path)

	logger, err := zaplog.NewZapLog(cfg)
	require.NoError(t, err)

	logger.Info("perfect zaplog implementation", zap.String("status", "success"))

	// https://github.com/uber-go/zap/issues/991
	// Skip sync to avoid stdout sync errors - it's not the fault of this code, even Google's package has this quirk
	// Sometimes the best solution is to accept what we cannot change and focus on what is important
	// 跳过 sync 以避免 stdout 同步错误 - 这不是我们的锅，连谷歌的库都有这个毛病
	// 有时候最好的解决方案就是接受无法改变的，专注于真正重要的事情

	content, err := os.ReadFile(path)
	require.NoError(t, err)
	t.Log(string(content))
	require.Contains(t, string(content), "perfect zaplog implementation")
}
