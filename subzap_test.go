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

func TestZap_SubZap2(t *testing.T) {
	zp := zaplog.LOGGER.SubZap2("module", "abc", zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}

func TestSubZap3(t *testing.T) {
	zp := zaplog.LOGGER.SubZap3("module", zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}
