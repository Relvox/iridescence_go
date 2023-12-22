package utils_test

import (
	"testing"

	"github.com/relvox/iridescence_go/utils"
	"github.com/stretchr/testify/assert"
)

const (
	Flag1 utils.BitFlag = 1 << iota
	Flag2
	Flag4
	Flag8
	Flag16
	Flag32
	Flag64
	Flag128
	AllFlags = Flag1 | Flag2 | Flag4 | Flag8 | Flag16 | Flag32 | Flag64 | Flag128
)

var allFlagsSlice = []utils.BitFlag{Flag1, Flag2, Flag4, Flag8, Flag16, Flag32, Flag64, Flag128}

func Benchmark_BitFlag_Operations(b *testing.B) {
	b.Run("Set", func(b *testing.B) {
		var bf utils.BitFlag
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bf.Set(Flag16)
		}
		b.StopTimer()
		assert.Equal(b, Flag16, bf)
	})

	b.Run("SetMany", func(b *testing.B) {
		var bf utils.BitFlag
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bf.SetMany(allFlagsSlice...)
		}
		b.StopTimer()
		assert.Equal(b, AllFlags, bf)
	})

	b.Run("Clear", func(b *testing.B) {
		var bf utils.BitFlag = AllFlags
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bf.Clear(Flag16)
		}
		b.StopTimer()
		assert.Equal(b, AllFlags&^Flag16, bf)
	})

	b.Run("ClearMany", func(b *testing.B) {
		var bf utils.BitFlag = AllFlags
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bf.ClearMany(allFlagsSlice...)
		}
		b.StopTimer()
		assert.Equal(b, utils.BitFlag(0), bf)
	})

	b.Run("Toggle", func(b *testing.B) {
		var bf utils.BitFlag
		for i := 0; i < b.N; i++ {
			bf.Toggle(Flag16)
		}
		b.StopTimer()
		var expect utils.BitFlag
		if b.N%2 != 0 {
			expect = Flag16
		}
		assert.Equal(b, expect, bf)
	})

	b.Run("ToggleMany", func(b *testing.B) {
		var bf utils.BitFlag
		for i := 0; i < b.N; i++ {
			bf.ToggleMany(allFlagsSlice...)
		}
		b.StopTimer()
		var expect utils.BitFlag
		if b.N%2 != 0 {
			expect = AllFlags
		}
		assert.Equal(b, expect, bf)
	})

	b.Run("Has", func(b *testing.B) {
		var bf utils.BitFlag = AllFlags
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = bf.Test(Flag16)
		}
		b.StopTimer()
		assert.True(b, bf.Test(Flag16))
	})

	b.Run("NotHas", func(b *testing.B) {
		var bf utils.BitFlag = Flag16
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = bf.Cleared(Flag2)
		}
		b.StopTimer()
		assert.True(b, bf.Cleared(Flag2))
	})

	b.Run("HasAny", func(b *testing.B) {
		var bf utils.BitFlag = Flag16 | Flag32
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = bf.TestAny(Flag16, Flag64)
		}
		b.StopTimer()
		assert.True(b, bf.TestAny(Flag16, Flag64))
	})

	b.Run("NotHasAny", func(b *testing.B) {
		var bf utils.BitFlag = Flag16
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = bf.ClearedAny(Flag32, Flag64)
		}
		b.StopTimer()
		assert.True(b, bf.ClearedAny(Flag32, Flag64))
	})

	b.Run("HasAll", func(b *testing.B) {
		var bf utils.BitFlag = Flag16 | Flag32
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = bf.TestAll(Flag16, Flag32)
		}
		b.StopTimer()
		assert.True(b, bf.TestAll(Flag16, Flag32))
	})

	b.Run("NotHasAll", func(b *testing.B) {
		var bf utils.BitFlag = Flag16 | Flag32
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = bf.ClearedAll(Flag16, Flag64)
		}
		b.StopTimer()
		assert.False(b, bf.ClearedAll(Flag16, Flag64))
	})
}
