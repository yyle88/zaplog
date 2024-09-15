package zaplog

import (
	"testing"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func TestSetLog(t *testing.T) {
	if false { //由于该操作会影响默认的全局变量，因此这里做个false表示不运行
		//假如你不想让自定义日志打印，也可以这样
		SetLog(zap.NewNop())

		LOG.Debug("abc") //这条日志不会被打印

		LOG.Panic("abc") //这条日志不会被打印，但还是会 panic(ce.Message) 的

		LOG.Panic("abc", zap.Error(errors.New("wrong"))) //这里panic不会打印wrong信息
	}
}

func TestZapLog(t *testing.T) {
	ZAP.LOG.Debug("abc", zap.Int("num", 123))
	ZAP.SUG.Debug("abc", "-|-", "num", "-|-", 123)
}
