package zaplog

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	LOG.Debug("1", zap.String("a", "0"))
	LOG.Error("1", zap.Error(errors.New("e")))
}

func TestError(t *testing.T) {
	LOG.Debug("1", zap.String("a", "0"))
	LOG.Error("1", zap.Error(errors.New("e")))
}

func TestDebug2(t *testing.T) {
	LOG.Debug("ok", zap.String("A", "aaa"))
	LOG.Debug("ok", zap.String("B", "bbb"))
}

func TestZap_SubZap(t *testing.T) {
	zp := LOGGER.SubZap(zap.String("module", "abc"), zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}

func TestZap_SubZap2(t *testing.T) {
	zp := LOGGER.SubZap2("module", "abc", zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}

func TestSubZap3(t *testing.T) {
	zp := LOGGER.SubZap3("module", zap.String("K", "V"))
	zp.LOG.Debug("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Debug(1, 2, 3)
	zp.LOG.Error("msg", zap.Int("a", 1), zap.Int("b", 2))
	zp.SUG.Error(1, 2, 3)
}

func TestDebug3(t *testing.T) {
	LOG.Debug("1")
	LOG.Error("2", zap.Error(errors.New("x")))
}

func TestDebug4(t *testing.T) {
	SUG.Debug("1")
	SUG.Error("2", errors.New("x"))
}

func TestDebug5(t *testing.T) {
	SUG.Debug(1, 2, 3, 4, 5, 6)
	SUG.Debug("1", 2, 3, 4, 5, 6)
	SUG.Debug()
	SUG.Debug(0)
	SUG.Debug([]int{0, 1, 2, 3, 4})
}

func TestNewZapConfig(t *testing.T) {
	{
		config := NewZapConfig(true, "debug", []string{"stdout"})
		zlg, err := config.Build()
		require.NoError(t, err)
		zlg.Info("abc", zap.String("xyz", "uvw"))
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Debug("abc", zap.String("xyz", "uvw"))
		zlg.Warn("abc", zap.String("xyz", "uvw"))
	}
	{
		config := NewZapConfig(false, "debug", []string{"stdout"})
		zlg, err := config.Build()
		require.NoError(t, err)
		zlg.Info("abc", zap.String("xyz", "uvw"))
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Debug("abc", zap.String("xyz", "uvw"))
		zlg.Warn("abc", zap.String("xyz", "uvw"))
	}
}
