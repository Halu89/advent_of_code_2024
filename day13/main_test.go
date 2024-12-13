package main

import "testing"

type testpair struct {
	input            Machine
	hasSolution      bool
	expectedSolution Vec
}

func TestHasSolution(t *testing.T) {
	// ...
	tests := []testpair{
		{Machine{Vec{26, 66}, Vec{67, 21}, Vec{12748, 12176}}, false, Vec{}},
		{Machine{Vec{17, 86}, Vec{84, 37}, Vec{7870, 6450}}, true, Vec{38, 86}},
		{Machine{Vec{69, 23}, Vec{27, 71}, Vec{18641, 10279}}, false, Vec{}},
		{Machine{Vec{94, 34}, Vec{22, 67}, Vec{8400, 5400}}, true, Vec{80, 40}},
	}

	for _, pair := range tests {
		v, ok := hasSolution(pair.input)

		if v != pair.expectedSolution {
			t.Error(
				"For", pair.input,
				"expected", pair.expectedSolution,
				"got", v,
			)
		}

		if ok != pair.hasSolution {
			t.Error(
				"For", pair.input,
				"expected", pair.hasSolution,
				"got", ok,
			)
		}
	}
}

func TestCalculatePrice(t *testing.T) {
	tests := []struct {
		input    Machine
		expected int64
	}{
		{Machine{Vec{26, 66}, Vec{67, 21}, Vec{12748, 12176}}, 0},
		{Machine{Vec{17, 86}, Vec{84, 37}, Vec{7870, 6450}}, 200},
		{Machine{Vec{69, 23}, Vec{27, 71}, Vec{18641, 10279}}, 0},
		{Machine{Vec{94, 34}, Vec{22, 67}, Vec{8400, 5400}}, 280},
	}

	for _, pair := range tests {
		v := calculatePrice(pair.input)

		if v != pair.expected {
			t.Error(
				"For", pair.input,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
