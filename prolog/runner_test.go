package prolog_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/relvox/iridescence_go/prolog"
)

func Test_Append(t *testing.T) {
	runner := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)
	for i := 0; i < 100; i++ {
		runner.AppendReplace(fmt.Sprintf("foo(a, %d).", i))
		res := runner.Query("zoo(A,B).")
		if !assert.Len(t, res, 2) {
			t.FailNow()
		}
	}
}

func Test_Rebuild(t *testing.T) {
	runner := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)
	for i := 0; i < 100; i++ {
		runner.RebuildRunnerWith(fmt.Sprintf("foo(a, %d).", i))
		res := runner.Query("zoo(A,B).")
		if !assert.Len(t, res, 2+i) {
			t.FailNow()
		}
	}
}

func Benchmark_Append(b *testing.B) {
	runner := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runner.AppendReplace(fmt.Sprintf("foo(a, %d).", n))
		runner.Query("zoo(A,B).")
	}
}

func Benchmark_Rebuild(b *testing.B) {
	runner := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runner.RebuildRunnerWith(fmt.Sprintf("foo(a, %d).", n))
		runner.Query("zoo(A,B).")
	}
}
