package main

import (
	"regexp"
	"testing"
)

type testpair struct {
	input    string
	expected []Operand
}

func TestParseLine(t *testing.T) {
	tests := []testpair{
		{"mul(1,2)", []Operand{{1, 2}}},
		{"mul(123,2)", []Operand{{123, 2}}},
		{"mul(123,2)asl;dkfjmul(432,12)", []Operand{{123, 2}, {432, 12}}},
		{"mul(#1,2)", []Operand{{0, 0}}},
		{"mul#(1,2)", []Operand{{0, 0}}},
		{"!*&mul(363,974)", []Operand{{363, 974}}},
	}

	for _, pair := range tests {
		operands, _ := parseLine(pair.input, true)
		if len(operands) != len(pair.expected) {
			t.Error(
				"For", pair.input,
				"expected", pair.expected,
				"got", operands,
			)
		}

		for i, expectedOperand := range pair.expected {
			if operands[i].a != expectedOperand.a {
				t.Error(
					"For", pair.input,
					"expected", pair.expected,
					"got", operands)
			}

			if operands[i].b != expectedOperand.b {
				t.Error(
					"For", pair.input,
					"expected", pair.expected,
					"got", operands)
			}
		}
	}
}

func TestRegexp(t *testing.T) {
	regex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")

	input := "mul(123,2)asl;dkfjmul(432,12)"

	matches := regex.FindAllString(input, -1)

	if len(matches) != 2 {
		t.Error(
			"Expected 2 matches, got", len(matches),
		)
	}

	if matches[0] != "mul(123,2)" {
		t.Error(
			"Expected mul(123,2), got", matches[0],
		)
	}

	if matches[1] != "mul(432,12)" {
		t.Error(
			"Expected mul(432,12), got", matches[1],
		)
	}

}
