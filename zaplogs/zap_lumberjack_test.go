package zaplogs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewLumberjackZapLog(t *testing.T) {
	cfgs := []*LumberjackZapLogConfig{
		NewLumberjackZapLogConfig(&LumberjackConfig{
			Filename:   "stdout",
			MaxSize:    500,
			MaxBackups: 5,
			MaxAge:     20160,
			Compress:   false,
			Level:      "debug",
		}),
		NewLumberjackZapLogConfig(&LumberjackConfig{
			Filename:   "stderr",
			MaxSize:    500,
			MaxBackups: 5,
			MaxAge:     20160,
			Compress:   false,
			Level:      "error",
		}),
	}

	{
		zlg := NewLumberjackZapLogX(cfgs, true, 0)
		zlg.Info("123", zap.String("k", "v"))
		zlg.Debug("abc", zap.String("k", "v"))
		zlg.Error("xyz", zap.String("k", "v")) // will be print twice(both to stdout and stderr output)
		zlg.Warn("uvw", zap.String("k", "v"))
	}
	{
		zlg := NewLumberjackZapLogX(cfgs, false, 0)
		zlg.Info("123", zap.String("k", "v"))
		zlg.Debug("abc", zap.String("k", "v"))
		zlg.Error("xyz", zap.String("k", "v")) // will be print twice(both to stdout and stderr output)
		zlg.Warn("uvw", zap.String("k", "v"))
	}
}

func TestNewLumberjackZapLogX(t *testing.T) {
	temp, err := os.MkdirTemp("", "zaplogs_case_simple")
	require.NoError(t, err)
	defer os.RemoveAll(temp)
	t.Log(temp)

	debugPath := filepath.Join(temp, "debug.log")
	errorPath := filepath.Join(temp, "error.log")

	cfgs := []*LumberjackZapLogConfig{
		NewLumberjackZapLogConfig(&LumberjackConfig{
			Filename:   debugPath,
			MaxSize:    500,
			MaxBackups: 5,
			MaxAge:     20160,
			Compress:   false,
			Level:      "debug",
		}),
		NewLumberjackZapLogConfig(&LumberjackConfig{
			Filename:   errorPath,
			MaxSize:    500,
			MaxBackups: 5,
			MaxAge:     20160,
			Compress:   false,
			Level:      "error",
		}),
	}

	zlg := NewLumberjackZapLog(cfgs)
	for i := 0; i < 3; i++ {
		zlg.Info("123", zap.String("k", "v"))
		zlg.Debug("abc", zap.String("k", "v"))
		zlg.Error("xyz", zap.String("k", "v"))
		zlg.Warn("uvw", zap.String("k", "v"))
	}
	require.NoError(t, zlg.Sync())

	showContent(t, debugPath)
	showContent(t, errorPath)
}

func showContent(t *testing.T, path string) {
	t.Log("path:", path)
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	t.Log("data:", string(data))
}
