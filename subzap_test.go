package zaplog_test

import (
	"testing"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func TestZap_SubZap(t *testing.T) {
	zp := zaplog.LOGGER.SubZap(zap.String("module", "abc"), zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}

func TestZap_NewZap(t *testing.T) {
	zp := zaplog.LOGGER.NewZap("module", "abc", zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}

func TestZap_SkipLog(t *testing.T) {
	zapLog := zaplog.LOGGER.SubZap(zap.String("module", "skip-log-test-case"))
	zapLog.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))

	run := func(t *testing.T) {
		subLog := zapLog.SkipLog(1, zap.Int("c", 3))
		subLog.Debug("abc", zap.Int("a", 1), zap.Int("b", 2))
	}
	run(t)
}

func TestZap_SkipZap(t *testing.T) {
	zapLog := zaplog.LOGGER.SubZap(zap.String("module", "skip-zap-test-case"))
	zapLog.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))

	run := func(t *testing.T) {
		subZap := zapLog.SkipZap(1, zap.Int("c", 3))
		subZap.LOG.Debug("abc", zap.Int("a", 1), zap.Int("b", 2))
	}
	run(t)
}
