package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

var LOGS = NewSkipLogs()

type skipLogs struct {
	P0 *zap.Logger
	P1 *zap.Logger
	P2 *zap.Logger
	P3 *zap.Logger
	P4 *zap.Logger
	mp *mutexmap.Map[int, *zap.Logger]
}

func NewSkipLogs() *skipLogs {
	return &skipLogs{
		P0: newSkipLog(LOG, 0),
		P1: newSkipLog(LOG, 1),
		P2: newSkipLog(LOG, 2),
		P3: newSkipLog(LOG, 3),
		P4: newSkipLog(LOG, 4),
		mp: mutexmap.NewMap[int, *zap.Logger](0),
	}
}

func (Z *skipLogs) Pn(skip int) *zap.Logger {
	switch skip {
	case 0:
		return Z.P0
	case 1:
		return Z.P1
	case 2:
		return Z.P2
	case 3:
		return Z.P3
	case 4:
		return Z.P4
	default:
		if skip > 0 {
			res, _ := Z.mp.GetOrzSet(skip, func() *zap.Logger {
				return newSkipLog(LOG, skip)
			})
			return res
		} else {
			return Z.P0
		}
	}
}

func newSkipLog(parent *zap.Logger, skip int) *zap.Logger {
	return parent.WithOptions(zap.AddCallerSkip(skip))
}
