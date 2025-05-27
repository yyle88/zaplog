package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func TestZapLog(t *testing.T) {
	ZAP.LOG.Debug("abc", zap.Int("num", 123))
	ZAP.SUG.Debug("abc", "-|-", "num", "-|-", 123)
}
