package random

import (
	"github.com/relvox/iridescence_go/maths"
)

type Range[TNum maths.Number] struct{ Min, Max TNum }

func NewRange[TNum maths.Number](min, max TNum) Range[TNum] {
	return Range[TNum]{Min: min, Max: max}
}

func NewZeroRange[TNum maths.Number](max TNum) Range[TNum] {
	return NewRange[TNum](0, max)
}

func (r Range[TNum]) Roll(rng IntnGenerator) TNum {
	if r.Min == r.Max {
		return r.Min
	}

	if r.Max-r.Min == maths.MaxValue[TNum]() {
		return TNum(rng.Intn(int(maths.MaxValue[TNum]())))
	}
	diff := r.Max - r.Min + 1
	return TNum(rng.Intn(int(diff))) + r.Min
}

func (r Range[TNum]) Gen() Gen[TNum] {
	return func(p IntnGenerator) TNum {
		return r.Roll(p)
	}
}

func (r Range[TNum]) Match(val TNum) bool {
	return r.Min <= val && val <= r.Max
}
