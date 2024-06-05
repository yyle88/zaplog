package main

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
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
