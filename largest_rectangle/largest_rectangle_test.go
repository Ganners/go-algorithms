package main

import "testing"

// Simple test for the largest rectangle problem
func TestLargestRectangle(t *testing.T) {

	histogram := []int{6, 2, 5, 4, 5, 1, 6}
	fixture := 12
	result := LargestRectangle(histogram)

	if result != fixture {
		t.Errorf("Expected %d, got %d", fixture, result)
	}
}
