package option_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/relvox/iridescence_go/experimental/option"
)

func TestResult(t *testing.T) {
	t.Run("Sanity", func(t *testing.T) {
		runResultTest[int](t, 5, errors.New("error"))
		runResultTest[uint](t, 5, errors.New("error"))
		runResultTest[float64](t, 5, errors.New("error"))
		runResultTest[dummy](t, dummy{"A", 1}, errors.New("error"))
		runResultTest[*dummy](t, &dummy{"A", 1}, errors.New("error"))
	})
}

func runResultTest[TOk any](t *testing.T, okValue TOk, errValue error) {
	t.Run(fmt.Sprintf("%T/%T", okValue, errValue), func(t *testing.T) {
		resOk := Ok[TOk](okValue)
		resErr := Err[TOk](errValue)

		if !resOk.IsOk() {
			t.Fatal("Ok result should be Ok")
		}
		if resOk.IsErr() {
			t.Fatal("Ok result should not be Err")
		}
		if !resErr.IsErr() {
			t.Fatal("Err result should be Err")
		}
		if resErr.IsOk() {
			t.Fatal("Err result should not be Ok")
		}

		if !resOk.Try(func(v TOk) {}) {
			t.Fatal("Try should succeed on Ok result")
		}
		if resErr.Try(func(v TOk) {}) {
			t.Fatal("Try should fail on Err result")
		}

		var handled bool
		resOk.Handle(func(v TOk) {
			handled = true
		}, func(e error) {
			t.Fatal("Handle should not call Err func on Ok result")
		})
		if !handled {
			t.Fatal("Handle should call Ok func on Ok result")
		}

		handled = false
		resErr.Handle(func(v TOk) {
			t.Fatal("Handle should not call Ok func on Err result")
		}, func(e error) {
			handled = true
		})
		if !handled {
			t.Fatal("Handle should call Err func on Err result")
		}
	})
}
