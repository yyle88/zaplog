package zaplog

import "go.uber.org/zap"

// SetLog 使用这个函数重新设置全部的全局变量，因为我所有的包都有可能需要打日志，因此把日志做成了全局变量
// 当需要修改时就在这里全局重新设置下就行
func SetLog(zlg *zap.Logger) {
	LOGGER = &ZapTuple{
		LOG: zlg,
		SUG: zlg.Sugar(),
	}
	LOG = LOGGER.LOG //底下的函数会依赖于这个变量，它们会重新构造新对象
	SUG = LOGGER.SUG

	LOGS = NewSkipLogs()
	ZAPS = NewSkipZaps()
}
