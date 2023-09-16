package namegen

import "math/rand"

type RandIntn func(n int) int
type Generator func(r RandIntn) string

func Blank(r RandIntn) string {
	return ""
}

func Concat(gens ...Generator) Generator {
	return func(r RandIntn) string {
		result := ""
		for _, gen := range gens {
			result += gen(r)
		}
		return result
	}
}

func Oneof(gens ...Generator) Generator {
	return func(r RandIntn) string {
		return gens[r(len(gens))](r)
	}
}

func Letters(letters string) Generator {
	return func(r RandIntn) string {
		i := r(len(letters))
		return letters[i : i+1]
	}
}

func FromWords(words ...string) Generator {
	return func(r RandIntn) string {
		i := r(len(words))
		return words[i]
	}
}

var (
	SampleLettersC  = Letters("smplnkbzdfgmpktrln")
	SampleLettersV  = Letters("aaeeoouieaouieay")
	SampleSyllableC = Concat(SampleLettersC, SampleLettersV, SampleLettersC)
	SampleSyllableO = Concat(SampleLettersC, SampleLettersV)
	sampleChoice    = Oneof(SampleSyllableC, SampleSyllableC, SampleSyllableO, SampleSyllableO, Blank)
	SampleName      = Concat(sampleChoice, sampleChoice, sampleChoice)

	Sample2LettersC  = Letters("mpnkbdgpktr")
	Sample2LettersV  = Letters("aaeoouieaou")
	Sample2SyllableC = Concat(Sample2LettersC, Sample2LettersV, Sample2LettersC)
	Sample2SyllableO = Concat(Sample2LettersC, Sample2LettersV)
	sample2Choice    = Oneof(Sample2SyllableC, Sample2SyllableO, Sample2SyllableO)
	Sample2Name      = Concat(sample2Choice, sample2Choice)
)

func Rng(seed int64) RandIntn {
	src := rand.NewSource(seed)
	return rand.New(src).Intn
}
