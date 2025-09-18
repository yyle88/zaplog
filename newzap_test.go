package zaplog_test

import (
	"testing"

	"github.com/pkg/errors"
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
	zaplog.SUG.Debugln(1, 2, 3, 4, 5, 6)
	zaplog.SUG.Debug("1", 2, 3, 4, "5", 6)
	zaplog.SUG.Debugln("1", 2, 3, 4, "5", 6)
	zaplog.SUG.Debug()
	zaplog.SUG.Debug(0)
	zaplog.SUG.Debug([]int{0, 1, 2, 3, 4})
}
