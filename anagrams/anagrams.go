package anagrams

func findAnagram(haystack, needle string) []string {

	needleLen := len(needle)
	lettersFound := 0
	firstIndex := -1
	needleMap := make(map[rune]int)
	anagrams := []string{}

	// Build the needle map to reduce runtime complexity
	for _, l := range needle {
		needleMap[rune(l)] = 0
	}

	// Loop our haystack, look for occurances of the needle
	for i := 0; i < len(haystack); i++ {

		l := rune(haystack[i])
		val, found := needleMap[rune(l)]

		switch {
		case found && val == 0:

			lettersFound++
			needleMap[l]++

			if firstIndex < 0 {
				firstIndex = i
			}

			// Check if we have a win
			if lettersFound < needleLen {
				continue
			}

			anagrams = append(anagrams, haystack[firstIndex:firstIndex+needleLen])

			// Fall through to the reset stuff
			fallthrough
		default:

			if lettersFound > 0 {
				i = firstIndex + 1
			}

			// Reset check vars
			lettersFound = 0
			firstIndex = -1

			// Reset the needle map
			for l, _ := range needleMap {
				needleMap[l] = 0
			}
		}
	}

	return anagrams
}
