package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func TestSkipZaps_Skip0(t *testing.T) {
	ZAPS.Skip0.LOG.Debug("msg")
}

func TestSkipZaps_Skip(t *testing.T) {
	ZAPS.Skip(0).LOG.Debug("s", zap.Int("i", 0))
	func() {
		ZAPS.Skip(1).LOG.Debug("s", zap.Int("i", 1))
	}()

	caseSkipZaps(t, 1)
}

func caseSkipZaps(t *testing.T, skipDepth int) {
	zpn := ZAPS.Skip(skipDepth)
	zpn.LOG.Debug("abc", zap.Int("skip", skipDepth))
	if skipDepth < 10 {
		caseSkipLogs(t, skipDepth+1)
	}
}
