package maths_test

import (
	"log"
	"math"
	"strings"
	"testing"

	"github.com/relvox/iridescence_go/maths"
	"github.com/relvox/iridescence_go/utils"
)

// Tests for MaxValue
func TestMaxValue(t *testing.T) {
	tests := []struct {
		name string
		want any
		got  any
	}{
		{"Max uint8", uint8(math.MaxUint8), maths.MaxValue[uint8]()},
		{"Max uint16", uint16(math.MaxUint16), maths.MaxValue[uint16]()},
		{"Max uint32", uint32(math.MaxUint32), maths.MaxValue[uint32]()},
		{"Max uint64", uint64(math.MaxUint64), maths.MaxValue[uint64]()},
		{"Max uint", uint(math.MaxUint64), maths.MaxValue[uint]()},
		{"Max int8", int8(math.MaxInt8), maths.MaxValue[int8]()},
		{"Max int16", int16(math.MaxInt16), maths.MaxValue[int16]()},
		{"Max int32", int32(math.MaxInt32), maths.MaxValue[int32]()},
		{"Max int64", int64(math.MaxInt64), maths.MaxValue[int64]()},
		{"Max int", int(math.MaxInt64), maths.MaxValue[int]()},
		{"Max float32", float32(math.MaxFloat32), maths.MaxValue[float32]()},
		{"Max float64", float64(math.MaxFloat64), maths.MaxValue[float64]()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("MaxValue() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

// Tests for MinValue
func TestMinValue(t *testing.T) {
	tests := []struct {
		name string
		want any
		got  any
	}{
		{"Min uint8", uint8(0), maths.MinValue[uint8]()},
		{"Min uint16", uint16(0), maths.MinValue[uint16]()},
		{"Min uint32", uint32(0), maths.MinValue[uint32]()},
		{"Min uint64", uint64(0), maths.MinValue[uint64]()},
		{"Min uint", uint(0), maths.MinValue[uint]()},
		{"Min int8", int8(math.MinInt8), maths.MinValue[int8]()},
		{"Min int16", int16(math.MinInt16), maths.MinValue[int16]()},
		{"Min int32", int32(math.MinInt32), maths.MinValue[int32]()},
		{"Min int64", int64(math.MinInt64), maths.MinValue[int64]()},
		{"Min int", int(math.MinInt64), maths.MinValue[int]()},
		{"Min float32", float32(-math.MaxFloat32), maths.MinValue[float32]()},
		{"Min float64", float64(-math.MaxFloat64), maths.MinValue[float64]()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("MinValue() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func Test_FOO(t *testing.T) {
	log.Println(strings.Join(utils.Repeat("|------|", 8), ""))
	log.Printf("%065b", math.MinInt8)
	log.Printf("%065b", maths.MinValue[int8]())
	log.Println(strings.Join(utils.Repeat("|------|", 8), ""))
	log.Printf("%065b", math.MinInt16)
	log.Printf("%065b", maths.MinValue[int16]())
	log.Println(strings.Join(utils.Repeat("|------|", 8), ""))
	log.Printf("%065b", math.MinInt32)
	log.Printf("%065b", maths.MinValue[int32]())
	log.Println(strings.Join(utils.Repeat("|------|", 8), ""))
	log.Printf("%065b", math.MinInt64)
	log.Printf("%065b", maths.MinValue[int64]())
	log.Println(strings.Join(utils.Repeat("|------|", 8), ""))
	log.Printf("%065b", math.MinInt)
	log.Printf("%065b", maths.MinValue[int]())
	log.Println(strings.Join(utils.Repeat("|------|", 8), ""))
}
