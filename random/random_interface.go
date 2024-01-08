package random

type Source32 interface {
	Uint32() uint32
	Seed(seed int64)
}

type IntnGenerator interface {
	Intn(n int) int
}

func Flip(rng IntnGenerator, tWeight, fWeight int) bool {
	return rng.Intn(tWeight+fWeight) < tWeight
}

func FairFlip[T any](rng IntnGenerator, tVal, fVal T) T {
	if Flip(rng, 100, 100) {
		return tVal
	}
	return fVal
}

type Uint32Generator interface {
	Uint32() uint32
}

type Uint64Generator interface {
	Uint64() uint64
}

type Int63Generator interface {
	Int63() int64
}
