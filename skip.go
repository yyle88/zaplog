package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

type SkipLogs struct {
	P0 *zap.Logger
	P1 *zap.Logger
	P2 *zap.Logger
	P3 *zap.Logger
	P4 *zap.Logger
	mp *mutexmap.Map[int, *zap.Logger]
}

func NewSkipLogs(zapLog *zap.Logger) *SkipLogs {
	return &SkipLogs{
		P0: newSkipDepth(zapLog, 0),
		P1: newSkipDepth(zapLog, 1),
		P2: newSkipDepth(zapLog, 2),
		P3: newSkipDepth(zapLog, 3),
		P4: newSkipDepth(zapLog, 4),
		mp: mutexmap.NewMap[int, *zap.Logger](0),
	}
}

func (Z *SkipLogs) Pn(skipDepth int) *zap.Logger {
	switch skipDepth {
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
		if skipDepth > 0 {
			res, _ := Z.mp.Getset(skipDepth, func() *zap.Logger {
				return newSkipDepth(LOG, skipDepth)
			})
			return res
		} else {
			return Z.P0
		}
	}
}

func newSkipDepth(zapLog *zap.Logger, skipDepth int) *zap.Logger {
	return zapLog.WithOptions(zap.AddCallerSkip(skipDepth))
}
