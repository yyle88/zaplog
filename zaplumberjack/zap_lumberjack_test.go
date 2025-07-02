package zaplumberjack

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig("stdout", "DEBUG")
	data, err := json.MarshalIndent(config, "", "\t")
	require.NoError(t, err)
	t.Log(string(data))
	output := `{
	"filename": "stdout",
	"max_size": 500,
	"maxbackups": 5,
	"maxage": 3650,
	"compress": false,
	"level": "DEBUG"
}`
	require.Equal(t, output, string(data))
}

func TestNewZapLumberjackFromConfig(t *testing.T) {
	zapLumberjacks := []*ZapLumberjack{
		NewZapLumberjackFromConfig(NewConfig("stdout", "DEBUG")),
		NewZapLumberjackFromConfig(NewConfig("stderr", "ERROR")),
	}

	{
		zapLog := NewLogger(zapLumberjacks, true, 0)
		zapLog.Info("123", zap.String("k", "v"))
		zapLog.Debug("abc", zap.String("k", "v"))
		zapLog.Error("xyz", zap.String("k", "v")) // will be print twice(both to stdout and stderr output)
		zapLog.Warn("uvw", zap.String("k", "v"))
	}
	{
		zapLog := NewLogger(zapLumberjacks, false, 0)
		zapLog.Info("123", zap.String("k", "v"))
		zapLog.Debug("abc", zap.String("k", "v"))
		zapLog.Error("xyz", zap.String("k", "v")) // will be print twice(both to stdout and stderr output)
		zapLog.Warn("uvw", zap.String("k", "v"))
	}
}

func TestNewZapLumberjack(t *testing.T) {
	temp, err := os.MkdirTemp("", "zaplogs_case_simple")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, os.RemoveAll(temp))
	}()
	t.Log(temp)

	debugPath := filepath.Join(temp, "debug.log")
	errorPath := filepath.Join(temp, "error.log")

	zapLumberjacks := []*ZapLumberjack{
		NewZapLumberjackFromConfig(NewConfig(debugPath, "DEBUG")),
		NewZapLumberjackFromConfig(NewConfig(errorPath, "ERROR")),
	}
	defer func() {
		for _, cfg := range zapLumberjacks {
			require.NoError(t, cfg.Close())
		}
	}()

	zapLog := GetLogger(zapLumberjacks)
	for i := 0; i < 3; i++ {
		zapLog.Info("123", zap.String("k", "v"))
		zapLog.Debug("abc", zap.String("k", "v"))
		zapLog.Error("xyz", zap.String("k", "v"))
		zapLog.Warn("uvw", zap.String("k", "v"))
	}
	require.NoError(t, zapLog.Sync())

	showContent(t, debugPath)
	showContent(t, errorPath)
}

func showContent(t *testing.T, path string) {
	t.Log("path:", path)
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	t.Log("data:", string(data))
}
