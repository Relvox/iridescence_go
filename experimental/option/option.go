package option

import (
	"fmt"

	"github.com/relvox/iridescence_go/utils"
)

type IOption[TVal any] interface {
	HasValue() bool
	Try(f func(TVal)) bool
	Must(f func(TVal))
	Do(fNone func(), fSome func(TVal))
}

type Option[T any] struct {
	value T
	isSet bool
}

func None[TVal any]() Option[TVal] {
	return Option[TVal]{}
}

func Some[TVal any](value TVal) Option[TVal] {
	return Option[TVal]{value, true}
}

func Maybe[TVal any](value TVal) Option[TVal] {
	if utils.IsNilOrZero(value) {
		return Option[TVal]{}
	}
	return Option[TVal]{value, true}
}

func ToOption[TVal any](value TVal, ok bool) Option[TVal] {
	if !ok {
		return Option[TVal]{}
	}
	return Option[TVal]{value, true}
}

func (o Option[TVal]) HasValue() bool {
	return o.isSet
}

func (o Option[TVal]) Try(f func(TVal)) bool {
	if o.isSet {
		f(o.value)
		return true
	}
	return false
}

func (o Option[TVal]) Do(fNone func(), fSome func(TVal)) {
	if o.isSet {
		fSome(o.value)
	} else {
		fNone()
	}
}

func (o Option[TVal]) Must(f func(TVal)) {
	if o.isSet {
		f(o.value)
	}
	panic(fmt.Errorf("option is none"))
}
