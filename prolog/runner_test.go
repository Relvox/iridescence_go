package prolog_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	asserts "github.com/relvox/iridescence_go/assert"
	"github.com/relvox/iridescence_go/prolog"
)

func Test_Append(t *testing.T) {
	runner, err := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)

	if err != nil {
		t.Failed()
	}
	for i := 0; i < 100; i++ {
		runner.AppendReplace(fmt.Sprintf("foo(a, %d).", i))
		res, err := runner.Query("zoo(A,B).")
		if !assert.Len(t, res, 2) || err != nil {
			t.FailNow()
		}
	}
}

func Test_Rebuild(t *testing.T) {
	runner, err := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)

	if err != nil {
		t.Failed()
	}
	for i := 0; i < 100; i++ {
		runner.RebuildRunnerWith(fmt.Sprintf("foo(a, %d).", i))
		res, err := runner.Query("zoo(A,B).")
		if !assert.Len(t, res, 2+i) || err != nil {
			t.FailNow()
		}
	}
}

func Benchmark_Append(b *testing.B) {
	runner, _ := prolog.NewRunner(`
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
	runner, _ := prolog.NewRunner(`
		:- discontiguous(foo/2).
		zoo(a, 666).
		zoo(X, Y) :- foo(X, Y).`)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runner.RebuildRunnerWith(fmt.Sprintf("foo(a, %d).", n))
		runner.Query("zoo(A,B).")
	}
}

func Test_TypedQueries(t *testing.T) {
	db := `
	foo(1, 'a').
	foo(2, 'b').
	foo(3, 'c').
	`

	type Foo struct {
		Id  int
		Tok string
	}

	expected := []Foo{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	r, err := prolog.NewRunner(db)
	if err != nil {
		panic(err)
	}
	t.Run("typed list", func(t *testing.T) {
		actual, err := prolog.TypedList[Foo](r)
		if err != nil {
			panic(err)
		}
		asserts.SameElements(t, expected, actual)
	})
	t.Run("filtered typed list", func(t *testing.T) {
		actual, err := prolog.TypedListFilter[Foo](r, "Id =:= 2")
		if err != nil {
			panic(err)
		}
		asserts.SameElements(t, expected[1:2], actual)
	})
	t.Run("insert object", func(t *testing.T) {
		actual, err := prolog.TypedList[Foo](r)
		if err != nil {
			panic(err)
		}
		asserts.SameElements(t, expected, actual)
		err = prolog.InsertObjects(r, Foo{4, "d"})
		if err != nil {
			panic(err)
		}
		actual, err = prolog.TypedList[Foo](r)
		if err != nil {
			panic(err)
		}
		expected = append(expected, Foo{4, "d"})
		asserts.SameElements(t, expected, actual)
	})
	t.Run("insert objects", func(t *testing.T) {
		actual, err := prolog.TypedList[Foo](r)
		if err != nil {
			panic(err)
		}
		asserts.SameElements(t, expected, actual)
		err = prolog.InsertObjects(r, Foo{5, "e"}, Foo{6, "f"})
		if err != nil {
			panic(err)
		}
		actual, err = prolog.TypedList[Foo](r)
		if err != nil {
			panic(err)
		}
		expected = append(expected, Foo{5, "e"})
		expected = append(expected, Foo{6, "f"})
		asserts.SameElements(t, expected, actual)
	})

}
