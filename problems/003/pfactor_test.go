package main

import (
	"testing"
)

func TestPrimeFacror(t *testing.T) {
	pf := PrimeFator(13195)
	if len(pf) != 4 {
		t.Errorf("Wrong number of pf, want 4, got %d.", len(pf))
	}
	for _, v := range []int{5, 7, 13, 29} {
		if _, ok := pf[v]; !ok {
			t.Errorf("%d not in factors.", v)
		}
	}
}
