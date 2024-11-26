package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func TestSkipZaps_P0(t *testing.T) {
	ZAPS.P0.LOG.Debug("msg")
}

func TestSkipZaps_Pn(t *testing.T) {
	ZAPS.Pn(0).LOG.Debug("s", zap.Int("i", 0))
	func() {
		ZAPS.Pn(1).LOG.Debug("s", zap.Int("i", 1))
	}()

	caseSkipZaps(t, 1)
}

func caseSkipZaps(t *testing.T, skipDepth int) {
	zpn := ZAPS.Pn(skipDepth)
	zpn.LOG.Debug("abc", zap.Int("skip", skipDepth))
	if skipDepth < 10 {
		caseSkipLogs(t, skipDepth+1)
	}
}
