package utils

type BitFlag uint8

func (bf *BitFlag) Set(flag BitFlag) {
	*bf |= flag
}

func (bf *BitFlag) SetMany(flags ...BitFlag) {
	for _, flag := range flags {
		*bf |= flag
	}
}

func (bf *BitFlag) Get(flag BitFlag) BitFlag {
	return *bf & flag
}

func (bf *BitFlag) GetMany(flags ...BitFlag) BitFlag {
	var mask BitFlag
	mask.SetMany(flags...)
	return *bf & mask
}

func (bf *BitFlag) Clear(flag BitFlag) {
	*bf &^= flag
}

func (bf *BitFlag) ClearMany(flags ...BitFlag) {
	for _, flag := range flags {
		*bf &^= flag
	}
}

func (bf *BitFlag) ClearAll() {
	*bf = 0
}

func (bf *BitFlag) Toggle(flag BitFlag) {
	*bf ^= flag
}

func (bf *BitFlag) ToggleMany(flags ...BitFlag) {
	for _, flag := range flags {
		*bf ^= flag
	}
}

// Test with multiple flags is effectively TestAny
func (bf *BitFlag) Test(flag BitFlag) bool {
	return *bf&flag != 0
}

func (bf *BitFlag) TestAny(flags ...BitFlag) bool {
	for _, flag := range flags {
		if *bf&flag != 0 {
			return true
		}
	}
	return false
}
func (bf *BitFlag) TestAll(flags ...BitFlag) bool {
	for _, flag := range flags {
		if *bf&flag == 0 {
			return false
		}
	}
	return true
}

// Cleared with multiple flags is effectively ClearedAll
func (bf *BitFlag) Cleared(flag BitFlag) bool {
	return *bf&flag == 0
}

func (bf *BitFlag) ClearedAny(flags ...BitFlag) bool {
	for _, flag := range flags {
		if *bf&flag == 0 {
			return true
		}
	}
	return false
}
func (bf *BitFlag) ClearedAll(flags ...BitFlag) bool {
	for _, flag := range flags {
		if *bf&flag != 0 {
			return false
		}
	}
	return true
}
