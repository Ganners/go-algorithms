package main

import "log"

// Given a pattern and a string, find if a string follows the same pattern
func main() {

	pattern := []rune{'a', 'b', 'b', 'a'}
	str := []string{"cat", "dog", "dog", "cat"}

	result := checkPatternsMatch(pattern, str)
	log.Println(result)
}

func checkPatternsMatch(pattern []rune, str []string) bool {

	// Can't match if the lengths are different
	if len(pattern) != len(str) {
		return false
	}

	// I suppose two empties are the same pattern
	if len(pattern) == 0 {
		return true
	}

	memo := make(map[rune]string)
	seenWords := make(map[string]struct{})

	for i, word := range str {

		if memoWord, found := memo[pattern[i]]; found {
			if memoWord != word {
				return false
			}
		} else {
			if _, seen := seenWords[word]; seen {
				// We've seen it and it wasn't stored in our memo, bad!
				return false
			}
			seenWords[word] = struct{}{}
			memo[pattern[i]] = word
		}
	}

	return true
}
