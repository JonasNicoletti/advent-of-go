package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("input_1.txt")
	expected := 2810

	if result != expected {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2("input_1.txt")
	expected := 69110

	if result != expected {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}
