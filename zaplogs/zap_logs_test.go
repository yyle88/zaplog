package zaplogs

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestParseLevelCode(t *testing.T) {
	require.Equal(t, zap.DebugLevel, ParseLevelCode("debug"))
	require.Equal(t, zap.InfoLevel, ParseLevelCode("info"))
	require.Equal(t, zap.WarnLevel, ParseLevelCode("warn"))
	require.Equal(t, zap.ErrorLevel, ParseLevelCode("error"))
	require.Equal(t, zap.PanicLevel, ParseLevelCode("panic"))

	require.Equal(t, zap.InfoLevel, ParseLevelCode("unknown"))
}
