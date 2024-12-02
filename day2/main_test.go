package main

import "testing"

type testpair struct {
	levels         []int
	expectedResult bool
}

func TestSafeLevels(t *testing.T) {

	var tests = []testpair{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 8}, false},
		{[]int{1, 2, 3, 4, 7}, true},
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, pair := range tests {
		report := NewReport(pair.levels)
		safeLevels := report.isSafe()
		if safeLevels != pair.expectedResult {
			t.Error(
				"For", pair.levels,
				"expected", pair.expectedResult,
				"got", safeLevels,
			)
		}
	}
}

func TestSafeLevelsWithDampener(t *testing.T) {
	var tests = []testpair{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 8}, true},
		{[]int{1, 2, 3, 4, 7}, true},
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
		{[]int{14, 17, 20, 21, 24, 26, 27, 24}, true},
		{[]int{4, 1, 2, 3, 4}, true},
	}

	for i, pair := range tests {
		report := Report{levels: pair.levels, dampener: true}
		safeLevels := report.isSafeWithDampen()
		if safeLevels != pair.expectedResult {
			t.Error(
				"For", pair.levels,
				"at index", i,
				"expected", pair.expectedResult,
				"got", safeLevels,
			)
		}
	}
}

func TestDampen(t *testing.T) {
	report := Report{levels: []int{1, 2, 3, 4, 8}, dampener: true}
	report.dampen(3)
	if report.dampener {
		t.Error(
			"Expected dampener to be false, got true",
		)
	}

	if len(report.levels) != 4 {
		t.Error(
			"Expected 4 levels, got", len(report.levels),
		)
	}

	if report.levels[3] != 8 {
		t.Error(
			"Expected 8, got", report.levels[3],
		)
	}

	report = Report{levels: []int{1, 2, 3, 4, 8}, dampener: true}
	report.dampen(4)
	if report.dampener {
		t.Error(
			"Expected dampener to be false, got true",
		)
	}

	if len(report.levels) != 4 {
		t.Error(
			"Expected 4 levels, got", len(report.levels),
		)
	}

	if report.levels[3] != 4 {
		t.Error(
			"Expected 4, got", report.levels[3],
		)
	}
}
