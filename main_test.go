package main

import "testing"

func TestAdd(t *testing.T) {
	var expected int = 12

	var actual int = Add(5, 6)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}
