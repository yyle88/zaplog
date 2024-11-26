package zaplog_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	zaplog.LOG.Debug("1", zap.String("a", "0"))
	zaplog.LOG.Error("1", zap.Error(errors.New("e")))
}

func TestError(t *testing.T) {
	zaplog.LOG.Debug("1", zap.String("a", "0"))
	zaplog.LOG.Error("1", zap.Error(errors.New("e")))
}

func TestDebug2(t *testing.T) {
	zaplog.LOG.Debug("ok", zap.String("A", "aaa"))
	zaplog.LOG.Debug("ok", zap.String("B", "bbb"))
}

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

func TestDebug3(t *testing.T) {
	zaplog.LOG.Debug("1")
	zaplog.LOG.Error("2", zap.Error(errors.New("x")))
}

func TestDebug4(t *testing.T) {
	zaplog.SUG.Debug("1")
	zaplog.SUG.Error("2", errors.New("x"))
}

func TestDebug5(t *testing.T) {
	zaplog.SUG.Debug(1, 2, 3, 4, 5, 6)
	zaplog.SUG.Debug("1", 2, 3, 4, 5, 6)
	zaplog.SUG.Debug()
	zaplog.SUG.Debug(0)
	zaplog.SUG.Debug([]int{0, 1, 2, 3, 4})
}

func TestNewZapConfig(t *testing.T) {
	{
		config := zaplog.NewZapConfig(true, "debug", []string{"stdout"})
		zlg, err := config.Build()
		require.NoError(t, err)
		zlg.Info("abc", zap.String("xyz", "uvw"))
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Debug("abc", zap.String("xyz", "uvw"))
		zlg.Warn("abc", zap.String("xyz", "uvw"))
	}
	{
		config := zaplog.NewZapConfig(false, "debug", []string{"stdout"})
		zlg, err := config.Build()
		require.NoError(t, err)
		zlg.Info("abc", zap.String("xyz", "uvw"))
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Debug("abc", zap.String("xyz", "uvw"))
		zlg.Warn("abc", zap.String("xyz", "uvw"))
	}
}
