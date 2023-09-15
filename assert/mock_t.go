package assert

import (
	"fmt"
	"testing"

	"golang.org/x/exp/maps"
	
	"github.com/relvox/iridescence_go/sets"
)

type MockT struct {
	Errors sets.Set[string]
}

func NewMockT() *MockT {
	return &MockT{
		Errors: make(sets.Set[string]),
	}
}

func (m *MockT) Errorf(format string, args ...any) {
	text := fmt.Sprintf(format, args...)
	m.Errors.Add(text)
}

func (m *MockT) Assert(t *testing.T, expectedErrors ...string) {
	errorValues := maps.Keys(m.Errors)
	SameElements(t, expectedErrors, errorValues)
}
