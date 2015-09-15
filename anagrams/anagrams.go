package anagrams

import "errors"

func findAnagram(haystack, needle string) (string, error) {

	needleLen := len(needle)
	lettersFound := 0
	firstIndex := -1
	needleMap := make(map[rune]int)

	// Build the needle map to reduce runtime complexity
	for i, l := range needle {
		needleMap[rune(l)] = i
	}

	// Loop our haystack, look for occurances of the needle
	for i, l := range haystack {
		if _, found := needleMap[rune(l)]; found {

			lettersFound++

			if firstIndex < 0 {
				firstIndex = i
			}

			// Check if we have a win
			if lettersFound >= needleLen {
				return haystack[firstIndex : firstIndex+needleLen], nil
			}
		} else {
			// Reset check vars
			lettersFound = 0
			firstIndex = -1
		}
	}

	return "", errors.New("No anagram found")
}
