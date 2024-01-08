package random_test

import (
	"testing"

	"github.com/relvox/iridescence_go/random"
	"github.com/relvox/iridescence_go/random/sources"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name        string
		rangeMin    uint32
		rangeMax    uint32
		testValue   uint32
		expectRoll  uint32
		expectMatch bool
	}{
		{"Normal Case", 10, 20, 15, 14, true},
		{"Minimum Edge Case", 0, 10, 0, 0, true},
		{"Maximum Edge Case", 0, 10, 10, 9, true},
		{"Below Range", 10, 20, 9, 17, false},
		{"Above Range", 10, 20, 21, 14, false},
		{"Uint32 Max Value", 0, 4294967295, 4294967295, 4236421601, true},
		{"Uint32 Min Value", 0, 4294967295, 0, 828186546, true},
		{"Overflow Case 1", 4294967294, 4294967295, 4294967294, 4294967294, true},
		{"Overflow Case 2", 4294967294, 4294967295, 4294967295, 4294967295, true},
		{"Overflow Case 3", 4294967293, 4294967295, 4294967294, 4294967293, true},
		{"Overflow Case 4", 4294967293, 4294967295, 4294967295, 4294967293, true},
		{"Match with Min Equal to Max", 10, 10, 10, 10, true},
		{"Mismatch with Min Equal to Max", 10, 10, 11, 10, false},
	}

	rng := sources.NewWELL512(123)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := random.Range[uint32]{Min: tt.rangeMin, Max: tt.rangeMax}

			// Test Roll method
			if roll := r.Roll(rng); roll != tt.expectRoll {
				t.Errorf("expected roll to be %d, got %d", tt.expectRoll, roll)
			}

			// Test Match method
			if match := r.Match(tt.testValue); match != tt.expectMatch {
				t.Errorf("expected match to be %v, got %v", tt.expectMatch, match)
			}
		})
	}
}
