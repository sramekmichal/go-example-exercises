package main

import (
	"math/cmplx"
	"testing"
)

func TestAbs(t *testing.T) {
	got := cmplx.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}