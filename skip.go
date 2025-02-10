package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

type SkipLogs struct {
	Skip0 *zap.Logger
	Skip1 *zap.Logger
	Skip2 *zap.Logger
	Skip3 *zap.Logger
	Skip4 *zap.Logger
	cache *mutexmap.Map[int, *zap.Logger]
}

func NewSkipLogs(zapLog *zap.Logger) *SkipLogs {
	return &SkipLogs{
		Skip0: newSkipDepth(zapLog, 0),
		Skip1: newSkipDepth(zapLog, 1),
		Skip2: newSkipDepth(zapLog, 2),
		Skip3: newSkipDepth(zapLog, 3),
		Skip4: newSkipDepth(zapLog, 4),
		cache: mutexmap.NewMap[int, *zap.Logger](0),
	}
}

func (Z *SkipLogs) Skip(skipDepth int) *zap.Logger {
	switch skipDepth {
	case 0:
		return Z.Skip0
	case 1:
		return Z.Skip1
	case 2:
		return Z.Skip2
	case 3:
		return Z.Skip3
	case 4:
		return Z.Skip4
	default:
		if skipDepth > 0 {
			res, _ := Z.cache.Getset(skipDepth, func() *zap.Logger {
				return newSkipDepth(LOG, skipDepth)
			})
			return res
		} else {
			return Z.Skip0
		}
	}
}

func newSkipDepth(zapLog *zap.Logger, skipDepth int) *zap.Logger {
	return zapLog.WithOptions(zap.AddCallerSkip(skipDepth))
}
