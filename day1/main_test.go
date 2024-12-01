package main

import "testing"

func TestDistance(t *testing.T) {
	dist := calculateDistance("test.txt")

	if dist != 11 {
		t.Errorf("Expected 11, got %d", dist)
	}
}

func TestSimilarities(t *testing.T) {
	sim := step2("test.txt")

	if sim != 31 {
		t.Errorf("Expected 31, got %d", sim)
	}
}
