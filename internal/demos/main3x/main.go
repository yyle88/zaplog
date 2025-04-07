package main

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
	{
		zapLog, err := zaplog.NewZapLog(zaplog.NewConfig())
		if err != nil {
			panic(errors.Wrap(err, "wrong"))
		}
		zaplog.SetLog(zapLog.With(zap.String("kkk", "vvv")))
	}
	{
		zaplog.LOG.Info("abc", zap.String("xyz", "uvw"))
		zaplog.LOG.Error("abc", zap.String("xyz", "uvw"))
		zaplog.LOG.Debug("abc", zap.String("xyz", "uvw"))
		zaplog.LOG.Warn("abc", zap.String("xyz", "uvw"))
	}
	{
		zaplog.SUG.Infof("abc xyz=%v", "uvw")
		zaplog.SUG.Errorf("abc xyz=%v", "uvw")
		zaplog.SUG.Debugf("abc xyz=%v", "uvw")
		zaplog.SUG.Warnf("abc xyz=%v", "uvw")
	}
}
