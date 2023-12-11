package day_1

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

func TestFindFirstAndLastPart1(t *testing.T) {
	first := FindFirstAndLastPart1("pqr31stu8vwx")
	expected := 38

	if first != expected {
		t.Errorf("expected '%q' but got '%q'", expected, first)
	}

	none := FindFirstAndLastPart1("pqrstuvwx")
	expected = 0

	if none != expected {
		t.Errorf("expected '%q' but got '%q'", expected, none)
	}
}

func TestPart1(t *testing.T) {
	result := Part1("input_part1_test.txt")
	expected := 142

	if result != expected {
		t.Errorf("expected '%q' but got '%q'", expected, result)
	}
}

func TestMatchNumberOrString(t *testing.T) {
	one := MatchNumberOrString("ab")
	expected := ""

	if one != expected {
		t.Errorf("expected '%v' but got '%v'", expected, one)
	}

	one = MatchNumberOrString("abcon")
	expected = ""

	if one != expected {
		t.Errorf("expected '%v' but got '%v'", expected, one)
	}

	one = MatchNumberOrString("abcone")
	expected = "one"

	if one != expected {
		t.Errorf("expected '%v' but got '%v'", expected, one)
	}
}

func TestFindFirstAndLastPart2(t *testing.T) {
	first := FindFirstAndLastPart2("two1nine")
	expected := 29

	if first != expected {
		t.Errorf("expected '%v' but got '%v'", expected, first)
	}

	first = FindFirstAndLastPart2("xtwone3four")
	expected = 24

	if first != expected {
		t.Errorf("expected '%v' but got '%v'", expected, first)
	}

	first = FindFirstAndLastPart2("eightwothree")
	expected = 83

	if first != expected {
		t.Errorf("expected '%v' but got '%v", expected, first)
	}

	first = FindFirstAndLastPart2("abcone2threexyz")
	expected = 13

	if first != expected {
		t.Errorf("expected '%v' but got '%v'", expected, first)
	}

	first = FindFirstAndLastPart2("xtwone3four")
	expected = 24

	if first != expected {
		t.Errorf("expected '%v' but got '%v'", expected, first)
	}

	first = FindFirstAndLastPart2("zoneight234")
	expected = 14

	if first != expected {
		t.Errorf("expected '%v' but got '%v'", expected, first)
	}

	first = FindFirstAndLastPart2("7pqrstsixteen")
	expected = 76

	if first != expected {
		t.Errorf("expected '%v' but got '%v'", expected, first)
	}

	none := FindFirstAndLastPart2("pqrstuvwx")
	expected = 0

	if none != expected {
		t.Errorf("expected '%v' but got '%v'", expected, none)
	}

}

func TestPart2(t *testing.T) {
	result := Part2("input_part2_test.txt")
	expected := 281

	if result != expected {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}
