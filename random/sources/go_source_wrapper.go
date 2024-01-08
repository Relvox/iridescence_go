package sources

import (
	"math"
	"math/rand"
)

type GoSource struct {
	rand.Source
	s64         rand.Source64
	quickUint32 int64
}

func NewGoSource(src rand.Source) *GoSource {
	s64, _ := src.(rand.Source64)
	return &GoSource{src, s64, -1}
}

func (r *GoSource) Uint32() uint32 {
	if r.quickUint32 != -1 {
		defer func() { r.quickUint32 = -1 }()
		return uint32(r.quickUint32)
	}
	if r.s64 != nil {
		rr := r.s64.Uint64()
		r.quickUint32 = int64((rr & (1<<32 - 1)))
		rr >>= 32
		return uint32(rr)
	}
	return uint32(r.Int63() >> 31)
}

func (r *GoSource) Uint64() uint64 {
	if r.s64 != nil {
		return r.s64.Uint64()
	}
	return uint64(r.Int63())>>31 | uint64(r.Int63())<<32
}

func (r *GoSource) Intn(n int) int {
	if n > math.MaxInt32 {
		return int(r.Uint64() % uint64(n))
	}
	if r.quickUint32 != -1 {
		defer func() { r.quickUint32 = -1 }()
		return int(r.quickUint32 % int64(n))
	}
	return int(r.Uint32() % uint32(n))
}
