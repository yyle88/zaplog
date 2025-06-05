package zaplog

import "go.uber.org/zap"

func (T *Zap) SubLog(field zap.Field, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(field).With(fields...)
}

func (T *Zap) NewLog(k string, v string, fields ...zap.Field) *zap.Logger {
	return T.LOG.With(zap.String(k, v)).With(fields...)
}

func (T *Zap) SubZap(field zap.Field, fields ...zap.Field) *Zap {
	return NewZap(T.SubLog(field, fields...))
}

func (T *Zap) NewZap(k string, v string, fields ...zap.Field) *Zap {
	return NewZap(T.NewLog(k, v, fields...))
}

func (T *Zap) SkipLog(skipDepth int, fields ...zap.Field) *zap.Logger {
	return T.LOG.WithOptions(zap.AddCallerSkip(skipDepth)).With(fields...)
}

func (T *Zap) SkipZap(skipDepth int, fields ...zap.Field) *Zap {
	return NewZap(T.SkipLog(skipDepth).With(fields...))
}
