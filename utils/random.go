package utils

type Rng interface {
	Int() int
}

func Flip(rng Rng, tWeight, fWeight uint) bool {
	return rng.Int()%(int(tWeight+fWeight)) < int(tWeight)
}

func FairFlip[T any](rng Rng, tVal, fVal T) T {
	if Flip(rng, 100, 100) {
		return tVal
	}
	return fVal
}
