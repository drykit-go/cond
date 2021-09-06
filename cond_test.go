package cond_test

import (
	"errors"
	"testing"

	"github.com/drykit-go/cond"
)

func TestPanicOnErr(t *testing.T) {
	newError := func(nonNil bool) error {
		if nonNil {
			return errors.New("error")
		}
		return nil
	}

	t.Run("nil errors", func(t *testing.T) {
		nilerrs := []error{
			newError(false),
			newError(false),
			newError(false),
		}
		defer assertNoPanic(t, "PanicOnErr/nil errors")
		cond.PanicOnErr(nilerrs...)
	})

	t.Run("nonnil errors", func(t *testing.T) {
		nilerrs := []error{
			newError(false),
			newError(true),
			newError(false),
		}
		defer assertPanic(t, "PanicOnErr/nonnil errors", nilerrs[1])
		cond.PanicOnErr(nilerrs...)
	})
}

func assertPanic(t *testing.T, label string, expv interface{}) {
	t.Helper()
	r := recover()
	if r == nil {
		t.Errorf("%s: expect to panic, did not", label)
	}
	if expv != nil && r != expv {
		t.Errorf("%s: bad panic value: exp %v, got %v", label, expv, r)
	}
}

func assertNoPanic(t *testing.T, label string) {
	t.Helper()
	if r := recover(); r != nil {
		t.Errorf("%s: unexpected panic", label)
	}
}
