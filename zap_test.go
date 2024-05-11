package zaplog

import (
	"testing"

	"github.com/pkg/errors"
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

func TestSubZap(t *testing.T) {
	zp := LOGGER.SubZap("module", zap.String("K", "V"))
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
