package random

import "strings"

type WordGen = Gen[string]

func FromString(set string) WordGen {
	return func(p IntnGenerator) string {
		index := NewZeroRange(uint32(len(set) - 1)).Roll(p)
		return set[index : index+1]
	}
}

func FromStrings(opts ...string) WordGen {
	return func(p IntnGenerator) string {
		index := NewZeroRange(uint32(len(opts) - 1)).Roll(p)
		return opts[index]
	}
}

func ConcatStr(gens ...WordGen) WordGen {
	return func(p IntnGenerator) string {
		sb := strings.Builder{}
		for _, g := range gens {
			sb.WriteString(g(p))
		}
		return sb.String()
	}
}
