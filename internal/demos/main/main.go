package main

import (
	"github.com/pkg/errors"
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
		show()
	}
	{
		config := zaplog.NewZapConfig(true, "debug", []string{"stdout"})
		zlg, err := config.Build()
		if err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		zlg.Info("abc", zap.String("xyz", "uvw"))
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Debug("abc", zap.String("xyz", "uvw"))
		zlg.Warn("abc", zap.String("xyz", "uvw"))
	}
	{
		config := zaplog.NewZapConfig(false, "debug", []string{"stdout"})
		zlg, err := config.Build()
		if err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		zlg.Info("abc", zap.String("xyz", "uvw"))
		zlg.Error("abc", zap.String("xyz", "uvw"))
		zlg.Debug("abc", zap.String("xyz", "uvw"))
		zlg.Warn("abc", zap.String("xyz", "uvw"))
	}
}

func show() {
	zaplog.LOGS.P1.Info("abc", zap.String("xyz", "uvw"))
	zaplog.LOGS.P1.Error("abc", zap.String("xyz", "uvw"))
	zaplog.LOGS.P1.Debug("abc", zap.String("xyz", "uvw"))
	zaplog.LOGS.P1.Warn("abc", zap.String("xyz", "uvw"))
}
