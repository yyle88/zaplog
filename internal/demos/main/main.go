package main

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
	{
		zaplog.LOG.Info("abc", zap.String("xyz", "uvw"))
		show(zaplog.LOG)
	}
	{
		zaplog.LOGS.P0.Info("abc", zap.String("xyz", "uvw"))
		show(zaplog.LOGS.P1)
	}
	{
		config := zaplog.NewZapConfig(true, "debug", []string{"stdout"})
		zlg, err := config.Build()
		if err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		zlg.Info("abc", zap.String("xyz", "uvw"))
		show(zlg)
	}
	{
		config := zaplog.NewZapConfig(false, "debug", []string{"stdout"})
		zlg, err := config.Build()
		if err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		zlg.Info("abc", zap.String("xyz", "uvw"))
		show(zlg)
	}
}

func show(zlg *zap.Logger) {
	zlg.Error("abc", zap.String("xyz", "uvw"))
	show2(zlg)
	zlg.Info("========================================")
}

func show2(zlg *zap.Logger) {
	zlg = zlg.With(zap.String("step2", "show2"))
	zlg.Debug("abc", zap.String("xyz", "uvw"))
	show3(zlg)
}

func show3(zlg *zap.Logger) {
	zlg = zlg.With(zap.String("step3", "show3"))
	zlg.Warn("abc", zap.String("xyz", "uvw"))
}
