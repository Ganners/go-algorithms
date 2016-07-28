package dash

// DictionaryDash will take a start and end word, and a dictionary.
// Using that dictionary it will find the length with the shortest
// transformation
func DictionaryDash(start, end string, dictionary []string) int {

	// Convert dictionary into a set
	dictSet := toSet(dictionary)

	if len(start) != len(end) {
		// return ErrLengthsDoNotMatch
		return -1
	}

	if start == end {
		return 0
	}

	transformations := 0

TryAgain:
	for i := 0; i < len(start); i++ {

		if start == end {
			break
		}

		// Difference
		if start[i] != end[i] {

			// Incorrect - This makes the assumption that there are linear one
			// character changes that will always exist in the dictionary that
			// match the end word.

			// Construct what the new word would be
			word := start[:i] + string(end[i]) + start[i+1:]

			if _, found := dictSet[word]; found {
				start = word
				transformations++
			}
		}
	}

	if start != end {
		goto TryAgain
	}

	return transformations
}

// Convert the slice of strings into a set (map of strings)
func toSet(strs []string) map[string]struct{} {
	set := make(map[string]struct{}, len(strs))
	for _, str := range strs {
		set[str] = struct{}{}
	}
	return set
}
