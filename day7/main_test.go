package main

import "testing"

func TestEnumerate(t *testing.T) {
	// Test cases
	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{3, 8},
	}

	operators := []Operator{Plus, Multiply}

	// Run tests
	for _, test := range tests {
		actual := enumerate(test.input, operators)

		if len(actual) != test.expected {
			t.Errorf("Expected: %v. Actual: %v", test.expected, actual)
		}

		for _, list := range actual {
			if len(list) != test.input {
				t.Errorf("Expected: %v. Actual: %v", test.input, len(list))
			}
		}
	}
}

func TestParseEquation(t *testing.T) {
	tests := []struct {
		input    string
		expected Equation
	}{
		{
			"1: 2 3 4",
			Equation{target: 1, operands: []int{2, 3, 4}},
		},

		{
			"21037: 9 7 18 13",
			Equation{target: 21037, operands: []int{9, 7, 18, 13}},
		},
	}

	for _, test := range tests {
		actual := parseEquation(test.input)

		if actual.target != test.expected.target {
			t.Errorf("Expected: %v. Actual: %v", test.expected.target, actual.target)
		}

		if len(actual.operands) != len(test.expected.operands) {
			t.Errorf("Expected: %v. Actual: %v", len(test.expected.operands), len(actual.operands))
		}

		for i, operand := range test.expected.operands {
			if actual.operands[i] != operand {
				t.Errorf("Expected: %v. Actual: %v", operand, actual.operands[i])
			}
		}

	}
}
