package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

type SkipZaps struct {
	Skip0 *Zap
	Skip1 *Zap
	Skip2 *Zap
	Skip3 *Zap
	Skip4 *Zap
	cache *mutexmap.Map[int, *Zap]
}

func NewSkipZaps(zapLog *zap.Logger) *SkipZaps {
	return &SkipZaps{
		Skip0: NewZapSkip(zapLog, 0),
		Skip1: NewZapSkip(zapLog, 1),
		Skip2: NewZapSkip(zapLog, 2),
		Skip3: NewZapSkip(zapLog, 3),
		Skip4: NewZapSkip(zapLog, 4),
		cache: mutexmap.NewMap[int, *Zap](0),
	}
}

func (Z *SkipZaps) Skip(skipDepth int) *Zap {
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
			res, _ := Z.cache.Getset(skipDepth, func() *Zap {
				return NewZapSkip(LOG, skipDepth)
			})
			return res
		} else {
			return Z.Skip0
		}
	}
}
