package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func Test_skipZaps(t *testing.T) {
	ZAPS.P0.LOG.Debug("msg")
}

func Test_skipZaps_Pn(t *testing.T) {
	ZAPS.Pn(0).LOG.Debug("s", zap.Int("i", 0))
	func() {
		ZAPS.Pn(1).LOG.Debug("s", zap.Int("i", 1))
	}()

	caseZap(t, 1)
}

func caseZap(t *testing.T, skip int) {
	zpn := ZAPS.Pn(skip)
	zpn.LOG.Debug("abc", zap.Int("skip", skip))
	if skip < 10 {
		caseRun(t, skip+1)
	}
}
