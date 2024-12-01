package main

import (
	"sort"
	"testing"
)

func TestParse(t *testing.T) {
	firstCol, secondCol := parseInput("test.txt")
	expectedFirstCol := []int{3, 4, 2, 1, 3, 3}
	expectedSecondCol := []int{4, 3, 5, 3, 9, 3}

	sort.Ints(expectedFirstCol)
	sort.Ints(expectedSecondCol)

	if len(firstCol) != 6 {
		t.Errorf("Expected 6 lines in the test file, got %d", len(firstCol))
	}

	if len(secondCol) != len(firstCol) {
		t.Errorf("Expected the same number of lines in the 2 columns")
	}

	for i := 0; i < len(firstCol); i++ {
		if firstCol[i] != expectedFirstCol[i] {
			t.Errorf("Expected %d at index %d for first column, got %d", expectedFirstCol[i], i, firstCol[i])
		}

		if secondCol[i] != expectedSecondCol[i] {
			t.Errorf("Expected %d at index %d for second column, got %d", expectedSecondCol[i], i, secondCol[i])
		}
	}
}

func TestSimilarities(t *testing.T) {
	firstCol, secondCol := parseInput("test.txt")
	similarities := calculateSimilarities(firstCol, secondCol)

	if similarities != 31 {
		t.Errorf("Expected 31, got %d", similarities)
	}
}

func TestDistance(t *testing.T) {
	firstCol, secondCol := parseInput("test.txt")
	distances := calculateDistance(firstCol, secondCol)

	if distances != 11 {
		t.Errorf("Expected 11, got %d", distances)
	}
}
