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
		zapLog, err := config.Build()
		if err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		zapLog.Info("abc", zap.String("xyz", "uvw"))
		show(zapLog)
	}
	{
		config := zaplog.NewZapConfig(false, "debug", []string{"stdout"})
		zapLog, err := config.Build()
		if err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		zapLog.Info("abc", zap.String("xyz", "uvw"))
		show(zapLog)
	}
}

func show(zapLog *zap.Logger) {
	zapLog.Error("abc", zap.String("xyz", "uvw"))
	show2(zapLog)
	zapLog.Info("========================================")
}

func show2(zapLog *zap.Logger) {
	zapLog = zapLog.With(zap.String("step2", "show2"))
	zapLog.Debug("abc", zap.String("xyz", "uvw"))
	show3(zapLog)
}

func show3(zapLog *zap.Logger) {
	zapLog = zapLog.With(zap.String("step3", "show3"))
	zapLog.Warn("abc", zap.String("xyz", "uvw"))
}
