package asserts

import (
	"github.com/stretchr/testify/assert"
)

func SameElements[T comparable](t assert.TestingT, expected, actual []T, extraMsgAndArgs ...any) bool {
	missingMap := make(map[T]int)
	for _, v := range expected {
		missingMap[v]++
	}

	extraMap := make(map[T]int)
	for _, v := range actual {
		if missingMap[v] == 0 {
			extraMap[v]++
			continue
		}
		missingMap[v]--
		if missingMap[v] == 0 {
			delete(missingMap, v)
		}
	}
	var failed bool
	if len(missingMap) != 0 {
		t.Errorf("Not all expected items found in actual:\n\t%+v\n", missingMap)
		failed = true
	}

	if len(extraMap) != 0 {
		t.Errorf("Extra items found in actual that were not expected:\n\t%+v\n", extraMap)
		failed = true
	}
	if !failed || len(extraMsgAndArgs) == 0 {
		return !failed
	}
	format, ok := extraMsgAndArgs[0].(string)
	if !ok {
		t.Errorf("expecting first extra argument to be format but found %T: %v", extraMsgAndArgs[0], extraMsgAndArgs[0])
	}
	t.Errorf(format, extraMsgAndArgs[1:]...)
	return !failed
}

func SameKeyValues[T, U comparable](t assert.TestingT, expected, actual map[T]U, extraMsgAndArgs ...any) {
	missingMap := make(map[T]U)
	for k, v := range expected {
		missingMap[k] = v
	}

	extraMap := make(map[T]U)
	for k, v := range actual {
		ev, ok := missingMap[k]
		if !ok || ev != v {
			extraMap[k] = v
			continue
		}
		delete(missingMap, k)
	}
	var failed bool
	if len(missingMap) != 0 {
		t.Errorf("Not all expected items found in actual:\n\t%+v\n", missingMap)
		failed = true
	}

	if len(extraMap) != 0 {
		t.Errorf("Extra items found in actual that were not expected:\n\t%+v\n", extraMap)
		failed = true
	}
	if !failed || len(extraMsgAndArgs) == 0 {
		return
	}
	format, ok := extraMsgAndArgs[0].(string)
	if !ok {
		t.Errorf("expecting first extra argument to be format but found %T: %v", extraMsgAndArgs[0], extraMsgAndArgs[0])
	}
	t.Errorf(format, extraMsgAndArgs[1:]...)
}
