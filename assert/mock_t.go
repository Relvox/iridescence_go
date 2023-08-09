package assert

import (
	"fmt"
	"testing"

	"github.com/relvox/iridescence_go/utils"
	"golang.org/x/exp/maps"
)

type MockT struct {
	Errors utils.Set[string]
}

func NewMockT() *MockT {
	return &MockT{
		Errors: make(utils.Set[string]),
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
