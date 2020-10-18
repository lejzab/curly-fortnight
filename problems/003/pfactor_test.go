package main

import (
	"testing"
)

func TestFactorize(t *testing.T) {
	n := NewNumber(13195)
	n.Factorize()
	if len(n.Factors) != 4 {
		t.Errorf("Wrong number of pf, want 4, got %d.", len(n.Factors))
	}
	for _, v := range []int{5, 7, 13, 29} {
		if _, ok := n.Factors[v]; !ok {
			t.Errorf("%d not in factors.", v)
		}
	}
	if n.LargestFactor != 29 {
		t.Errorf("Wrong largest factor, want: 29, got %d.", n.LargestFactor)
	}
}
