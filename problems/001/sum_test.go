package main

import "testing"

func TestSimpleSum(t *testing.T) {
	total := SimpleSum(10)
	if total != 23 {
		t.Errorf("SimpleSum was incorrect, want: %d, got: %d.", 23, total)
	}
}

func TestArithmeticSum(t *testing.T) {
	total := ArithmeticSum(10)
	if total != 23 {
		t.Errorf("SimpleSum was incorrect, want: %d, got: %d.", 23, total)
	}
}
