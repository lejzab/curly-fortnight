package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(10)
	if total != 23 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 23)
	}
}
