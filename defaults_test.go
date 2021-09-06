package cond_test

import (
	"testing"

	"github.com/drykit-go/cond"
)

func TestDefaultValue(t *testing.T) {
	for _, tc := range []struct {
		val, def, exp interface{}
	}{
		{val: 0, def: "default", exp: "default"},
		{val: 42, def: "default", exp: 42},
		{val: "", def: "default", exp: "default"},
		{val: "hi", def: "default", exp: "hi"},
		{val: struct{ b bool }{}, def: "default", exp: "default"},
		{val: struct{ b bool }{true}, def: "default", exp: struct{ b bool }{true}},
	} {
		if got := cond.DefaultValue(tc.val, tc.def); got != tc.exp {
			t.Errorf(
				"cond.DefaultValue(%v, %v) -> exp %v, got %v",
				tc.val, tc.def, tc.exp, got,
			)
		}
	}
}

func TestDefaultString(t *testing.T) {
	for _, tc := range []struct {
		val, def, exp string
	}{
		{val: "", def: "default", exp: "default"},
		{val: "hi", def: "default", exp: "hi"},
	} {
		if got := cond.DefaultString(tc.val, tc.def); got != tc.exp {
			t.Errorf(
				"cond.DefaultString(%v, %v) -> exp %v, got %v",
				tc.val, tc.def, tc.exp, got,
			)
		}
	}
}

func TestDefaultInt(t *testing.T) {
	for _, tc := range []struct {
		val, def, exp int
	}{
		{val: 0, def: -1, exp: -1},
		{val: 42, def: -1, exp: 42},
	} {
		if got := cond.DefaultInt(tc.val, tc.def); got != tc.exp {
			t.Errorf(
				"cond.DefaultInt(%v, %v) -> exp %v, got %v",
				tc.val, tc.def, tc.exp, got,
			)
		}
	}
}
