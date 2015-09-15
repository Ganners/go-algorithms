package anagrams

import (
	"log"
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {

	haystack := "abcdogodfa"
	needle := "odg"
	expected := []string{"dog", "god"}

	result := findAnagram(haystack, needle)
	log.Println(result, expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result incorrect, expected %v got %v", expected, result)
	}

}
