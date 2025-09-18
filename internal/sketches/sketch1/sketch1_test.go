package sketch1

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func TestExample(t *testing.T) {
	//假如你不想让自定义日志打印，也可以这样
	zaplog.SetLog(zap.NewNop())

	zaplog.LOG.Debug("abc") //这条日志不会被打印

	require.Panics(t, func() {
		zaplog.LOG.Panic("abc") //这条日志不会被打印，但还是会 panic(ce.Message) 的
	})

	require.Panics(t, func() {
		zaplog.LOG.Panic("abc", zap.Error(errors.New("wrong"))) //这里panic不会打印wrong信息
	})
}
