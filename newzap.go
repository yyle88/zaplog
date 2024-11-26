package zaplog

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Zap struct {
	LOG *zap.Logger
	SUG *zap.SugaredLogger //比较慢但也是种简单的调用接口
}

func NewZap(zapLog *zap.Logger) *Zap {
	return &Zap{
		LOG: zapLog,
		SUG: zapLog.Sugar(),
	}
}

func NewZapSkip(zapLog *zap.Logger, skipDepth int) *Zap {
	return NewZap(zapLog.WithOptions(zap.AddCallerSkip(skipDepth)))
}

func MustNewZap(cfg *Config) *Zap {
	zapLog, err := NewZapLog(cfg)
	if err != nil {
		panic(errors.Wrap(err, "ERROR WHEN NEW ZAP LOG"))
	}
	return NewZap(zapLog)
}

func (T *Zap) Close() error {
	return T.LOG.Sync()
}
