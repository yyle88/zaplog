package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func Test_skipLogs_Pn(t *testing.T) {
	LOGS.Pn(0).Debug("s", zap.Int("i", 0))
	func() {
		LOGS.Pn(1).Debug("s", zap.Int("i", 1))
	}()

	caseRun(t, 1)
}

func caseRun(t *testing.T, skip int) {
	zlg := LOGS.Pn(skip)
	zlg.Debug("abc", zap.Int("skip", skip))
	if skip < 10 {
		caseRun(t, skip+1)
	}
}
