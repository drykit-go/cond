package cond_test

import (
	"testing"

	"github.com/drykit-go/cond"
)

func TestValue(t *testing.T) {
	const (
		vtrue  = 42
		vfalse = "default"
	)

	for b, exp := range map[bool]interface{}{true: vtrue, false: vfalse} {
		if got := cond.Value(vtrue, vfalse, b); got != exp {
			t.Errorf("cond.Value(%v, %v, %v) -> exp %v, got %v", vtrue, vfalse, b, exp, got)
		}
	}
}

func TestString(t *testing.T) {
	const (
		vtrue  = "hello"
		vfalse = "default"
	)

	for b, exp := range map[bool]interface{}{true: vtrue, false: vfalse} {
		if got := cond.String(vtrue, vfalse, b); got != exp {
			t.Errorf("cond.String(%v, %v, %v) -> exp %v, got %v", vtrue, vfalse, b, exp, got)
		}
	}
}

func TestInt(t *testing.T) {
	const (
		vtrue  = 42
		vfalse = -1
	)

	for b, exp := range map[bool]interface{}{true: vtrue, false: vfalse} {
		if got := cond.Int(vtrue, vfalse, b); got != exp {
			t.Errorf("cond.Int(%v, %v, %v) -> exp %v, got %v", vtrue, vfalse, b, exp, got)
		}
	}
}
