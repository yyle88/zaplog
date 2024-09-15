package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

var ZAPS = NewSkipZaps(LOG)

type SkipZaps struct {
	P0 *Zap
	P1 *Zap
	P2 *Zap
	P3 *Zap
	P4 *Zap
	mp *mutexmap.Map[int, *Zap]
}

func NewSkipZaps(zlg *zap.Logger) *SkipZaps {
	return &SkipZaps{
		P0: NewZapSkip(zlg, 0),
		P1: NewZapSkip(zlg, 1),
		P2: NewZapSkip(zlg, 2),
		P3: NewZapSkip(zlg, 3),
		P4: NewZapSkip(zlg, 4),
		mp: mutexmap.NewMap[int, *Zap](0),
	}
}

func (Z *SkipZaps) Pn(skip int) *Zap {
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
			res, _ := Z.mp.GetOrzSet(skip, func() *Zap {
				return NewZapSkip(LOG, skip)
			})
			return res
		} else {
			return Z.P0
		}
	}
}
