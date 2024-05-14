package zaplogw

import "go.uber.org/zap"

// ZapLogw 在某些场景下需要按照 kvs/fields 的格式来打印日志，因此做个转换的类型，把普通的zap转换成满足 kvs/fields 接口的
// 通常是传给其它自定义的日志接口，但不知道接口是需要 Warn 函数，还是需要 Warning 函数，因此两个都实现就行，这样就更能适应接口
// 需要注意的是因为有这一层的封装，因此传进去的 sug 对象需要 skip 这层，否则打印的位置就是不完美的
type ZapLogw struct {
	SUG *zap.SugaredLogger
}

func NewZapLogw(sug *zap.SugaredLogger) *ZapLogw {
	return &ZapLogw{SUG: sug}
}

func (l *ZapLogw) Debug(msg string, kvs ...interface{}) {
	l.SUG.Debugw(msg, kvs...)
}

func (l *ZapLogw) Info(msg string, kvs ...interface{}) {
	l.SUG.Infow(msg, kvs...)
}

func (l *ZapLogw) Error(msg string, kvs ...interface{}) {
	l.SUG.Errorw(msg, kvs...)
}

func (l *ZapLogw) Fatal(msg string, kvs ...interface{}) {
	l.SUG.Fatalw(msg, kvs...)
}

// Panic is for log warning level
func (l *ZapLogw) Panic(msg string, kvs ...interface{}) {
	l.SUG.Panicw(msg, kvs...)
}

func (l *ZapLogw) Warning(msg string, kvs ...interface{}) {
	l.SUG.Warnw(msg, kvs...)
}

func (l *ZapLogw) Warn(msg string, kvs ...interface{}) {
	l.SUG.Warnw(msg, kvs...)
}
