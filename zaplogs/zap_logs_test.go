package zaplogs

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestParseLevel(t *testing.T) {
	require.Equal(t, zap.DebugLevel, ParseLevel("debug"))
	require.Equal(t, zap.InfoLevel, ParseLevel("info"))
	require.Equal(t, zap.WarnLevel, ParseLevel("warn"))
	require.Equal(t, zap.ErrorLevel, ParseLevel("error"))
	require.Equal(t, zap.PanicLevel, ParseLevel("panic"))

	require.Equal(t, zap.InfoLevel, ParseLevel("unknown"))
}
