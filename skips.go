package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

type SkipZaps struct {
	P0 *Zap
	P1 *Zap
	P2 *Zap
	P3 *Zap
	P4 *Zap
	mp *mutexmap.Map[int, *Zap]
}

func NewSkipZaps(zapLog *zap.Logger) *SkipZaps {
	return &SkipZaps{
		P0: NewZapSkip(zapLog, 0),
		P1: NewZapSkip(zapLog, 1),
		P2: NewZapSkip(zapLog, 2),
		P3: NewZapSkip(zapLog, 3),
		P4: NewZapSkip(zapLog, 4),
		mp: mutexmap.NewMap[int, *Zap](0),
	}
}

func (Z *SkipZaps) Pn(skipDepth int) *Zap {
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
			res, _ := Z.mp.Getset(skipDepth, func() *Zap {
				return NewZapSkip(LOG, skipDepth)
			})
			return res
		} else {
			return Z.P0
		}
	}
}
