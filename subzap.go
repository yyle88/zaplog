package zaplog

import "go.uber.org/zap"

func (T *Zap) SubLog(zp zap.Field, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zp).With(fields...)
}

func (T *Zap) SubLog2(k, v string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String(k, v)).With(fields...)
}

func (T *Zap) SubModuleLog(module string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String("module", module)).With(fields...)
}

func (T *Zap) SubZap(zp zap.Field, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog(zp, fields...))
}

func (T *Zap) SubZap2(k, v string, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog2(k, v, fields...))
}

func (T *Zap) SubModuleZap(module string, fields ...zap.Field) *Zap {
	return NewZap(T.SubModuleLog(module, fields...))
}
