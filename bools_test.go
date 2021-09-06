package cond_test

import (
	"testing"

	"github.com/drykit-go/cond"
)

func TestXOR(t *testing.T) {
	for _, tc := range []struct {
		a, b, exp bool
	}{
		{a: true, b: true, exp: false},
		{a: true, b: false, exp: true},
		{a: false, b: true, exp: true},
		{a: false, b: false, exp: false},
	} {
		if got := cond.XOR(tc.a, tc.b); got != tc.exp {
			t.Errorf("cond.XOR(%v, %v) -> exp %v, got %v", tc.a, tc.b, tc.exp, got)
		}
	}
}

func TestAnd(t *testing.T) {
	for _, tc := range []struct {
		in  []bool
		exp bool
	}{
		{in: []bool{true, true, true}, exp: true},
		{in: []bool{true, false, true}, exp: false},
		{in: []bool{false, false, false}, exp: false},
	} {
		if got := cond.And(tc.in...); got != tc.exp {
			t.Errorf("cond.And(%v) -> exp %v, got %v", tc.in, tc.exp, got)
		}
	}
}

func TestOr(t *testing.T) {
	for _, tc := range []struct {
		in  []bool
		exp bool
	}{
		{in: []bool{true, true, true}, exp: true},
		{in: []bool{false, true, false}, exp: true},
		{in: []bool{false, false, false}, exp: false},
	} {
		if got := cond.Or(tc.in...); got != tc.exp {
			t.Errorf("cond.Or(%v) -> exp %v, got %v", tc.in, tc.exp, got)
		}
	}
}

func TestNone(t *testing.T) {
	for _, tc := range []struct {
		in  []bool
		exp bool
	}{
		{in: []bool{true, true, true}, exp: false},
		{in: []bool{false, true, false}, exp: false},
		{in: []bool{false, false, false}, exp: true},
	} {
		if got := cond.None(tc.in...); got != tc.exp {
			t.Errorf("cond.None(%v) -> exp %v, got %v", tc.in, tc.exp, got)
		}
	}
}

func TestNotAll(t *testing.T) {
	for _, tc := range []struct {
		in  []bool
		exp bool
	}{
		{in: []bool{true, true, true}, exp: false},
		{in: []bool{false, true, false}, exp: true},
		{in: []bool{false, false, false}, exp: true},
	} {
		if got := cond.NotAll(tc.in...); got != tc.exp {
			t.Errorf("cond.NotAll(%v) -> exp %v, got %v", tc.in, tc.exp, got)
		}
	}
}
