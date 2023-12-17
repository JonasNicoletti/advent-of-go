package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("input_1.txt")
	expected := 2810

	if result != expected {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}
