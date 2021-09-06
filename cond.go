package cond

import "reflect"

// Zero returns true if v is a zero value.
func Zero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

// OnErr calls f(err) on each non-nil error.
func OnErr(f func(error), errs ...error) {
	for _, err := range errs {
		if err != nil {
			f(err)
		}
	}
}

// PanicOnErr panics on the first non-nil error.
func PanicOnErr(errs ...error) {
	OnErr(func(err error) { panic(err) }, errs...)
}
