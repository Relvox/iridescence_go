package random

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/rand"

	"github.com/relvox/iridescence_go/random/sources"
)

type NumGen32 struct {
	Source Source32
	Index  uint64
}

func NewNumGen32(source32 Source32) *NumGen32 {
	return &NumGen32{
		Source: source32,
		Index:  0,
	}
}

func NewNumGen32FromIndex(source32 Source32, index uint64) *NumGen32 {
	for i := 0; i < int(index); i++ {
		source32.Uint32()
	}
	return &NumGen32{
		Source: source32,
		Index:  index,
	}
}

func NewNumGen32FromGoSource32(source rand.Source) *NumGen32 {
	return &NumGen32{
		Source: sources.NewGoSource(source),
		Index:  0,
	}
}

func (r *NumGen32) Seed(seed uint32) {
	r.Index = 0
	r.Source.Seed(int64(seed))
}

func (r *NumGen32) SeedFrom(seed uint32, index uint64) {
	r.Index = index
	r.Source.Seed(int64(seed))
	for i := 0; i < int(index); i++ {
		r.Source.Uint32()
	}
}

func (r *NumGen32) Uint32() uint32 {
	r.Index++
	return r.Source.Uint32()
}

func (r *NumGen32) Intn(n int) int {
	r.Index++
	return int(r.Source.Uint32() % uint32(n))
}

func (r *NumGen32) Roll(range_ Range[uint32]) uint32 {
	return range_.Roll(r)
}

func (r *NumGen32) Rolls(ranges ...Range[uint32]) []uint32 {
	result := make([]uint32, len(ranges))
	for i := 0; i < len(result); i++ {
		result[i] = ranges[i].Roll(r)
	}
	return result
}

func (r *NumGen32) GetState() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(r)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	return buf.Bytes()
}

func (r *NumGen32) RestoreState(state []byte) {
	buf := bytes.NewBuffer(state)
	dec := gob.NewDecoder(buf)
	var result *NumGen32
	err := dec.Decode(&result)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	r.Source = result.Source
	r.Index = result.Index
}
