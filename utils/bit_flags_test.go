package utils_test

import (
	"testing"

	"github.com/relvox/iridescence_go/utils"
)

func TestBitFlagComplexScenario(t *testing.T) {
	// Define some fictional flags for the scenario
	const (
		FlagA utils.BitFlag = 1 << iota
		FlagB
		FlagC
		FlagD
		FlagE // Extra flag to test edge cases
		FlagF // Additional flag for more coverage
	)

	var bf utils.BitFlag

	// Initial setting of individual flags
	bf.Set(FlagA)
	if !bf.Test(FlagA) {
		t.Error("Expected FlagA to be set")
	}

	// Setting multiple flags and checking
	bf.SetMany(FlagB, FlagC)
	if !bf.TestAll(FlagA, FlagB, FlagC) {
		t.Error("Expected FlagA, FlagB, and FlagC to be set")
	}

	// Clearing a flag and checking
	bf.Clear(FlagA)
	if bf.Test(FlagA) {
		t.Error("Expected FlagA to be cleared")
	}

	// Clearing multiple flags
	bf.ClearMany(FlagB, FlagC)
	if bf.TestAny(FlagA, FlagB, FlagC) {
		t.Error("Expected FlagA, FlagB, and FlagC to be cleared")
	}

	// Toggling flags
	bf.ToggleMany(FlagA, FlagB)
	if !bf.TestAny(FlagA, FlagB) {
		t.Error("Expected either FlagA or FlagB to be toggled on")
	}

	// Edge case: Toggling an unset flag and a set flag
	bf.ToggleMany(FlagB, FlagC)
	if bf.Test(FlagB) || !bf.Test(FlagC) {
		t.Error("Expected FlagB to be off and FlagC to be on after toggling")
	}

	// Edge case: Setting a flag that is already set
	bf.Set(FlagC)
	if !bf.Test(FlagC) {
		t.Error("Expected FlagC to remain set")
	}

	// Using Get and GetMany
	combined := bf.Get(FlagD)
	combinedMany := bf.GetMany(FlagD, FlagE)
	if !combined.Test(FlagD) || !combinedMany.TestAll(FlagD, FlagE) {
		t.Error("Expected combined flags to include FlagD and combinedMany to include FlagD, FlagE")
	}

	// Clearing all flags
	bf.ClearAll()
	if bf.TestAny(FlagA, FlagB, FlagC, FlagD, FlagE) {
		t.Error("Expected all flags to be cleared")
	}

	// Edge case: Testing Cleared and ClearedAny on cleared flags
	if !bf.Cleared(FlagA) || !bf.ClearedAny(FlagA, FlagB) {
		t.Error("Expected FlagA to be cleared and at least one of FlagA or FlagB to be cleared")
	}

	// Edge case: Testing ClearedAll when all flags are cleared
	if !bf.ClearedAll(FlagA, FlagB, FlagC, FlagD, FlagE) {
		t.Error("Expected all flags to be cleared")
	}

	// Additional edge cases and use cases

	// Setting a flag with a value outside the defined constants
	bf.Set(FlagF)
	if !bf.Test(FlagF) {
		t.Error("Expected FlagF to be set")
	}

	// Passing an empty variadic slice to SetMany and ClearMany
	bf.SetMany()
	bf.ClearMany()
	// No assert needed; this is to ensure no panic or error occurs

	// Toggling flags with an empty variadic slice
	bf.ToggleMany() // No assert needed; this is to ensure no panic or error occurs

	// Checking TestAny and TestAll with an empty variadic slice
	if bf.TestAny() {
		t.Error("Expected TestAny with no arguments to return false")
	}
	if !bf.TestAll() {
		t.Error("Expected TestAll with no arguments to return true")
	}

	// Checking ClearedAny and ClearedAll with an empty variadic slice
	if bf.ClearedAny() {
		t.Error("Expected ClearedAny with no arguments to return false")
	}
	if !bf.ClearedAll() {
		t.Error("Expected ClearedAll with no arguments to return true")
	}
}

func TestBitFlagHas(t *testing.T) {
	var bf utils.BitFlag = Flag1
	if !bf.TestAll(Flag1) {
		t.Errorf("Expected Flag1 to be set")
	}
	if bf.TestAll(Flag2) {
		t.Errorf("Expected Flag2 to not be set")
	}
}

func TestBitFlagSetVariadic(t *testing.T) {
	var bf utils.BitFlag
	bf.SetMany(Flag1, Flag2)
	if !bf.TestAll(Flag1) || !bf.TestAll(Flag2) {
		t.Errorf("Expected Flag1 and Flag2 to be set")
	}
	if bf.TestAll(Flag4) {
		t.Errorf("Expected Flag4 to not be set")
	}
}

func TestBitFlagClearVariadic(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2 | Flag4
	bf.ClearMany(Flag1, Flag4)
	if bf.TestAll(Flag1) || bf.TestAll(Flag4) {
		t.Errorf("Expected Flag1 and Flag4 to be cleared")
	}
	if !bf.TestAll(Flag2) {
		t.Errorf("Expected Flag2 to remain set")
	}
}

func TestBitFlagToggleVariadic(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	bf.ToggleMany(Flag1, Flag4)
	if bf.TestAll(Flag1) {
		t.Errorf("Expected Flag1 to be toggled off")
	}
	if !bf.TestAll(Flag2) {
		t.Errorf("Expected Flag2 to remain on")
	}
	if !bf.TestAll(Flag4) {
		t.Errorf("Expected Flag4 to be toggled on")
	}
}

func TestBitFlagHasVariadic(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	if !bf.TestAll(Flag1, Flag2) {
		t.Errorf("Expected Flag1 and Flag2 to be set")
	}
	if bf.TestAll(Flag1, Flag4) {
		t.Errorf("Expected combination of Flag1 and Flag4 to not be set")
	}
}

func TestBitFlagSet(t *testing.T) {
	var bf utils.BitFlag
	bf.Set(Flag1)
	if !bf.Test(Flag1) {
		t.Errorf("Expected Flag1 to be set")
	}
}

func TestBitFlagSetMany(t *testing.T) {
	var bf utils.BitFlag
	bf.SetMany(Flag1, Flag2)
	if !bf.TestAll(Flag1, Flag2) {
		t.Errorf("Expected Flag1 and Flag2 to be set")
	}
}

func TestBitFlagGet(t *testing.T) {
	var bf utils.BitFlag = Flag1
	result := bf.Get(Flag2)
	if !result.Test(Flag1) || !result.Test(Flag2) {
		t.Errorf("Expected Flag1 and Flag2 to be present in the result")
	}
}

func TestBitFlagGetMany(t *testing.T) {
	var bf utils.BitFlag = Flag1
	result := bf.GetMany(Flag2, Flag4)
	if !result.TestAll(Flag1, Flag2, Flag4) {
		t.Errorf("Expected Flag1, Flag2, and Flag4 to be present in the result")
	}
}

func TestBitFlagClear(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	bf.Clear(Flag1)
	if bf.Test(Flag1) {
		t.Errorf("Expected Flag1 to be cleared")
	}
}

func TestBitFlagClearMany(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2 | Flag4
	bf.ClearMany(Flag1, Flag4)
	if bf.TestAny(Flag1, Flag4) {
		t.Errorf("Expected Flag1 and Flag4 to be cleared")
	}
	if !bf.Test(Flag2) {
		t.Errorf("Expected Flag2 to remain set")
	}
}

func TestBitFlagToggle(t *testing.T) {
	var bf utils.BitFlag = Flag1
	bf.Toggle(Flag1)
	if bf.Test(Flag1) {
		t.Errorf("Expected Flag1 to be toggled off")
	}
	bf.Toggle(Flag1)
	if !bf.Test(Flag1) {
		t.Errorf("Expected Flag1 to be toggled on")
	}
}

func TestBitFlagToggleMany(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	bf.ToggleMany(Flag1, Flag4)
	if bf.Test(Flag1) {
		t.Errorf("Expected Flag1 to be toggled off")
	}
	if !bf.Test(Flag2) {
		t.Errorf("Expected Flag2 to remain on")
	}
	if !bf.Test(Flag4) {
		t.Errorf("Expected Flag4 to be toggled on")
	}
}

func TestBitFlagTest(t *testing.T) {
	var bf utils.BitFlag = Flag1
	if !bf.Test(Flag1) {
		t.Errorf("Expected Flag1 to be set")
	}
	if bf.Test(Flag2) {
		t.Errorf("Expected Flag2 to not be set")
	}
}

func TestBitFlagTestAny(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	if !bf.TestAny(Flag1, Flag4) {
		t.Errorf("Expected either Flag1 or Flag4 to be set")
	}
	if bf.TestAny(Flag4, Flag8) {
		t.Errorf("Expected neither Flag4 nor Flag8 to be set")
	}
}

func TestBitFlagTestAll(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	if !bf.TestAll(Flag1, Flag2) {
		t.Errorf("Expected both Flag1 and Flag2 to be set")
	}
	if bf.TestAll(Flag1, Flag4) {
		t.Errorf("Expected not all of Flag1 and Flag4 to be set")
	}
}

func TestBitFlagCleared(t *testing.T) {
	var bf utils.BitFlag = Flag1
	if !bf.Cleared(Flag2) {
		t.Errorf("Expected Flag2 to be cleared")
	}
	if bf.Cleared(Flag1) {
		t.Errorf("Expected Flag1 to not be cleared")
	}
}

func TestBitFlagClearedAny(t *testing.T) {
	var bf utils.BitFlag = Flag1 | Flag2
	if !bf.ClearedAny(Flag4, Flag8) {
		t.Errorf("Expected at least one of Flag4 or Flag8 to be cleared")
	}
	if bf.ClearedAny(Flag1, Flag2) {
		t.Errorf("Expected neither Flag1 nor Flag2 to be cleared")
	}
}

func TestBitFlagClearedAll(t *testing.T) {
	var bf utils.BitFlag = Flag1
	if !bf.ClearedAll(Flag2, Flag4) {
		t.Errorf("Expected all of Flag2 and Flag4 to be cleared")
	}
	if bf.ClearedAll(Flag1, Flag4) {
		t.Errorf("Expected not all of Flag1 and Flag4 to be cleared")
	}
}
