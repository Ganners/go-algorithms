package anagrams

import "testing"

func TestFindAnagrams(t *testing.T) {

	haystack := "abcdogfa"
	needle := "odg"
	expected := "dog"

	result, err := findAnagram(haystack, needle)

	if err != nil {
		t.Errorf("Did not expect an error, got %s", err)
	}

	if result != expected {
		t.Errorf("Result incorrect, expected %s got %s", expected, result)
	}

}
