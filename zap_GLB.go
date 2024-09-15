package zaplog

import "go.uber.org/zap"

// SetLog 重新设置该包里的全局变量，当你不喜欢默认的日志格式时可以用这个函数设置新的默认风格
// 这些全局变量让使用变得更方便，其它包可以定义日志格式就能使用
// 但默认的日志配置和风格也不是所有人都喜欢的，提供个修改的函数以便于使用者修改它
// 在单元测试代码里还演示了如何取消自定义日志打印
func SetLog(zlg *zap.Logger) {
	LOGGER = NewZap(zlg)
	ZAP = LOGGER
	LOG = LOGGER.LOG //底下的函数会依赖于这个变量，它们会重新构造新对象
	SUG = LOGGER.SUG

	LOGS = NewSkipLogs(zlg)
	ZAPS = NewSkipZaps(zlg)
}
