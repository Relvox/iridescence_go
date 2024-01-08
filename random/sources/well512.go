package sources

import "math"

type WELL512 struct {
	State  [16]uint32
	Cursor uint8
}

func NewWELL512(seed uint32) *WELL512 {
	w := &WELL512{
		State: [16]uint32{},
	}
	w.Seed(int64(seed))
	return w
}

func (w *WELL512) Seed(seed int64) {
	w.Cursor = 0
	w.State[0] = uint32(seed)
	for i := 1; i < 16; i++ {
		w.State[i] = (1812433253*(w.State[i-1]^(w.State[i-1]>>30)) + uint32(i))
	}
}

func (w *WELL512) Uint32() uint32 {
	a, b, c, d := w.State[w.Cursor], w.State[(w.Cursor+13)&15], w.State[(w.Cursor+9)&15], w.State[(w.Cursor+5)&15]
	y := a ^ (a << 16) ^ b ^ (b << 15)
	z := c ^ (c >> 11) ^ d ^ (d >> 2)
	w.State[w.Cursor] = y ^ z
	w.Cursor = (w.Cursor + 15) & 15
	return w.State[w.Cursor]
}

func (w *WELL512) Uint64() uint64 {
	return uint64(w.Uint32())<<32 | uint64(w.Uint32())
}

func (w *WELL512) Int63() int64 {
	return int64(w.Uint64() >> 1)
}

func (w *WELL512) Intn(n int) int {
	if n > math.MaxInt32 {
		return int(w.Uint64() % uint64(n))
	}
	return int(w.Uint32() % uint32(n))
}
