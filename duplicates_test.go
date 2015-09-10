package main

import (
	"reflect"
	"testing"
)

func TestDuplicates(t *testing.T) {

	fixture := []int{
		5, 21, 59, 41, 4, 12, 1, 5, 59, 12, 7, 14, 5, 11, 23, 41, 5, 21, 4, 99,
		5, 21, 59, 41, 4, 12, 1, 5, 59, 12, 7, 14, 5, 11, 23, 41, 5, 21, 4, 99,
	}
	expected := []int{
		5, 21, 59, 41, 4, 12, 1, 99, 11, 7, 23, 14,
	}

	expectedRemoved := len(fixture) - len(expected)
	result, numFound := removeDuplicates(fixture)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"Result from duplicate removal did not match expected %v", expected)
	}

	if numFound != expectedRemoved {
		t.Error(
			"Number of duplicates not found (%d) does not match expected (%d)",
			expectedRemoved, numFound)
	}
}
