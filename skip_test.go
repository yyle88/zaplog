package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func TestSkipLogs_Skip(t *testing.T) {
	LOGS.Skip(0).Debug("s", zap.Int("i", 0))
	func() {
		LOGS.Skip(1).Debug("s", zap.Int("i", 1))
	}()

	caseSkipLogs(t, 1)
}

func caseSkipLogs(t *testing.T, skipDepth int) {
	zapLog := LOGS.Skip(skipDepth)
	zapLog.Debug("abc", zap.Int("skip", skipDepth))
	if skipDepth < 10 {
		caseSkipLogs(t, skipDepth+1)
	}
}
