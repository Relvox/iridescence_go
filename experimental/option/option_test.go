package option_test

import (
	"fmt"
	"testing"

	. "github.com/relvox/iridescence_go/experimental/option"
)

type dummy struct {
	name string
	num  int
}

func Test(t *testing.T) {
	t.Run("Sanity", func(t *testing.T) {
		runOptionTest[int](t, 5)
		runOptionTest[uint](t, 5)
		runOptionTest[float64](t, 5)
		runOptionTest[dummy](t, dummy{"A", 1})
		runOptionTest[*dummy](t, &dummy{"A", 1})
		runOptionTest[*dummy](t, nil)
		runOptionTest[any](t, dummy{"A", 1})
		runOptionTest[any](t, &dummy{"A", 1})
		runOptionTest[any](t, nil)
	})
	t.Run("Stuff", func(t *testing.T) {
		var boo, zoo *dummy = &dummy{"boo", 6}, nil
		if !Maybe(boo).HasValue() {
			t.Fatal("Should have value")
		}
		if Maybe(zoo).HasValue() {
			t.Fatal("Should NOT have value")
		}
	})
	t.Run("Stuff_int", func(t *testing.T) {
		var boo, zoo int = 7, 0
		if !Maybe(boo).HasValue() {
			t.Fatal("Should have value")
		}
		if Maybe(zoo).HasValue() {
			t.Fatal("Should NOT have value")
		}
	})
}

func runOptionTest[OT any](t *testing.T, someValue OT) {
	t.Run(fmt.Sprintf("%T", someValue), func(t *testing.T) {
		o_none := None[OT]()
		o_some := Some[OT](someValue)

		t.Run("HasValue", func(t *testing.T) {
			if o_none.HasValue() {
				t.Fatal("none has value")
			}
			if !o_some.HasValue() {
				t.Fatal("some has no value")
			}
		})

		t.Run("Try", func(t *testing.T) {
			if o_none.Try(func(v OT) {
				t.Fatalf("expected no value but found: %v", v)
			}) {
				t.Fatal("expected Try to return false")
			}
			var stuff bool
			if !o_some.Try(func(v OT) {
				stuff = true
			}) {
				t.Fatal("expected Try to return true")
			}
			if !stuff {
				t.Fatal("expected stuff to be true")
			}
		})

		t.Run("Do", func(t *testing.T) {
			var stuff2 bool
			o_none.Do(func() {
				stuff2 = true
			}, func(v OT) {
				t.Fatal("did not expect to Do some on none")
			})
			if !stuff2 {
				t.Fatal("expected stuff2 to be true")
			}

			var stuff3 bool
			o_some.Do(func() {
				t.Fatal("did not expect to Do none on some")
			}, func(v OT) {
				stuff3 = true
			})
			if !stuff3 {
				t.Fatal("expected stuff3 to be true")
			}
		})

		t.Run("Maybe", func(t *testing.T) {
			opt := Maybe(o_none)
			if opt.HasValue() {
				t.Errorf("Maybe should return None for zero value")
			}

			opt = Maybe(o_some)
			if !opt.HasValue() {
				t.Errorf("Maybe should return Some for non-zero value")
			}
		})
	})
}
