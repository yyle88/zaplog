package zaplogw

import (
	"log/slog"
	"testing"

	"github.com/yyle88/zaplog"
)

func TestZapLogw_Debug(t *testing.T) {
	zpw := NewZapLogw(zaplog.ZAPS.Skip1.SUG)
	zpw.Debug("abc")
	zpw.Debug("xyz", "num", 123)
	zpw.Debug("uvw", "res", "x")
}

func TestSlogExample(t *testing.T) {
	slog.Info("message", "x", 1, "y", 2, "z", 3)
}
