package utils

import "reflect"

func TypeFor[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func IsNilOrZero[T any](val T) bool {
	v := reflect.ValueOf(val)
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Interface, reflect.Func:
		return v.IsNil()
	default:
		zeroValue := reflect.Zero(v.Type()).Interface()
		return reflect.DeepEqual(val, zeroValue)
	}
}
