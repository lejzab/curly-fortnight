package main

import (
	"testing"
)

func TestFibb(t *testing.T) {
	control := []int64{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	fibb := Fibb(100)
	if len(control) != len(fibb) {
		t.Errorf("Wrong length of data. want: %d, got: %d", len(control), len(fibb))
		return
	}
	for i := 0; i < len(control); i++ {
		if fibb[i] != control[i] {
			t.Errorf("Wrong data in %d position, want: %d, got: %d.", i, control[i], fibb[i])
		}
	}
}

func TestSumEven(t *testing.T) {
	sum := SumEven(Fibb(100))
	if sum != 44 {
		t.Errorf("SumEven was incorrect, want: %d, got: %d.", 44, sum)
	}
}
