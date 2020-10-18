package main

import "testing"

func TestSimpleSum(t *testing.T) {
	total := SimpleSum(10)
	if total != 23 {
		t.Errorf("SimpleSum was incorrect, got: %d, want: %d.", total, 23)
	}
}
