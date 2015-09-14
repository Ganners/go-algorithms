package main

import "log"

func isWord(word string) bool {

	switch word {
	case "hi":
		return true
		break
	case "go":
		return true
		break
	}

	return false
}

// Helper
func getLetterPermutations(letter string) []string {

	switch letter {
	case "g":
		return []string{"f", "g", "h"}
		break
	case "i":
		return []string{"k", "i", "o"}
		break
	}

	return []string{}
}

// Runtime complexity: O(N^M)
//
// fi
//  fk
//  fi
//  fo
// gi
//  gk
//  gi
//  go
// hi
//  fk
//  fi
//  fo
//
// How to reduce...
func permutationIterator(
	word string,
	index int,
	filter func(word string) bool,
) []string {

	if index < 0 || index > len(word)-1 {
		return []string{}
	}

	words := make([]string, 0)

	for _, letter := range getLetterPermutations(string(word[index])) {

		// Swap out the first letter
		newWord := word[:index] + letter + word[index+1:]

		// If we cannot recurse then we run filter on the word
		if index == len(word)-1 && filter(newWord) {
			words = append(words, newWord)
			continue
		}

		// Append words from children
		words = append(words, permutationIterator(newWord, index+1, filter)...)
	}

	return words
}

func getNearbyWords(word string) []string {

	if len(word) < 1 {
		return []string{}
	}

	return permutationIterator(word, 0, func(word string) bool {

		if isWord(word) {
			return true
		}
		return false
	})
}

func main() {

	words := getNearbyWords("gi")
	log.Println(words)
}
