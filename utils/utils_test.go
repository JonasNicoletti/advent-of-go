package utils

import (
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := ReadInput("input_part1_test.txt")
	expected := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	if input == nil {
		t.Errorf("expected '%q' but got mil", expected)

	}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("expected '%q' but got '%q'", expected, input)
	}
}
