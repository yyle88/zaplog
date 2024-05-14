package zaplogw

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestZapLogw_Debug(t *testing.T) {
	zapLogw := NewZapLogw(zaplog.ZAPS.P1.SUG)
	zapLogw.Debug("abc")
	zapLogw.Debug("xyz", "num", 123)
	zapLogw.Debug("uvw", "res", "xxx")
}
