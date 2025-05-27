package sketch2

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestExample(t *testing.T) {
	core, logs := observer.New(zap.InfoLevel)
	zaplog.SetLog(zap.New(core))

	zaplog.LOG.Debug("abc")
	t.Log(logs.Len()) // output: 1. Because we set the level to InfoLevel, Debug messages are not in logs.

	zaplog.LOG.Info("123")
	t.Log(logs.Len())

	zaplog.LOG.Info("xyz", zap.Int("num", 1024), zap.String("msg", "ok"))
	t.Log(logs.Len())

	entries := logs.All()
	require.Len(t, entries, 2)
	{
		item := entries[0]
		require.Equal(t, zap.InfoLevel, item.Level)
		require.Equal(t, "123", item.Message)
		require.Empty(t, item.ContextMap())
	}
	{
		item := entries[1]
		require.Equal(t, zap.InfoLevel, item.Level)
		require.Equal(t, "xyz", item.Message)
		require.Equal(t, int64(1024), item.ContextMap()["num"])
		require.Equal(t, "ok", item.ContextMap()["msg"])
	}
}
