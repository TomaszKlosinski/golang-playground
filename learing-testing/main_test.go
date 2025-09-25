package main

import "testing"

func TestAdd(t *testing.T) {
	l := 1
	r := 2
	expected := 3

	got := add(l, r)

	if expected != got {
		t.Errorf("Expected: %v + %v = %v, but got: %v", l, r, expected, got)
	}
}
