package option

import (
	"fmt"
)

type IResult[TOk any] interface {
	IsOk() bool
	IsErr() bool
	Try(func(TOk)) bool
	Do(func(TOk)) error
	Must(func(TOk))
	Handle(func(TOk), func(error))
}

type Result[TOk any] struct {
	ok   TOk
	err  error
	isOk bool
}

func Ok[TOk any](value TOk) Result[TOk] {
	return Result[TOk]{ok: value, isOk: true}
}

func Err[TOk any](err error) Result[TOk] {
	return Result[TOk]{err: err, isOk: false}
}

func ToResult[TOk any](value TOk, err error) Result[TOk] {
	if err != nil {
		return Result[TOk]{ok: value, isOk: true}
	}
	return Result[TOk]{err: err, isOk: false}
}

func (r Result[TOk]) IsOk() bool {
	return r.isOk
}

func (r Result[TOk]) IsErr() bool {
	return !r.isOk
}
func (r Result[TOk]) Try(f func(TOk)) bool {
	if r.isOk {
		f(r.ok)
	}
	return r.isOk
}

func (r Result[TOk]) Do(f func(TOk)) error {
	if r.isOk {
		f(r.ok)
		return *new(error)
	}
	return r.err
}

func (r Result[TOk]) Must(f func(TOk)) {
	if r.isOk {
		f(r.ok)
	} else {
		panic(fmt.Errorf("result is error: %v", r.err))
	}
}

func (r Result[TOk]) Handle(fOk func(TOk), fErr func(error)) {
	if r.isOk {
		fOk(r.ok)
	} else {
		fErr(r.err)
	}
}
