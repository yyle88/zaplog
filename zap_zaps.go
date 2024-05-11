package zaplog

import (
	"github.com/yyle88/mutexmap"
	"go.uber.org/zap"
)

var ZAPS = NewSkipZaps()

type skipZaps struct {
	P0 *ZapTuple
	P1 *ZapTuple
	P2 *ZapTuple
	P3 *ZapTuple
	P4 *ZapTuple
	mp *mutexmap.Map[int, *ZapTuple]
}

func NewSkipZaps() *skipZaps {
	return &skipZaps{
		P0: newSkipZapTuple(LOG, 0),
		P1: newSkipZapTuple(LOG, 1),
		P2: newSkipZapTuple(LOG, 2),
		P3: newSkipZapTuple(LOG, 3),
		P4: newSkipZapTuple(LOG, 4),
		mp: mutexmap.NewMap[int, *ZapTuple](0),
	}
}

func (Z *skipZaps) Pn(skip int) *ZapTuple {
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
			res, _ := Z.mp.GetOrzSet(skip, func() *ZapTuple {
				return newSkipZapTuple(LOG, skip)
			})
			return res
		} else {
			return Z.P0
		}
	}
}

func newSkipZapTuple(parent *zap.Logger, skip int) *ZapTuple {
	zlg := parent.WithOptions(zap.AddCallerSkip(skip))
	return &ZapTuple{
		LOG: zlg,
		SUG: zlg.Sugar(),
	}
}
