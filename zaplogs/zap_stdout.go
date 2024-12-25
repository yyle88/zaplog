package zaplogs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewStdoutLog 日志输出到stdio里
func NewStdoutLog(level zapcore.Level) *zap.Logger {
	return NewStdoutZap(true, level, 0)
}

// NewStdoutZap skip+1的设计思路是，我们要把封装关系在逐层慢慢消化掉，否则，在不同层调用时就会有问题
// 但事实上认为这个没什么作用
// 除非是专门封装个日志包，在里面全部重新调用 debug/info/warn 的时候，能够有预见的只跳过1层函数调用，打印调用位置
func NewStdoutZap(debug bool, level zapcore.Level, skipDepth int) *zap.Logger {
	wSync := zapcore.AddSync(os.Stdout)
	cores := make([]zapcore.Core, 0)

	encSimple := NewEncoderSimple(debug)

	cores = append(cores, zapcore.NewCore(encSimple, wSync, level))
	tee := zapcore.NewTee(cores...)

	options := NewOptionsSimple(debug, skipDepth)

	zapLog := zap.New(tee, options...) //修改堆栈深度
	return zapLog
}